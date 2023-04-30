package main

import (
	"context"
	"fmt"
	"github/Jiang-Gianni/grpc-test/proto"
	"time"
)

func main() {
	go func() {
		time.Sleep(2 * time.Second)
		client, err := NewGRPCClient()
		if err != nil {
			fmt.Println(err)
		}
		client.SayHello(context.Background(), &proto.HelloRequest{
			Name: "Me",
		})
	}()
	MakeGRPCServerAndRun()
}
