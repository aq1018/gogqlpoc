package model

import (
	"time"
)

type Property struct {
	ID        string    `json:"id" gorm:"primary_key"`
	Address1  string    `json:"address1"`
	Address2  string    `json:"address2"`
	City      string    `json:"city"`
	State     string    `json:"state"`
	Zip       string    `json:"zip"`
	Loans     []Loan    `json:"loans" gorm:"many2many:loan_properties"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
