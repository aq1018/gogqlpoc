package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/PeerStreet/aqgqlpoc/graph/generated"
	"github.com/PeerStreet/aqgqlpoc/graph/model"
)

func (r *loanResolver) CreatedAt(ctx context.Context, obj *model.Loan) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *loanResolver) UpdatedAt(ctx context.Context, obj *model.Loan) (string, error) {
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

func (r *propertyResolver) CreatedAt(ctx context.Context, obj *model.Property) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *propertyResolver) UpdatedAt(ctx context.Context, obj *model.Property) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetLoans(ctx context.Context) ([]*model.Loan, error) {
	panic(fmt.Errorf("not implemented"))
}

// Loan returns generated.LoanResolver implementation.
func (r *Resolver) Loan() generated.LoanResolver { return &loanResolver{r} }

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Property returns generated.PropertyResolver implementation.
func (r *Resolver) Property() generated.PropertyResolver { return &propertyResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type loanResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type propertyResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *loanResolver) Valuation(ctx context.Context, obj *model.Loan) (string, error) {
	panic(fmt.Errorf("not implemented"))
}
func (r *loanResolver) Principal(ctx context.Context, obj *model.Loan) (string, error) {
	panic(fmt.Errorf("not implemented"))
}
func (r *loanResolver) Interest(ctx context.Context, obj *model.Loan) (string, error) {
	panic(fmt.Errorf("not implemented"))
}
func (r *loanResolver) Properties(ctx context.Context, obj *model.Loan) ([]*model.Property, error) {
	panic(fmt.Errorf("not implemented"))
}
func (r *propertyResolver) Loans(ctx context.Context, obj *model.Property) ([]*model.Loan, error) {
	panic(fmt.Errorf("not implemented"))
}
