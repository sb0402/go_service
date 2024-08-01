package main

import (
	"context"
	"encoding/json"
	"log"
	"net"

	"github.com/osteele/liquid"
	pb "github.com/sb0402/go_service/liquidparsing/liquidparsingpb"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedLiquidParsingGrpcServer
}

func (s *server) ParseAndRenderString(ctx context.Context, in *pb.Data) (*pb.ParsedResult, error) {
	engine := liquid.NewEngine()
	var result string
	var err error
	template := make(map[string]interface{})
	for key, value := range in.Variables {
		if key == "html_part" {
			template[key] = value // Directly assigning the HTML part
			continue
		}
		var jsonData map[string]interface{}
		// Attempt to unmarshal the value as JSON
		if err := json.Unmarshal([]byte(value), &jsonData); err == nil {
			template[key] = jsonData
		}
	}
	html := template["html_part"].(string)
	for i := 0; i < 123000; i++ {
		result, err = engine.ParseAndRenderString(html, template)
		if err != nil {
			log.Fatal(err)
		}
	}
	return &pb.ParsedResult{Result: result}, nil

}

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:5005")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	log.Printf("Server listening at %v", lis.Addr())

	s := grpc.NewServer()
	pb.RegisterLiquidParsingGrpcServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
