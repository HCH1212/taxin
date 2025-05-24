package main

import (
	"context"
	"log"

	pb "github.com/HCH1212/taxin/demo/grpc/hello"
	"google.golang.org/grpc"
)

func main() {
	// 连接到服务器
	conn, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	// 创建客户端
	c := pb.NewHelloServiceClient(conn)

	// 调用服务方法
	name := "World"
	response, err := c.SayHello(context.Background(), &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	log.Printf("Greeting: %s", response.GetMessage())
}
