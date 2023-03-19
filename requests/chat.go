package requests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/eloonstra/gopenai/response"
	"net/http"
)

type Role string

const (
	RoleSystem    Role = "system"
	RoleUser      Role = "user"
	RoleAssistant Role = "assistant"
)

type Message struct {
	Role    Role   `json:"role"`
	Content string `json:"content"`
}

type CreateChatCompletionBody struct {
	Model            string    `json:"model"`
	Messages         []Message `json:"messages"`
	Temperature      float64   `json:"temperature"`
	TopP             float64   `json:"top_p"`
	N                int       `json:"n"`
	Stream           bool      `json:"stream"`
	Stop             []string  `json:"stop"`
	MaxTokens        int       `json:"max_tokens"`
	PresencePenalty  float64   `json:"presence_penalty"`
	FrequencyPenalty float64   `json:"frequency_penalty"`
	User             string    `json:"user"`
}

type CreateChatCompletion struct {
	ApiKey string
	Body   CreateChatCompletionBody
}

// NewCreateChatCompletion creates a new CreateChatCompletion request.
func NewCreateChatCompletion(apiKey, model string, messages []Message) *CreateChatCompletion {
	return &CreateChatCompletion{
		ApiKey: apiKey,
		Body: CreateChatCompletionBody{
			Model:       model,
			Messages:    messages,
			Temperature: 1,
			TopP:        1,
			N:           1,
			MaxTokens:   1000,
		},
	}
}

// WithModel sets the model for the request.
func (c *CreateChatCompletion) WithModel(model string) *CreateChatCompletion {
	c.Body.Model = model
	return c
}

// WithMessages sets the messages for the request.
func (c *CreateChatCompletion) WithMessages(messages []Message) *CreateChatCompletion {
	c.Body.Messages = messages
	return c
}

// WithTemperature sets the temperature for the request.
func (c *CreateChatCompletion) WithTemperature(temperature float64) *CreateChatCompletion {
	c.Body.Temperature = temperature
	return c
}

// WithTopP sets the top_p for the request.
func (c *CreateChatCompletion) WithTopP(topP float64) *CreateChatCompletion {
	c.Body.TopP = topP
	return c
}

// WithN sets the number of results to return.
func (c *CreateChatCompletion) WithN(n int) *CreateChatCompletion {
	c.Body.N = n
	return c
}

// Do sends the request to the OpenAI API and returns the response.
func (c *CreateChatCompletion) Do() (response.Chat, error) {
	body, err := json.Marshal(c.Body)
	req, err := http.NewRequest("POST", "https://api.openai.com/v1/chat/completions", bytes.NewBuffer(body))
	if err != nil {
		return response.Chat{}, err
	}
	req.Header.Add("Authorization", "Bearer "+c.ApiKey)
	req.Header.Add("Content-Type", "application/json")

	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		return response.Chat{}, err
	}

	if resp.StatusCode != http.StatusOK {
		var errRes response.Error
		err = json.NewDecoder(resp.Body).Decode(&errRes)
		return response.Chat{}, fmt.Errorf("%s: %s", resp.Status, errRes.Error.Message)
	}

	var model response.Chat
	err = json.NewDecoder(resp.Body).Decode(&model)
	if err != nil {
		return response.Chat{}, err
	}

	return model, nil
}
