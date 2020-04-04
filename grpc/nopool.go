package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"

	"GoToy/grpc/pb"
)

const (
	port = ":50051"
)

type Server struct{
}

func (s *Server) GetHelloWorld(ctx context.Context, in *pb.GetHelloWorldRequest) (*pb.GetHelloWorldResponse, error) {
	// now := time.Now()
	log.Printf("Received: %v", in.GetName())
	return &pb.GetHelloWorldResponse{
		Name: in.Name,
		Now: nil,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to lister: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterTestServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}