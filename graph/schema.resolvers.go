package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.43

import (
	"context"
	"fmt"

	"github.com/psbernardo/graphql/graph/model"
)

// SaveSearchResult is the resolver for the saveSearchResult field.
func (r *mutationResolver) SaveSearchResult(ctx context.Context, input model.SearchResult) (*model.SaveResultResponse, error) {
	panic(fmt.Errorf("not implemented: SaveSearchResult - saveSearchResult"))
}

// Search is the resolver for the search field.
func (r *queryResolver) Search(ctx context.Context, filter string) ([]*model.People, error) {
	panic(fmt.Errorf("not implemented: Search - search"))
}

// GetSearchHistory is the resolver for the getSearchHistory field.
func (r *queryResolver) GetSearchHistory(ctx context.Context) ([]string, error) {
	panic(fmt.Errorf("not implemented: GetSearchHistory - getSearchHistory"))
}

// GetSaveResult is the resolver for the getSaveResult field.
func (r *queryResolver) GetSaveResult(ctx context.Context, key string) ([]*model.People, error) {
	panic(fmt.Errorf("not implemented: GetSaveResult - getSaveResult"))
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
