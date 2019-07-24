package main

import (
	"net"
	"os"

	pb "github.com/jspc/sine-service/service"
	"google.golang.org/grpc"
)

func main() {
	redis, err := NewRedis(os.Getenv("REDIS_MASTER"))
	if err != nil {
		panic(err)
	}

	service := Service{
		redis: redis,
	}

	lis, err := net.Listen("tcp", "0.0.0.0:8080")
	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterSineServiceServer(grpcServer, service)

	grpcServer.Serve(lis)
}
