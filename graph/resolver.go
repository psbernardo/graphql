package graph

import (
	"context"

	"github.com/psbernardo/graphql/graph/model"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type APIService interface {
	Search(ctx context.Context, filter string) (*model.QueryResult, error)
	PageResult(ctx context.Context, url string) (*model.QueryResult, error)
}

type Repository interface {
	SaveSearchResult(ctx context.Context, key string, searchResult *model.QueryResult) error
	GetSaveResult(ctx context.Context, key string) (*model.QueryResult, error)
	GetSearchHistoryKey(ctx context.Context) ([]string, error)
}

type Resolver struct {
	apiService APIService
	repository Repository
}

func NewResolver(apiService APIService, repository Repository) *Resolver {
	return &Resolver{
		apiService: apiService,
		repository: repository,
	}
}
