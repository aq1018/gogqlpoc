package model

type Loan struct {
	ID          string `json:"id"`
	PropertyIDs string `json:"properties"`
	Valuation   string `json:"valuation"`
	Principal   string `json:"principal"`
	Interest    string `json:"interest"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
}
