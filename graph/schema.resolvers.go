package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/aq1018/gogqlpoc/graph/generated"
	"github.com/aq1018/gogqlpoc/graph/model"
	"github.com/aq1018/gogqlpoc/operation"
)

func (r *loanResolver) LoadProperties(ctx context.Context, obj *model.Loan) ([]*model.Property, error) {
	return operation.ForContext(ctx).LoadLoanProperties(ctx, obj)
}

func (r *mutationResolver) AddLoan(ctx context.Context, request model.LoanRequest) (*model.LoanResponse, error) {
	return operation.ForContext(ctx).AddLoan(ctx, request)
}

func (r *mutationResolver) RemoveLoan(ctx context.Context, loanID int64) (*model.LoanResponse, error) {
	return operation.ForContext(ctx).RemoveLoan(ctx, loanID)
}

func (r *mutationResolver) UpdateLoan(ctx context.Context, loanID int64, request model.LoanRequest) (*model.LoanResponse, error) {
	return operation.ForContext(ctx).UpdateLoan(ctx, loanID, request)
}

func (r *queryResolver) GetLoans(ctx context.Context) ([]*model.Loan, error) {
	return operation.ForContext(ctx).GetLoans(ctx)
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
