package main

import (
	"context"
	pb "go-grpc/pb"
	"log"
	"net"

	"google.golang.org/grpc"
)

type HelloService struct {
	pb.UnimplementedHelloServer
}

func (s *HelloService) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResp, error) {
	log.Println(req.Name)

	return &pb.HelloResp{
		Message: "hello " + req.Name,
	}, nil
}

const (
	Addr    string = ":8080"
	Network string = "tcp"
)

func main() {
	listener, err := net.Listen(Network, Addr)
	if err != nil {
		log.Panic("net.Listen err: %v", err)
	}

	log.Println(Addr + " net.Listening...")

	grpcServer := grpc.NewServer()

	pb.RegisterHelloServer(grpcServer, &HelloService{})

	err = grpcServer.Serve(listener)

	if err != nil {
		log.Panic("grpcServer.Serve err: %v", err)
	}
}
