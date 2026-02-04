package main

import "fmt"

func main() {
	rules := []DiscountRule{
		{MinAmount: 100.0, DiscountPercent: 5.0, Description: "5% off for orders over $100"},
		{MinAmount: 500.0, DiscountPercent: 10.0, Description: "10% off for orders over $500"},
		{MinAmount: 1000.0, DiscountPercent: 15.0, Description: "15% off for orders over $1000"},
	}

	// Create orders
	orders := []Order{
		{
			ID: 1,
			Items: []OrderItem{
				{ProductID: 1, Quantity: 2, Price: 50.0},
				{ProductID: 2, Quantity: 1, Price: 30.0},
			},
			Status: "pending",
		},
		{
			ID: 2,
			Items: []OrderItem{
				{ProductID: 3, Quantity: 10, Price: 100.0},
			},
			Status: "pending",
		},
	}

	// Process orders
	processedOrders := ProcessOrders(orders, rules)

	// Filter completed orders
	completedOrders := FilterOrdersByStatus(processedOrders, "completed")
	fmt.Println(completedOrders)

	// Calculate statistics
	stats := CalculateOrderStatistics(processedOrders)
	fmt.Println(stats)
}
