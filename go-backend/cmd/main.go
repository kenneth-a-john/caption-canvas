package main

import (
	"context"
	"log"
	"net"

	"github.com/kenneth-a-john/greetings-gen/client"
	"github.com/kenneth-a-john/greetings-gen/config"
	greet "github.com/kenneth-a-john/greetings-gen/greeting"
	greeting "github.com/kenneth-a-john/greetings-gen/greeting/proto"
	"google.golang.org/grpc"
)

var oac client.OpenAIClient

type myGreetingServer struct {
	greeting.UnimplementedImageServiceServer
}

func (s myGreetingServer) GenerateGreeting(ctx context.Context, req *greeting.GreetingGenerationRequest) (*greeting.GreetingGenerationResponse, error) {
	log.Println(req.ImagePrompt, req.MessagePrompt)
	return greet.GenGreeting(ctx, oac, req)
}

func main() {
	log.Println("Starting greetings generator...")
	lis, err := net.Listen("tcp", ":8089")
	if err != nil {
		log.Fatalf("unable to create listener: %s", err)
	}

	serverRegistrar := grpc.NewServer()
	service := &myGreetingServer{}
	greeting.RegisterImageServiceServer(serverRegistrar, service)

	log.Println("Loading configuration")
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("unable to load config: ", err)
	}

	if cfg.Environment == "development" {
		oac = client.NewMockOpenAIClient()
	} else {
		oac = client.NewOpenAIClient(cfg.OpenAIBaseURL, cfg.OpenAIToken)
	}

	err = serverRegistrar.Serve(lis)
	if err != nil {
		log.Fatalf("unable to serve: %s", err)
	}

}
