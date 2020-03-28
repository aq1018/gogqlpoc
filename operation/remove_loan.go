package operation

import (
	"context"

	"github.com/aq1018/gogqlpoc/graph/model"
	"github.com/jinzhu/gorm"
)

func (op *Operation) RemoveLoan(ctx context.Context, loanID int64) (*model.LoanResponse, error) {
	loan := &model.Loan{}

	if err := op.DB.First(loan, loanID).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return &model.LoanResponse{Errors: []string{"Not found"}}, nil
		}
		return nil, err
	}

	if err := op.DB.Delete(loan).Error; err != nil {
		return nil, err
	}

	return &model.LoanResponse{Loan: loan}, nil
}
