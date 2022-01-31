package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"grpc_interpretor/proto"
	"time"
)

func main() {
	myInterceptor := func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		start := time.Now()
		err := invoker(ctx, method, req, reply, cc, opts...)
		fmt.Println("耗时", time.Since(start))
		return err
	}
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	opts = append(opts, grpc.WithUnaryInterceptor(myInterceptor))
	dial, err := grpc.Dial("127.0.0.1:50052", opts...)
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
