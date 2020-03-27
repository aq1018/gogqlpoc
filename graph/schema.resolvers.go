package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/PeerStreet/aqgqlpoc/graph/generated"
	"github.com/PeerStreet/aqgqlpoc/graph/model"
)

func (r *loanResolver) Properties(ctx context.Context, obj *model.Loan) ([]*model.Property, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) AddLoan(ctx context.Context, request model.LoanRequest) (*model.LoanResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) RemoveLoan(ctx context.Context, loanID string) (*model.LoanResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateLoan(ctx context.Context, loanID string, request model.LoanRequest) (*model.LoanResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetLoans(ctx context.Context) ([]*model.Loan, error) {
	panic(fmt.Errorf("not implemented"))
}

// Loan returns generated.LoanResolver implementation.
func (r *Resolver) Loan() generated.LoanResolver { return &loanResolver{r} }

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type loanResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
