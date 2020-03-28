package operation

import (
	"context"

	"github.com/aq1018/gogqlpoc/graph/model"
	"github.com/gocraft/work"
	"github.com/jinzhu/gorm"
)

func (op *Operation) AddLoan(ctx context.Context, request model.LoanRequest) (*model.LoanResponse, error) {
	response := &model.LoanResponse{}

	err := op.DB.Transaction(func(tx *gorm.DB) error {
		var properties []*model.Property

		for _, address := range request.Addresses {
			property := model.NewPropertyFromAddress(address)
			if err := op.DB.FirstOrCreate(property, property).Error; err != nil {
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

		if err := op.DB.Create(response.Loan).Error; err != nil {
			return err
		}

		op.Worker.Enqueue("send_email", work.Q{"loanID": response.Loan.ID})

		return nil
	})

	return response, err
}
