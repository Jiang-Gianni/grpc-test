package main

import (
	"github/Jiang-Gianni/grpc-test/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewGRPCClient() (proto.GreeterClient, error) {
	conn, err := grpc.Dial(":3000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	c := proto.NewGreeterClient(conn)

	return c, nil
}
