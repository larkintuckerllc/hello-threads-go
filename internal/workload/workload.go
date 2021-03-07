package workload

import (
	"fmt"
	"log"
	"net"
	"os"

	greet "github.com/larkintuckerllc/hellothreadsgo/internal/helloworld"
	helloworldPb "github.com/larkintuckerllc/hellothreadsgo/pkg/helloworld"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func Execute() {
	var port string
	envPort := os.Getenv("PORT")
	if envPort == "" {
		port = ":50051"
	} else {
		port = fmt.Sprintf(":%s", envPort)
	}
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	helloworldPb.RegisterGreeterServer(s, &greet.GreeterServer{})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
