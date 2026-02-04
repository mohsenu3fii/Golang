package main

type OrderItem struct {
	ProductID int
	Quantity  int
	Price     float64
}

type Order struct {
	ID         int
	Items      []OrderItem
	Discount   float64 // Discount percentage (0-100)
	SubTotal   float64
	FinalPrice float64
	Status     string // "pending", "processing", "completed", "cancelled"
}

type DiscountRule struct {
	MinAmount       float64
	DiscountPercent float64
	Description     string
}
