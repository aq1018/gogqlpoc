package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/aq1018/gogqlpoc/db"
	"github.com/aq1018/gogqlpoc/graph/generated"
	"github.com/aq1018/gogqlpoc/graph/model"
	"github.com/jinzhu/gorm"
)

func (r *loanResolver) LoadProperties(ctx context.Context, obj *model.Loan) ([]*model.Property, error) {
	err := db.DB.Model(obj).Related(&obj.Properties, "Properties").Error

	return obj.Properties, err
}

func (r *mutationResolver) AddLoan(ctx context.Context, request model.LoanRequest) (*model.LoanResponse, error) {
	response := &model.LoanResponse{}

	err := db.DB.Transaction(func(tx *gorm.DB) error {
		var properties []*model.Property

		for _, address := range request.Addresses {
			property := model.NewPropertyFromAddress(address)
			if err := db.DB.FirstOrCreate(property, property).Error; err != nil {
				return err
			}
			properties = append(properties, property)
		}

		response.Loan = &model.Loan{
			Properties: properties,
			Valuation:  request.Valuation,
			Principal:  request.Principal,
			Interest:   request.Interest,
		}

		if err := db.DB.Create(response.Loan).Error; err != nil {
			return err
		}

		return nil
	})

	return response, err
}

func (r *mutationResolver) RemoveLoan(ctx context.Context, loanID int64) (*model.LoanResponse, error) {
	loan := &model.Loan{}

	if err := db.DB.First(loan, loanID).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return &model.LoanResponse{Errors: []string{"Not found"}}, nil
		}
		return nil, err
	}

	if err := db.DB.Delete(loan).Error; err != nil {
		return nil, err
	}

	return &model.LoanResponse{Loan: loan}, nil
}

func (r *mutationResolver) UpdateLoan(ctx context.Context, loanID int64, request model.LoanRequest) (*model.LoanResponse, error) {
	response := &model.LoanResponse{}

	err := db.DB.Transaction(func(tx *gorm.DB) error {
		response.Loan = &model.Loan{}
		if err := db.DB.First(&response.Loan, loanID).Error; err != nil {
			if gorm.IsRecordNotFoundError(err) {
				response.Errors = append(response.Errors, "Not Found")
				response.Loan = nil
				return nil
			}

			return err
		}

		var properties []*model.Property

		for _, address := range request.Addresses {
			property := model.NewPropertyFromAddress(address)
			if err := db.DB.FirstOrCreate(property, property).Error; err != nil {
				return err
			}
			properties = append(properties, property)
		}

		response.Loan.Valuation = request.Valuation
		response.Loan.Principal = request.Principal
		response.Loan.Interest = request.Interest

		if err := db.DB.Save(response.Loan).Error; err != nil {
			return err
		}

		return db.DB.Model(response.Loan).Association("Properties").Replace(properties).Error
	})

	return response, err
}

func (r *queryResolver) GetLoans(ctx context.Context) ([]*model.Loan, error) {
	loans := []*model.Loan{}

	if err := db.DB.Find(&loans).Error; err != nil {
		return nil, err
	}

	return loans, nil
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
