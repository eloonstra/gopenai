package requests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/eloonstra/gopenai/response"
	"net/http"
)

type CreateEditBody struct {
	Model       string  `json:"model"`
	Input       string  `json:"input"`
	Instruction string  `json:"instruction"`
	N           int     `json:"n"`
	Temperature float64 `json:"temperature"`
	TopP        float64 `json:"top_p"`
}

type CreateEdit struct {
	ApiKey string
	Body   CreateEditBody
}

// NewCreateEdit creates a new CreateEdit request.
func NewCreateEdit(apiKey, model, instruction string) *CreateEdit {
	return &CreateEdit{
		ApiKey: apiKey,
		Body: CreateEditBody{
			Model:       model,
			Instruction: instruction,
			N:           1,
			Temperature: 1,
			TopP:        1,
		},
	}
}

// WithInput sets the input for the request.
func (c *CreateEdit) WithInput(input string) *CreateEdit {
	c.Body.Input = input
	return c
}

// WithN sets the number of results to return.
func (c *CreateEdit) WithN(n int) *CreateEdit {
	c.Body.N = n
	return c
}

// WithTemperature sets the temperature for the request.
func (c *CreateEdit) WithTemperature(temperature float64) *CreateEdit {
	c.Body.Temperature = temperature
	return c
}

// WithTopP sets the top_p for the request.
func (c *CreateEdit) WithTopP(topP float64) *CreateEdit {
	c.Body.TopP = topP
	return c
}

// Do sends the request.
func (c *CreateEdit) Do() (response.Edit, error) {
	body, err := json.Marshal(c.Body)
	req, err := http.NewRequest("POST", "https://api.openai.com/v1/edits", bytes.NewBuffer(body))
	if err != nil {
		return response.Edit{}, err
	}
	req.Header.Add("Authorization", "Bearer "+c.ApiKey)
	req.Header.Add("Content-Type", "application/json")

	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		return response.Edit{}, err
	}

	if resp.StatusCode != http.StatusOK {
		var errRes response.Error
		err = json.NewDecoder(resp.Body).Decode(&errRes)
		return response.Edit{}, fmt.Errorf("%s: %s", resp.Status, errRes.Error.Message)
	}

	var model response.Edit
	err = json.NewDecoder(resp.Body).Decode(&model)
	if err != nil {
		return response.Edit{}, err
	}

	return model, nil
}
