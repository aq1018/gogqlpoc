package model

import "github.com/thoas/go-funk"

func (r *LoanRequest) Validate() []string {
	var errors []string
	uniq_addresses := funk.Uniq(r.Addresses).([]Address)
	if len(uniq_addresses) < len(r.Addresses) {
		errors = append(errors, "Addresses must be unique")
	}

	return errors
}
