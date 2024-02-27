package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.44

import (
	"context"
	"crypto/rand"
	"fmt"
	"math/big"

	"github.com/BurrrY/obstwiesen-server/graph/model"
)

// CreateMeadow is the resolver for the createMeadow field.
func (r *mutationResolver) CreateMeadow(ctx context.Context, input model.NewMeadow) (*model.Meadow, error) {
	randNumber, _ := rand.Int(rand.Reader, big.NewInt(100))
	meadow := &model.Meadow{
		ID:    fmt.Sprintf("T%d", randNumber),
		Name:  input.Name,
		Trees: nil,
	}
	storage.StoreMeadow(meadow)
	r.meadows = append(r.meadows, meadow)
	return meadow, nil
}

// CreateTree is the resolver for the createTree field.
func (r *mutationResolver) CreateTree(ctx context.Context, input model.NewTree) (*model.Tree, error) {
	randNumber, _ := rand.Int(rand.Reader, big.NewInt(100))

	tree := &model.Tree{
		ID:   fmt.Sprintf("T%d", randNumber),
		Name: input.Name,
	}

	r.meadows[0].Trees = append(r.meadows[0].Trees, tree)
	return tree, nil
}

// Meadows is the resolver for the meadows field.
func (r *queryResolver) Meadows(ctx context.Context) ([]*model.Meadow, error) {
	return r.meadows, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
