package main

import (
	"context"
	"log"
	"net"

	"github.com/osteele/liquid"
	pb "github.com/sb0402/go_service/liquidparsing/liquidparsingpb"
	"google.golang.org/grpc"
)

type server struct{}

func (s *server) ParseAndRenderString(ctx context.Context, in *pb.Template) (*pb.ParsedResult, error) {
	engine := liquid.NewEngine()
	template := make(map[string]interface{})
	for key, value := range in.Variables {
		template[key] = value
	}
	html := template["html_part"].(string)
	delete(template, "html_part")
	result, err := engine.ParseAndRenderString(html, template)
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
