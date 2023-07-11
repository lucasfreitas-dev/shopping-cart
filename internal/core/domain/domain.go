package domain

import "github.com/shopspring/decimal"

type Item struct {
	ID    string
	Name  string
	Price decimal.Decimal
}

type Cart struct {
	UserID string
	Items  []Item
}

type CartTotalPrice struct {
	Cart
	TotalPrice decimal.Decimal
}
