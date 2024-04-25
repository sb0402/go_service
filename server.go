package main

import (
	"context"
	"log"
	"net"

	"github.com/osteele/liquid"
	"google.golang.org/grpc"

	pb "liquid" // Update with your Go package name
)

type server struct{}

func (s *server) ParseAndRenderString(ctx context.Context, in *pb.Template) (*pb.ParsedResult, error) {
	engine := liquid.NewEngine()
	result, err := engine.ParseAndRenderString(in.HtmlPart, in.Variables)
	if err != nil {
		return nil, err
	}
	return &pb.ParsedResult{Result: result}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":50051") // Use any available port
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterLiquidParsingServer(s, &server{})
	if err := s.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
