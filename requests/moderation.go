package requests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/eloonstra/gopenai/response"
	"net/http"
)

type CreateModerationBody struct {
	Input []string `json:"input"`
	Model string   `json:"model"`
}

type CreateModeration struct {
	ApiKey string
	Body   CreateModerationBody
}

// NewCreateModeration creates a new CreateModeration request.
func NewCreateModeration(apiKey string, input []string) *CreateModeration {
	return &CreateModeration{
		ApiKey: apiKey,
		Body: CreateModerationBody{
			Input: input,
			Model: "text-moderation-latest",
		},
	}
}

// WithModel sets the model for the request.
func (c *CreateModeration) WithModel(model string) *CreateModeration {
	c.Body.Model = model
	return c
}

// Do sends the request to the OpenAI API and returns the response.
func (c *CreateModeration) Do() (response.Moderation, error) {
	body, err := json.Marshal(c.Body)
	req, err := http.NewRequest("POST", "https://api.openai.com/v1/moderations", bytes.NewBuffer(body))
	if err != nil {
		return response.Moderation{}, err
	}
	req.Header.Add("Authorization", "Bearer "+c.ApiKey)
	req.Header.Add("Content-Type", "application/json")

	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		return response.Moderation{}, err
	}

	if resp.StatusCode != http.StatusOK {
		var errRes response.Error
		err = json.NewDecoder(resp.Body).Decode(&errRes)
		return response.Moderation{}, fmt.Errorf("%s: %s", resp.Status, errRes.Error.Message)
	}

	var model response.Moderation
	err = json.NewDecoder(resp.Body).Decode(&model)
	if err != nil {
		return response.Moderation{}, err
	}

	return model, nil
}
