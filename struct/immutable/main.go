package main

import "math/big"

type Currency int

const (
	JPY Currency = iota
	USD
	CNY
)

// struct-immutable
type ImmutableMoney struct {
	currency Currency
	amount   *big.Int
}

func (im ImmutableMoney) Currency() Currency {
	return im.currency
}

func (im ImmutableMoney) SetCurrency(c Currency) ImmutableMoney {
	return ImmutableMoney{
		currency: c,
		amount:   im.amount,
	}
}

// struct-immutable

// struct-mutable
type MutableMoney struct {
	currency Currency
	amount   *big.Int
}

func (m MutableMoney) Currency() Currency {
	return m.currency
}

func (m *MutableMoney) SetCurrency(c Currency) {
	m.currency = c
}

// struct-mutable
