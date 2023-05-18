package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	pb "github.com/azusachino/maackia/protobuf/hello"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	serverAddr = flag.String("addr", "localhost:50052", "grpc server addr")
)

func main() {
	flag.Parse()

	conn, err := grpc.Dial(*serverAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()

	client := pb.NewHelloClient(conn)
	ctx := context.TODO()
	res, err := client.SayHello(ctx, &pb.HelloReq{Name: "alice"})

	if err != nil {
		log.Fatalf("fail to grpc: %v", err)
	}

	fmt.Printf("%s\n", res.Message)
}
