package main

import (
	"context"
	"log"

	pb "go-grpc/pb"
	// "github.com/CodeFish-xiao/blogs/gRPCAction/code/grpc-5/pb"
	"google.golang.org/grpc"
)

const (
	ServerAddr string = ":8080"
)

func main() {
	// 连接服务器
	conn, err := grpc.Dial(ServerAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("net.Connect error: %v", err)
	}

	defer conn.Close()

	grpcClient := pb.NewHelloClient(conn)

	req := pb.HelloRequest{
		Name: "grpc",
	}

	res, err := grpcClient.SayHello(context.Background(), &req)
	if err != nil {
		log.Fatalf("Call SayHello error: %v", err)
	}

	log.Println(res)
}
