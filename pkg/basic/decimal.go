package basic

import "github.com/shopspring/decimal"

func StrToDecimal(str string) decimal.Decimal {
	v, _ := decimal.NewFromString(str)
	return v
}
