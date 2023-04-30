package main

import (
	"context"
	"fmt"
	"github/Jiang-Gianni/grpc-test/proto"
	"net"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type GRPCServer struct {
	proto.UnimplementedGreeterServer
}

func (s *GRPCServer) SayHello(ctx context.Context, req *proto.HelloRequest) (*proto.HelloResponse, error) {
	start := time.Now()
	defer func(start time.Time) {
		fmt.Println("Time since start ", time.Since(start).Microseconds())
	}(start)
	msg := "Hello " + req.Name
	fmt.Println(msg)
	return &proto.HelloResponse{
		Message: msg,
	}, nil
}

func MakeGRPCServerAndRun() error {
	grpcServer := &GRPCServer{}
	ln, err := net.Listen("tcp", ":3000")
	if err != nil {
		return err
	}
	fmt.Println("listen address ", ln.Addr().String())
	fmt.Printf("ln is %+v \n", ln)

	opts := []grpc.ServerOption{}
	server := grpc.NewServer(opts...)
	reflection.Register(server)
	proto.RegisterGreeterServer(server, grpcServer)
	return server.Serve(ln)
}
