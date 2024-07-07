package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/kenneth-a-john/greetings-gen/client"
	"github.com/kenneth-a-john/greetings-gen/config"
	greet "github.com/kenneth-a-john/greetings-gen/greeting"
	greeting "github.com/kenneth-a-john/greetings-gen/greeting/proto"
)

var oac client.OpenAIClient

type myGreetingServer struct {
	greeting.UnimplementedImageServiceServer
}

func (s myGreetingServer) GenerateGreeting(ctx context.Context, req *greeting.GreetingGenerationRequest) (*greeting.GreetingGenerationResponse, error) {
	log.Println(req.ImagePrompt, req.MessagePrompt)
	return greet.GenGreeting(ctx, oac, req)
}

// use struct to represent the data
// to recieve and send
type ImageReq struct {
	ImagePrompt   string `json:"image_prompt"`
	MessagePrompt string `json:"message_prompt"`
}

func main() {
	log.Println("Starting caption canvas go-backend...")
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
	ctx := context.Background()
	mux := http.NewServeMux()
	log.Println("starting server...")

	mux.HandleFunc("/generate", func(w http.ResponseWriter, r *http.Request) {
		imageReq := &ImageReq{}
		err := json.NewDecoder(r.Body).Decode(imageReq)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		fmt.Println("got imageReq:", imageReq)

		req := &greeting.GreetingGenerationRequest{
			MessagePrompt: imageReq.MessagePrompt,
			ImagePrompt:   imageReq.ImagePrompt,
		}

		response, err := greet.GenGreeting(ctx, oac, req)
		if err != nil {
			log.Fatal("Unable to generate: ", err)
		}
		resp, err := json.Marshal(response)
		w.Write(resp)
	})

	if err := http.ListenAndServe(":8089", mux); err != http.ErrServerClosed {
		panic(err)
	}
	// lis, err := net.Listen("tcp", ":8089")
	// if err != nil {
	// 	log.Fatalf("unable to create listener: %s", err)
	// }

	// serverRegistrar := grpc.NewServer()
	// service := &myGreetingServer{}
	// greeting.RegisterImageServiceServer(serverRegistrar, service)

	// err = serverRegistrar.Serve(lis)
	// if err != nil {
	// 	log.Fatalf("unable to serve: %s", err)
	// }

}
