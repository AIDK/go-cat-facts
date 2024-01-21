package service

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/aidk/gocatfacts/cmd/internal/api/types"
)

// Service is the interface that provides facts.
type Service interface {
	GetFact(context.Context) (*types.Fact, error)
}

// FactService is a simple implementation of Service.
type FactService struct {
	url string
}

// NewFactService returns a new FactService.
func NewFactService(url string) Service {
	return &FactService{
		url: url,
	}
}

// GetFact makes a GET request to the fact API and returns the response.
func (s *FactService) GetFact(ctx context.Context) (*types.Fact, error) {

	// we make a GET request to the fact API
	response, err := http.Get(s.url)
	if err != nil {
		return nil, err
	}

	// sample response:
	//{"fact":"A cat's field of vision is about 200 degrees.","length":45}

	// we defer the closing of the response body so that we can
	// ensure it is closed before we return from the GetFact function
	defer response.Body.Close()

	// we initialize our Fact struct
	fact := &types.Fact{}

	// we decode the response body into our Fact struct
	if err := json.NewDecoder(response.Body).Decode(fact); err != nil {
		return nil, err
	}

	return fact, nil
}
