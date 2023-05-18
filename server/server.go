package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	pb "github.com/azusachino/maackia/protobuf/hello"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50052, "server port")
)

type helloServer struct {
	pb.UnimplementedHelloServer
}

func (s *helloServer) SayHello(ctx context.Context, req *pb.HelloReq) (*pb.HelloRes, error) {
	r := pb.HelloRes{
		Code: 200,
	}
	r.Message = fmt.Sprintf("Hello %s", req.GetName())
	return &r, nil
}

func newServer() *helloServer {
	s := &helloServer{}
	return s
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("127.0.0.1:%d", *port))
	if err != nil {
		log.Fatalf("fail to listen %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterHelloServer(grpcServer, newServer())
	grpcServer.Serve(lis)
}
