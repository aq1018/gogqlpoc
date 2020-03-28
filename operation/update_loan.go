package operation

import (
	"context"

	"github.com/aq1018/gogqlpoc/graph/model"
	"github.com/jinzhu/gorm"
)

func (op *Operation) UpdateLoan(ctx context.Context, loanID int64, request model.LoanRequest) (*model.LoanResponse, error) {
	response := &model.LoanResponse{}

	response.Errors = request.Validate()
	if response.Errors != nil {
		return response, nil
	}

	err := op.DB.Transaction(func(tx *gorm.DB) error {
		response.Loan = &model.Loan{}
		if err := op.DB.First(&response.Loan, loanID).Error; err != nil {
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
			if err := op.DB.FirstOrCreate(property, property).Error; err != nil {
				return err
			}
			properties = append(properties, property)
		}

		response.Loan.Valuation = request.Valuation
		response.Loan.Principal = request.Principal
		response.Loan.Interest = request.Interest

		if err := op.DB.Save(response.Loan).Error; err != nil {
			return err
		}

		return op.DB.Model(response.Loan).Association("Properties").Replace(properties).Error
	})

	return response, err
}
