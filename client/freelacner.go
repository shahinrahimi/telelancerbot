package client

import (
	"context"
	"fmt"
	"log"

	"github.com/shahinrahimi/go-freelancer-sdk/v1"
)

type Freelancer struct {
	client *freelancer.Client
	l      *log.Logger
}

func New(l *log.Logger, token string) *Freelancer {
	c := freelancer.NewClient(token)
	return &Freelancer{client: c, l: l}
}

func (f *Freelancer) GetCountries(ctx context.Context) ([]freelancer.Country, error) {
	res, err := f.client.NewListCountriesService().Do(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get countries: %v", err)
	}
	return res.Result.Countries, nil
}

func (f *Freelancer) GetCategories(ctx context.Context) ([]freelancer.Category, error) {
	res, err := f.client.NewListCategoriesService().Do(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get categories: %v", err)
	}
	return res.Result.Categories, nil
}
