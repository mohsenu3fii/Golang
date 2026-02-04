package main

func CalculateSubtotal(order Order) float64 {
	// Sum of (Quantity * Price) for all items
	var subtotal float64
	for _, items := range order.Items {
		subtotal += float64(items.Quantity) * float64(items.Price)
	}
	return subtotal
}

func (order *Order) ApplyDiscountRules(rules []DiscountRule) Order {
	// Check if order total meets minimum for each rule
	// Apply the maximum possible discount
	// Update order.Discount and order.Total

	best_rule_index := 0
	discount_matched := false
	order.SubTotal = CalculateSubtotal(*order)
	for i, value := range rules {
		if order.SubTotal > value.MinAmount {
			discount_matched = true
			best_rule_index = i
		}
	}
	if discount_matched {
		order.Discount = order.SubTotal * (rules[best_rule_index].DiscountPercent / 100)
	}
	return *order

}

func ProcessOrders(orders []Order, rules []DiscountRule) []Order {
	// For each order:
	for i := range orders {
		orderPointer := &orders[i]

		// 1. Change status to "processing"
		orderPointer.Status = "processing"

		// 2. Apply discount
		orderPointer.ApplyDiscountRules(rules)

		// 3. Calculate final total
		orderPointer.FinalPrice = orderPointer.SubTotal - orderPointer.Discount

		// 4. Change status to "completed"
		// If order total is less than 0, set status to "cancelled"
		if orderPointer.FinalPrice <= 0 {
			orderPointer.Status = "cancelled"
		} else {
			orderPointer.Status = "completed"
		}
	}
	return orders
}

func FilterOrdersByStatus(orders []Order, status string) []Order {
	// Return only orders with the specified status
	query_matched := make([]Order, 0)
	for i := range orders {
		if orders[i].Status == status {
			query_matched = append(query_matched, orders[i])
		}
	}
	return query_matched
}

func CalculateOrderStatistics(orders []Order) map[string]interface{} {

	var total_orders, completed_orders int
	var total_revenue, total_discount, average_order_value float64

	for _, order := range orders {

		//   - "total_revenue": total revenue
		total_revenue += order.FinalPrice

		//   - "total_discount": total discounts applied
		total_discount += order.Discount
	}

	//   - "average_order_value": average order value
	if total_orders > 0 {
		average_order_value = total_revenue / float64(total_orders)
	}

	//   - "completed_orders": number of completed orders
	completed_orders = len(FilterOrdersByStatus(orders, "completed"))

	//   - "total_orders": total number of orders
	return map[string]any{
		"total_orders":        len(orders),
		"total_revenue":       total_revenue,
		"total_discount":      total_discount,
		"completed_orders":    completed_orders,
		"average_order_value": average_order_value,
	}
}
