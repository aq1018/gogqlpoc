// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type MutationResponse interface {
	IsMutationResponse()
}

type Address struct {
	Address1 string `json:"address1"`
	Address2 string `json:"address2"`
	City     string `json:"city"`
	State    string `json:"state"`
	Zip      string `json:"zip"`
}

type LoanRequest struct {
	Addresses []*Address `json:"addresses"`
	Valuation string     `json:"valuation"`
	Principal string     `json:"principal"`
	Interest  string     `json:"interest"`
}

type LoanResponse struct {
	Errors []string `json:"errors"`
	Loan   *Loan    `json:"loan"`
}

func (LoanResponse) IsMutationResponse() {}
