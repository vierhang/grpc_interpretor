package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"grpc_interpretor/proto"
)

func main() {
	dial, err := grpc.Dial("127.0.0.1:50052", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer dial.Close()
	sayClient := proto.NewGreeterClient(dial)
	hello, err := sayClient.SayHello(context.Background(), &proto.HelloRequest{
		Name: "weihang",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(hello)
}
