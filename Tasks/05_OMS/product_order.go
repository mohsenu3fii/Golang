package main

import "fmt"

func NewOrderItem(product Product, qty int) (OrderItem, error) {
	// - Prevent zero/negative quantity.
	// - No invalid construction.
	if qty < 1 || product.Price < 1 {
		return OrderItem{}, fmt.Errorf("Add product failed. Quantity must be greater than zero")
	}
	return OrderItem{
		Quantity: qty,
		Product:  product,
	}, nil
}

func (oi OrderItem) TotalPrice() int {

	// - TotalPrice = Price × Quantity.
	return oi.Product.Price * oi.Quantity
}
