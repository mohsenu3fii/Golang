package main

import "fmt"

func main() {
	// Create order
	order := NewOrder(1001, "customer-123")

	// Add items
	order.AddItem(OrderItem{
		ProductID: 101,
		Name:      "Laptop",
		Price:     999.99,
		Quantity:  1,
	})
	order.AddItem(OrderItem{
		ProductID: 102,
		Name:      "Mouse",
		Price:     29.99,
		Quantity:  2,
	})
	order.AddItem(OrderItem{
		ProductID: 102,
		Name:      "Mouse",
		Price:     79.99,
		Quantity:  3,
	})

	// Print summary
	order.PrintSummary()

	// Update status
	fmt.Println("\n--- Updating Status ---")
	order.UpdateStatus(string(StatusProcessing))
	fmt.Printf("New status: %s\n", order.Status)

}
