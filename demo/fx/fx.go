package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/HCH1212/taxin/demo/grpc/hello"
	"go.uber.org/fx"
	"google.golang.org/grpc"
)

// 实现 HelloService 接口
type server struct {
	pb.UnimplementedHelloServiceServer
}

// 实现 SayHello 方法
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.HelloResponse{Message: "Hello, " + in.GetName()}, nil
}

// 创建监听套接字
func newListener() (net.Listener, error) {
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		return nil, fmt.Errorf("failed to listen: %w", err)
	}
	return lis, nil
}

// 创建 gRPC 服务器
func newGRPCServer(lc fx.Lifecycle, lis net.Listener) *grpc.Server {
	s := grpc.NewServer()
	pb.RegisterHelloServiceServer(s, &server{})

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			fmt.Println("Server is listening on port 50052")
			go func() {
				if err := s.Serve(lis); err != nil {
					log.Printf("failed to serve: %v", err)
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			s.GracefulStop()
			return nil
		},
	})

	return s
}

func main() {
	app := fx.New(
		fx.Provide(
			newListener,
			newGRPCServer,
		),
		fx.Invoke(func(*grpc.Server) {}), // 触发服务器启动
		fx.NopLogger,
	)

	app.Run()
}

// 启动这个就想相当于grpc包的server/server.go了，再启动client/client.go就能调用SayHello方法了
