package service

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	pb "github.com/HCH1212/taxin/api/pb/user"
	"github.com/HCH1212/taxin/internal/dao"
	"github.com/HCH1212/taxin/internal/model"
	"github.com/HCH1212/taxin/internal/utils"
	"github.com/go-redis/redis/v8"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type UserService struct {
	pb.UnimplementedUserServiceServer
}

// Register 注册新用户
func (u *UserService) Register(ctx context.Context, req *pb.RegisterReq) (*pb.RegisterResp, error) {
	// 创建一个新的 span
	tr := otel.Tracer("user-service")
	_, span := tr.Start(ctx, "Register")
	defer span.End()
	// 参数校验
	if req.Password == "" || len(req.Like) == 0 || req.Username == "" {
		span.SetStatus(codes.Error, "invalid request")
		return nil, errors.New("invalid request")
	}
	// 注册幂等性校验
	redisKey := "register:redis:" + req.Username
	if userID, err := dao.RedisClient.Get(ctx, redisKey).Result(); err == nil {
		span.AddEvent("register success")
		return &pb.RegisterResp{UserId: userID}, nil
	} else if err != redis.Nil {
		span.SetStatus(codes.Error, "redis error")
		return nil, err
	}
	userID, err := model.GetUserIDByUsername(dao.DB, req.Username)
	if err == nil {
		span.AddEvent("register success")
		return &pb.RegisterResp{UserId: userID}, nil
	} else if err != gorm.ErrRecordNotFound {
		span.SetStatus(codes.Error, "database error")
		return nil, err
	}
	// 以下开始注册新用户
	// 密码加密
	hashPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		span.SetStatus(codes.Error, "hash password failed")
		return nil, err
	}
	// 生成词嵌入向量
	embedding, err := utils.GenerateEmbeddingForLikes(ctx, req.Like)
	if err != nil {
		span.SetStatus(codes.Error, "generate embedding failed")
		return nil, err
	}
	// 生成用户ID并创建用户
	userID = utils.GenerateUUID()
	span.SetAttributes(attribute.String("user_id", userID))
	// 爱好转json
	likeJSON, err := json.Marshal(req.Like)
	if err != nil {
		span.SetStatus(codes.Error, "marshal like failed")
		return nil, err
	}
	user := model.User{
		Username:      req.Username,
		UserID:        userID,
		Password:      hashPassword,
		Like:          datatypes.JSON(likeJSON),
		LikeEmbedding: embedding,
	}
	// 存储注册信息到redis
	err = dao.RedisClient.Set(ctx, redisKey, userID, time.Hour*24).Err()
	if err != nil {
		span.SetStatus(codes.Error, "store register info to redis failed")
		return nil, err
	}
	// 存储用户信息到数据库
	err = model.CreateUser(dao.DB, &user)
	if err != nil {
		span.SetStatus(codes.Error, "create user failed")
		return nil, err
	}

	return &pb.RegisterResp{UserId: userID}, nil
}

// Login 登录
func (u *UserService) Login(ctx context.Context, req *pb.LoginReq) (*pb.LoginResp, error) {
	// 创建一个新的 span
	tr := otel.Tracer("user-service")
	_, span := tr.Start(ctx, "Login")
	defer span.End()
	// 添加自定义标签
	span.SetAttributes(attribute.String("user_id", req.UserId))

	// 参数校验
	if req.UserId == "" || req.Password == "" {
		span.SetStatus(codes.Error, "invalid request")
		return nil, errors.New("invalid request")
	}
	// 查询用户信息
	user, err := model.GetUserByUserID(dao.DB, req.UserId)
	if err != nil {
		span.SetStatus(codes.Error, "user not found")
		return nil, err
	}
	// 验证密码
	if !utils.VerifyPassword(user.Password, req.Password) {
		span.SetStatus(codes.Error, "invalid password")
		return nil, errors.New("invalid password")
	}
	// 生成 access_token
	accessToken, err := utils.GetToken(req.UserId)
	if err != nil {
		span.SetStatus(codes.Error, "generate access token failed")
		return nil, err
	}

	// 添加自定义事件
	span.AddEvent("login success")

	return &pb.LoginResp{AccessToken: accessToken}, nil
}

// GetUserInfo 获取用户信息
func (u *UserService) GetUserInfo(ctx context.Context, req *pb.UserInfoReq) (*pb.UserInfoResp, error) {
	tr := otel.Tracer("user-service")
	_, span := tr.Start(ctx, "GetUserInfo")
	defer span.End()
	// 从上下文中获取用户 ID
	userID := ctx.Value("user_id")
	span.SetAttributes(attribute.String("user_id", userID.(string)))
	if userID == nil {
		span.SetStatus(codes.Error, "missing user ID in context")
		return nil, errors.New("missing user ID in context")
	}
	// 查询用户信息
	user, err := model.GetUserByUserID(dao.DB, userID.(string))
	if err != nil {
		span.SetStatus(codes.Error, "get user failed")
		return nil, err
	}
	// 组装响应
	likeEmbedding := make([]float32, 1536)
	copy(likeEmbedding, user.LikeEmbedding.Slice())
	return &pb.UserInfoResp{
		UserId:        user.UserID,
		Like:          user.GetLikeList(),
		LikeEmbedding: likeEmbedding,
		CreateAt:      user.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdateAt:      user.UpdatedAt.Format("2006-01-02 15:04:05"),
		Username:      user.Username,
	}, nil
}
