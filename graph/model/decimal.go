package model

import (
	"fmt"
	"io"

	"github.com/99designs/gqlgen/graphql"
	"github.com/shopspring/decimal"
)

func MarshalDecimal(v decimal.Decimal) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		w.Write([]byte(v.String()))
	})
}

func UnmarshalDecimal(v interface{}) (d decimal.Decimal, err error) {
	switch v := v.(type) {
	case string:
		d, err = decimal.NewFromString(v)
	default:
		err = fmt.Errorf("%T must be a string representing a decimal", v)
	}

	return d, err
}
