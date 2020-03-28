package operation

import (
	"context"

	"github.com/aq1018/gogqlpoc/graph/model"
)

func (op *Operation) GetLoans(ctx context.Context) ([]*model.Loan, error) {
	loans := []*model.Loan{}

	if err := op.DB.Find(&loans).Error; err != nil {
		return nil, err
	}

	return loans, nil
}
