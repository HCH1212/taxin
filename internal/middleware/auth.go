package middleware

// grpc的认证中间件

import (
	"context"
	"errors"
	"strings"

	"github.com/HCH1212/taxin/internal/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// AuthInterceptor 是一个 gRPC 一元拦截器，用于鉴权
func AuthInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		// 只拦截 GetUserInfo 方法
		if info.FullMethod != "/user.UserService/GetUserInfo" {
			return handler(ctx, req)
		}

		// 从元数据中获取 Authorization 头
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, errors.New("missing metadata")
		}

		// 获取 token
		authHeader := md.Get("authorization")
		if len(authHeader) == 0 {
			return nil, errors.New("missing authorization token")
		}

		tokenString := authHeader[0]

		// 去除 "Bearer " 前缀（忽略大小写）
		const bearerPrefix = "Bearer "
		if len(tokenString) > len(bearerPrefix) &&
			strings.EqualFold(tokenString[:len(bearerPrefix)], bearerPrefix) {
			tokenString = tokenString[len(bearerPrefix):]
		}

		// 解析 token
		claims, err := utils.ParseAccessToken(tokenString)
		if err != nil {
			return nil, err
		}

		// 将用户 ID 添加到上下文
		ctx = context.WithValue(ctx, "user_id", claims.UserID)

		// 调用下一个处理程序
		return handler(ctx, req)
	}
}
