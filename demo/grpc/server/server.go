package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/HCH1212/taxin/demo/grpc/hello"
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

func main() {
	// 创建 gRPC 服务器
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	// 注册服务
	pb.RegisterHelloServiceServer(s, &server{})

	fmt.Println("Server is listening on port 50052")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
