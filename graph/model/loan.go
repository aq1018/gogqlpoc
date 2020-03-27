package model

import (
	"time"

	"github.com/shopspring/decimal"
)

type Loan struct {
	ID         int64           `json:"id" gorm:"primary_key;type: serial;"`
	Properties []Property      `json:"properties" gorm:"many2many:loan_properties"`
	Valuation  decimal.Decimal `json:"valuation" sql:"type:decimal(16,2)"`
	Principal  decimal.Decimal `json:"principal" sql:"type:decimal(16,2)"`
	Interest   decimal.Decimal `json:"interest" sql:"type:decimal(8,6)"`
	CreatedAt  time.Time       `json:"createdAt"`
	UpdatedAt  time.Time       `json:"updatedAt"`
}
