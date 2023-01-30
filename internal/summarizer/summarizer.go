package summarizer

import (
	"context"
	"fmt"

	gogpt "github.com/sashabaranov/go-gpt3"
	"github.com/spf13/viper"
)

func SummarizeMeeting(transcript string, apiKey string) (string, error) {
	gptClient := gogpt.NewClient(viper.GetString("OPENAI_API_KEY"))
	ctx := context.Background()

	var requestTemplate string = "Summarize the meeting using this transcript: %s"
	req := gogpt.CompletionRequest{
		Model:            gogpt.GPT3TextDavinci003,
		MaxTokens:        500,
		Prompt:           fmt.Sprintf(requestTemplate, transcript),
		Temperature:      0,
		BestOf:           1,
		PresencePenalty:  0,
		FrequencyPenalty: 0,
		TopP:             1,
	}
	resp, err := gptClient.CreateCompletion(ctx, req)
	if err != nil {
		return "", err
	}

	return resp.Choices[0].Text, nil
}