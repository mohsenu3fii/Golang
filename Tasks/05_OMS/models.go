package main

type OrderState int

// State machine models
const (
	Created OrderState = iota
	Paid
	VendorAccepted
	Shipped
	Delivered
	Cancelled
)

var valid_transition = map[OrderState][]OrderState{
	Created:        {Paid, VendorAccepted, Cancelled},
	Paid:           {VendorAccepted, Cancelled},
	VendorAccepted: {Shipped},
	Shipped:        {Delivered},
	Cancelled:      {},
}

// Product and OrderItem models
type Product struct {
	ID    int
	Name  string
	Price int
}

type OrderItem struct {
	Product  Product
	Quantity int
}

// Voucher System models
type Voucher interface {
	Apply(total int) int
	Code() string
}

type PercentageVoucher struct {
	code    string
	percent int
}

type FixedAmountVoucher struct {
	code   string
	amount int
}

// Order Aggregate models
type Order struct {
	ID      int
	Items   []OrderItem
	Voucher Voucher
	State   OrderState
}
