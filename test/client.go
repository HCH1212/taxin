package main

import (
	"context"
	"fmt"
	"io"
	"log"

	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	pb_system "github.com/HCH1212/taxin/api/pb/system"
	pb_user "github.com/HCH1212/taxin/api/pb/user"
)

func main() {
	ctx := context.Background()
	tr := otel.Tracer("all-services-client")
	ctx, span := tr.Start(ctx, "ClientCallAllServices")
	defer span.End()

	// 连接到 gRPC 服务器
	conn, err := grpc.Dial("localhost:50052", grpc.WithInsecure(), grpc.WithStatsHandler(otelgrpc.NewClientHandler()))
	if err != nil {
		span.SetStatus(codes.Error, "Failed to connect to server")
		log.Fatalf("Failed to connect to server: %v", err)
	}
	defer conn.Close()

	// 测试 UserService
	testUserService(ctx, conn)

	// 测试 SystemService
	testSystemService(ctx, conn)

	fmt.Println("All tests completed successfully")
}

func testUserService(ctx context.Context, conn *grpc.ClientConn) {
	tr := otel.Tracer("user-service-client")
	ctx, span := tr.Start(ctx, "TestUserService")
	defer span.End()

	// 创建 UserService 客户端
	client := pb_user.NewUserServiceClient(conn)

	// 测试注册
	registerReq := &pb_user.RegisterReq{
		Password: "testpassword",
		Like:     []string{"reading", "swimming"},
		Username: "testuser7",
	}
	registerResp, err := client.Register(ctx, registerReq)
	if err != nil {
		span.SetStatus(codes.Error, "Failed to register")
		log.Fatalf("Failed to register: %v", err)
	}
	fmt.Printf("Register User ID: %s\n", registerResp.UserId)

	// 测试登录
	loginReq := &pb_user.LoginReq{
		UserId:   registerResp.UserId,
		Password: "testpassword",
	}
	loginResp, err := client.Login(ctx, loginReq)
	if err != nil {
		span.SetStatus(codes.Error, "Failed to login")
		log.Fatalf("Failed to login: %v", err)
	}
	fmt.Printf("Login Access Token: %s\n", loginResp.AccessToken)

	// 测试获取用户信息
	ctxWithToken := metadata.AppendToOutgoingContext(ctx, "authorization", "Bearer "+loginResp.AccessToken)
	userInfoReq := &pb_user.UserInfoReq{}
	userInfoResp, err := client.GetUserInfo(ctxWithToken, userInfoReq)
	if err != nil {
		span.SetStatus(codes.Error, "Failed to get user info")
		log.Fatalf("Failed to get user info: %v", err)
	}
	fmt.Printf("User ID: %s\n", userInfoResp.UserId)
	// fmt.Printf("Likes: %v\n", userInfoResp.Like)
	// fmt.Printf("Like Embedding: %v\n", userInfoResp.LikeEmbedding)
	// fmt.Printf("Create At: %s\n", userInfoResp.CreateAt)
	// fmt.Printf("Update At: %s\n", userInfoResp.UpdateAt)
	// fmt.Printf("Username %s\n", userInfoResp.Username)

	// 添加自定义标签和事件
	span.SetAttributes(attribute.String("user_id", registerResp.UserId))
	span.AddEvent("User service tests completed successfully")
}

func testSystemService(ctx context.Context, conn *grpc.ClientConn) {
	tr := otel.Tracer("system-service-client")
	ctx, span := tr.Start(ctx, "TestSystemService")
	defer span.End()

	// 创建 SystemService 客户端
	client := pb_system.NewSystemServiceClient(conn)

	// 测试发送文件
	sendFileReq := &pb_system.SendFileReq{
		FilePath: "test/test.txt",
	}
	stream, err := client.SendFile(ctx, sendFileReq)
	if err != nil {
		span.SetStatus(codes.Error, "Failed to send file request")
		log.Fatalf("Failed to send file request: %v", err)
	}

	// 接收文件流
	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			span.SetStatus(codes.Error, "Failed to receive file stream")
			log.Fatalf("Failed to receive file stream: %v", err)
		}
		fmt.Printf("Received %d bytes of file content\n", len(resp.Content))
	}

	// 添加自定义标签和事件
	span.SetAttributes(attribute.String("file_path", sendFileReq.FilePath))
	span.AddEvent("System service SendFile test completed successfully")
}
