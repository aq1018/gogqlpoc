package model

import (
	"fmt"
	"io"
	"strconv"

	"github.com/99designs/gqlgen/graphql"
	"github.com/shopspring/decimal"
)

func MarshalDecimal(v decimal.Decimal) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		io.WriteString(w, strconv.Quote(v.String()))
	})
}

func UnmarshalDecimal(v interface{}) (decimal.Decimal, error) {
	if str, ok := v.(string); ok {
		return decimal.NewFromString(str), nil
	}

	return decimal.Zero, fmt.Errorf("%T must be a string representing a decimal", v)
}
