package operation

import (
	"context"

	"github.com/aq1018/gogqlpoc/graph/model"
)

func (op *Operation) LoadLoanProperties(ctx context.Context, loan *model.Loan) ([]*model.Property, error) {
	err := op.DB.Model(loan).Related(&loan.Properties, "Properties").Error

	return loan.Properties, err
}
