package model

type Property struct {
	ID        string   `json:"id"`
	Address1  string   `json:"address1"`
	Address2  string   `json:"address2"`
	City      string   `json:"city"`
	State     string   `json:"state"`
	Zip       string   `json:"zip"`
	LoanIDs   []string `json:"loans"`
	CreatedAt string   `json:"createdAt"`
	UpdatedAt string   `json:"updatedAt"`
}
