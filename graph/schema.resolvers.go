package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/PeerStreet/aqgqlpoc/db"
	"github.com/PeerStreet/aqgqlpoc/graph/generated"
	"github.com/PeerStreet/aqgqlpoc/graph/model"
)

func (r *mutationResolver) AddLoan(ctx context.Context, request model.LoanRequest) (*model.LoanResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) RemoveLoan(ctx context.Context, loanID int64) (*model.LoanResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateLoan(ctx context.Context, loanID int64, request model.LoanRequest) (*model.LoanResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetLoans(ctx context.Context) ([]*model.Loan, error) {
	loans := []*model.Loan{}

	db.DB.Find(&loans)

	return loans, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
