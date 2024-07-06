package greeting

import (
	"context"
	"log"
	"path/filepath"
	"strings"
	"time"

	"github.com/kenneth-a-john/greetings-gen/client"
	greeting "github.com/kenneth-a-john/greetings-gen/greeting/proto"
	"github.com/kenneth-a-john/greetings-gen/processor"
)

func getFileNames(req *greeting.GreetingGenerationRequest) (string, string, string) {
	image_prompt := req.GetImagePrompt()
	image_prompt = strings.ReplaceAll(image_prompt, " ", "_")
	if len(image_prompt) >= 5 {
		image_prompt = image_prompt[:5]
	}
	text_prompt := req.GetMessagePrompt()
	text_prompt = strings.ReplaceAll(text_prompt, " ", "_")
	if len(text_prompt) >= 5 {
		text_prompt = text_prompt[:5]
	}
	file_name := image_prompt + "_" + text_prompt + ".png"
	originalpath := filepath.Join("..", "..", "react-client", "public", "generated_images", file_name)
	resultpath := filepath.Join("..", "..", "react-client", "public", "generated_images", file_name)
	imagepath := filepath.Join("generated_images", file_name)
	return originalpath, resultpath, imagepath
}

func GenGreeting(ctx context.Context, oac client.OpenAIClient, req *greeting.GreetingGenerationRequest) (*greeting.GreetingGenerationResponse, error) {
	log.Println("Generating image")
	resp, err := oac.GenerateImage(ctx, req.GetImagePrompt(), 1, "1024x1024")
	if err != nil {
		log.Fatal("unable to generate image: ", err)
	}
	originalpath, resultpath, imagepath := getFileNames(req)

	log.Println("Downloading image")
	wc := client.NewWebClient()
	err = wc.DownloadImageToFile(ctx, resp.Data[0].URL, originalpath)
	if err != nil {
		log.Fatal("unable to download image to file: ", err)
	}

	log.Println("Adding text to image")
	ip := processor.NewImageProcessor()
	//using same original_path and result_path to overwrite

	messagePrompt := req.GetMessagePrompt()
	if len(messagePrompt) > 32 {
		messagePrompt = messagePrompt[0:32]
	}
	err = ip.AddTextToImage(ctx, originalpath, resultpath, messagePrompt)
	if err != nil {
		log.Fatal("unable to add text to image: ", err)
	}
	time.Sleep(5 * time.Second)
	// imagepath := "generated_images/daft__human.png"
	log.Println("Done, saved at:", imagepath)
	return &greeting.GreetingGenerationResponse{
		ImagePath: imagepath,
	}, nil
}
