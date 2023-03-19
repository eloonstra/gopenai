package requests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/eloonstra/gopenai/response"
	"net/http"
)

type CreateCompletionBody struct {
	Model            string   `json:"model"`
	Prompt           string   `json:"prompt"`
	Suffix           string   `json:"suffix"`
	MaxTokens        int      `json:"max_tokens"`
	Temperature      float64  `json:"temperature"`
	TopP             float64  `json:"top_p"`
	N                int      `json:"n"`
	Stream           bool     `json:"stream"`
	LogProbs         int      `json:"logprobs"`
	Echo             bool     `json:"echo"`
	Stop             []string `json:"stop"`
	PresencePenalty  float64  `json:"presence_penalty"`
	FrequencyPenalty float64  `json:"frequency_penalty"`
	BestOf           int      `json:"best_of"`
	User             string   `json:"user"`
}

type CreateCompletion struct {
	ApiKey string
	Body   CreateCompletionBody
}

// NewCreateCompletion creates a new CreateCompletion request.
func NewCreateCompletion(apiKey, model string) *CreateCompletion {
	return &CreateCompletion{
		ApiKey: apiKey,
		Body: CreateCompletionBody{
			Model:       model,
			MaxTokens:   16,
			Temperature: 1,
			TopP:        1,
			N:           1,
			BestOf:      1,
		},
	}
}

// WithPrompt sets the prompt for the request.
func (c *CreateCompletion) WithPrompt(prompt string) *CreateCompletion {
	c.Body.Prompt = prompt
	return c
}

// WithSuffix sets the suffix for the request.
func (c *CreateCompletion) WithSuffix(suffix string) *CreateCompletion {
	c.Body.Suffix = suffix
	return c
}

// WithMaxTokens sets the max tokens for the request.
func (c *CreateCompletion) WithMaxTokens(maxTokens int) *CreateCompletion {
	c.Body.MaxTokens = maxTokens
	return c
}

// WithTemperature sets the temperature for the request.
func (c *CreateCompletion) WithTemperature(temperature float64) *CreateCompletion {
	c.Body.Temperature = temperature
	return c
}

// WithTopP sets the top p for the request.
func (c *CreateCompletion) WithTopP(topP float64) *CreateCompletion {
	c.Body.TopP = topP
	return c
}

// WithN sets the n for the request.
func (c *CreateCompletion) WithN(n int) *CreateCompletion {
	c.Body.N = n
	return c
}

// WithStream sets the stream for the request.
func (c *CreateCompletion) WithStream(stream bool) *CreateCompletion {
	c.Body.Stream = stream
	return c
}

// WithLogProbs sets the log probs for the request.
func (c *CreateCompletion) WithLogProbs(logProbs int) *CreateCompletion {
	c.Body.LogProbs = logProbs
	return c
}

// WithEcho sets the echo for the request.
func (c *CreateCompletion) WithEcho(echo bool) *CreateCompletion {
	c.Body.Echo = echo
	return c
}

// WithStop sets the stop for the request.
func (c *CreateCompletion) WithStop(stop []string) *CreateCompletion {
	c.Body.Stop = stop
	return c
}

// WithPresencePenalty sets the presence penalty for the request.
func (c *CreateCompletion) WithPresencePenalty(presencePenalty float64) *CreateCompletion {
	c.Body.PresencePenalty = presencePenalty
	return c
}

// WithFrequencyPenalty sets the frequency penalty for the request.
func (c *CreateCompletion) WithFrequencyPenalty(frequencyPenalty float64) *CreateCompletion {
	c.Body.FrequencyPenalty = frequencyPenalty
	return c
}

// WithBestOf sets the best of for the request.
func (c *CreateCompletion) WithBestOf(bestOf int) *CreateCompletion {
	c.Body.BestOf = bestOf
	return c
}

// WithUser sets the user for the request.
func (c *CreateCompletion) WithUser(user string) *CreateCompletion {
	c.Body.User = user
	return c
}

// Do sends the request to the OpenAI API and returns the response.
func (c *CreateCompletion) Do() (response.Completion, error) {
	body, err := json.Marshal(c.Body)
	req, err := http.NewRequest("POST", "https://api.openai.com/v1/completions", bytes.NewBuffer(body))
	if err != nil {
		return response.Completion{}, err
	}
	req.Header.Add("Authorization", "Bearer "+c.ApiKey)
	req.Header.Add("Content-Type", "application/json")

	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		return response.Completion{}, err
	}

	if resp.StatusCode != http.StatusOK {
		var errRes response.Error
		err = json.NewDecoder(resp.Body).Decode(&errRes)
		return response.Completion{}, fmt.Errorf("%s: %s", resp.Status, errRes.Error.Message)
	}

	var model response.Completion
	err = json.NewDecoder(resp.Body).Decode(&model)
	if err != nil {
		return response.Completion{}, err
	}

	return model, nil
}
