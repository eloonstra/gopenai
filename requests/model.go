package requests

import (
	"encoding/json"
	"fmt"
	"github.com/eloonstra/gopenai/response"
	"net/http"
)

type ListModels struct {
	ApiKey string
}

// NewListModels creates a new ListModels request.
func NewListModels(apiKey string) *ListModels {
	return &ListModels{
		ApiKey: apiKey,
	}
}

// Do sends the request to the OpenAI API and returns the response.
func (l *ListModels) Do() (response.Models, error) {
	req, err := http.NewRequest("GET", "https://api.openai.com/v1/models", nil)
	if err != nil {
		return response.Models{}, err
	}
	req.Header.Add("Authorization", "Bearer "+l.ApiKey)
	req.Header.Add("Content-Type", "application/json")

	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		return response.Models{}, err
	}

	var models response.Models
	err = json.NewDecoder(resp.Body).Decode(&models)
	if err != nil {
		return response.Models{}, err
	}

	return models, nil
}

type RetrieveModel struct {
	ApiKey string
	Id     string
}

// NewRetrieveModel creates a new RetrieveModel request.
func NewRetrieveModel(apiKey, id string) *RetrieveModel {
	return &RetrieveModel{
		ApiKey: apiKey,
		Id:     id,
	}
}

// Do sends the request to the OpenAI API and returns the response.
func (r *RetrieveModel) Do() (response.Model, error) {
	req, err := http.NewRequest("GET", "https://api.openai.com/v1/models/"+r.Id, nil)
	if err != nil {
		return response.Model{}, err
	}
	req.Header.Add("Authorization", "Bearer "+r.ApiKey)
	req.Header.Add("Content-Type", "application/json")

	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		return response.Model{}, err
	}

	if resp.StatusCode != http.StatusOK {
		var errRes response.Error
		err = json.NewDecoder(resp.Body).Decode(&errRes)
		return response.Model{}, fmt.Errorf("%s: %s", resp.Status, errRes.Error.Message)
	}

	var model response.Model
	err = json.NewDecoder(resp.Body).Decode(&model)
	if err != nil {
		return response.Model{}, err
	}

	return model, nil
}
