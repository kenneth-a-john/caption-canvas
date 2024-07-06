package client

import (
	"context"

	"github.com/go-resty/resty/v2"
	"github.com/kenneth-a-john/greetings-gen/model"
)

type OpenAIClient interface {
	GenerateImage(ctx context.Context, prompt string, n uint8, size string) (model.ImageGenerationResponse, error)
}

type openAIClient struct {
	*resty.Client
}

type mockOpenAIClient struct{}

func (m *mockOpenAIClient) GenerateImage(context.Context, string, uint8, string) (model.ImageGenerationResponse, error) {
	return model.ImageGenerationResponse{
		Data: []model.ImageGenerationData{
			{
				URL: "https://oaidalleapiprodscus.blob.core.windows.net/private/org-wmFC2rGGykmuzJ9S5y6N9iku/user-A2OLbPg26dTIPUxfZSt7t3Ty/img-rVzdLgMawV9TsFYsNqnq85xn.png?st=2024-06-07T03%3A15%3A36Z&se=2024-06-07T05%3A15%3A36Z&sp=r&sv=2023-11-03&sr=b&rscd=inline&rsct=image/png&skoid=6aaadede-4fb3-4698-a8f6-684d7786b067&sktid=a48cca56-e6da-484e-a814-9c849652bcb3&skt=2024-06-06T17%3A12%3A55Z&ske=2024-06-07T17%3A12%3A55Z&sks=b&skv=2023-11-03&sig=4hAU3bMkVP6HXqNtGj%2BbGXhQriau24u9B2FzzlF%2Ba70%3D",
			},
		},
	}, nil
}

func (o *openAIClient) GenerateImage(ctx context.Context, prompt string, n uint8, size string) (model.ImageGenerationResponse, error) {
	var success model.ImageGenerationResponse
	var fail model.OpenAIErrorResponse

	resp, err := o.R().
		SetContext(ctx).
		SetBody(model.ImageGenerationRequest{
			Model:  "dall-e-3",
			Prompt: prompt,
			N:      n,
			Size:   size,
		}).
		SetResult(&success).
		SetError(&fail).
		Post("/v1/images/generations")

	if err != nil {
		return model.ImageGenerationResponse{}, err
	}

	if resp.IsError() {
		return model.ImageGenerationResponse{}, &model.RESTClientError{
			Message:  "OpenAI error",
			Response: fail,
		}
	}

	return success, nil
}

func NewOpenAIClient(baseURL string, token string) OpenAIClient {
	client := resty.New()
	client.SetBaseURL(baseURL)
	client.SetAuthScheme("Bearer")
	client.SetAuthToken(token)

	return &openAIClient{client}
}

func NewMockOpenAIClient() OpenAIClient {
	return &mockOpenAIClient{}
}
