package gopenai

import (
	"github.com/eloonstra/gopenai/requests"
)

type OpenAI struct {
	ApiKey string
}

func New(apiKey string) *OpenAI {
	return &OpenAI{
		ApiKey: apiKey,
	}
}

func (o *OpenAI) CreateCompletion(model string) *requests.CreateCompletion {
	return requests.NewCreateCompletion(o.ApiKey, model)
}

func (o *OpenAI) CreateChatCompletion(model string, messages []requests.Message) *requests.CreateChatCompletion {
	return requests.NewCreateChatCompletion(o.ApiKey, model, messages)
}

func (o *OpenAI) ListModels() *requests.ListModels {
	return requests.NewListModels(o.ApiKey)
}

func (o *OpenAI) RetrieveModel(modelId string) *requests.RetrieveModel {
	return requests.NewRetrieveModel(o.ApiKey, modelId)
}

func (o *OpenAI) CreateEdit(model, instruction string) *requests.CreateEdit {
	return requests.NewCreateEdit(o.ApiKey, model, instruction)
}

func (o *OpenAI) CreateModeration(input []string) *requests.CreateModeration {
	return requests.NewCreateModeration(o.ApiKey, input)
}
