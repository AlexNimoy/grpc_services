package main

import (
	"context"
	"log"
	"net"

	pb "server/proto"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedMessageServiceServer
}

func (s *server) SendMessage(ctx context.Context, in *pb.Message) (*pb.Response, error) {
	log.Printf("Received message: %s", in.Text)
	return &pb.Response{Reply: "Message received"}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to open port: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterMessageServiceServer(s, &server{})
	log.Println("Server is running on port 50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to start gRPC server: %v", err)
	}
}
