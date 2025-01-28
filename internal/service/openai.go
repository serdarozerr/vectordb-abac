package service

import (
	"context"
	"github.com/openai/openai-go"
)
import "github.com/openai/openai-go/option"

type LLM struct {
	client *openai.Client
}

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
