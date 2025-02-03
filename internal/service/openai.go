package service

import (
	"context"
	"errors"
	"github.com/openai/openai-go"
)
import "github.com/openai/openai-go/option"

type LLM struct {
	client *openai.Client
}

// Define the structure of the JSON response
type Response struct {
	Index   int     `json:"index"`
	Message Message `json:"message"`
}

type Message struct {
	Role    string  `json:"role"`
	Content string  `json:"content"`
	Refusal *string `json:"refusal"` // Nullable field
}

const systemPrompt = `
You are a helpful AI assistant specialized in answering questions and providing insights in various domains, including programming, data science, and general knowledge.
You always provide structured, accurate, and concise responses.
`

func NewLLM(key string) *LLM {
	client := openai.NewClient(option.WithAPIKey(key))
	llm := &LLM{client: client}
	return llm
}

func (l *LLM) EmbedText(ctx context.Context, text openai.EmbeddingNewParamsInputArrayOfStrings) ([]openai.Embedding, error) {
	enp := openai.EmbeddingNewParams{Input: openai.F[openai.EmbeddingNewParamsInputUnion](text),
		Model: openai.F[openai.EmbeddingModel](openai.EmbeddingModelTextEmbeddingAda002)}

	response, err := l.client.Embeddings.New(ctx, enp)
	if err != nil {
		return nil, err
	}

	return response.Data, nil
}

func (l *LLM) CompleteText(ctx context.Context, userText string) (string, error) {
	model := openai.F(openai.ChatModelGPT4o)

	messages := openai.F([]openai.ChatCompletionMessageParamUnion{
		openai.ChatCompletionSystemMessageParam{
			Role: openai.F(openai.ChatCompletionSystemMessageParamRoleSystem),
			Content: openai.F([]openai.ChatCompletionContentPartTextParam{{
				Text: openai.F(systemPrompt),
				Type: openai.F(openai.ChatCompletionContentPartTextTypeText)}})},
		openai.ChatCompletionUserMessageParam{
			Role: openai.F(openai.ChatCompletionUserMessageParamRoleUser),
			Content: openai.F([]openai.ChatCompletionContentPartUnionParam{
				openai.ChatCompletionContentPartTextParam{
					Text: openai.F(userText),
					Type: openai.F(openai.ChatCompletionContentPartTextTypeText)}})}},
	)

	response, err := l.client.Chat.Completions.New(ctx, openai.ChatCompletionNewParams{Model: model, Messages: messages})
	if response == nil {
		return "", errors.New("OpenAI response is nil")
	}

	return response.Choices[0].Message.Content, err
}
