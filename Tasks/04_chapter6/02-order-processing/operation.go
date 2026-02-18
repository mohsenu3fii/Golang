package main

import (
	"fmt"
	"time"
)

func NewOrder(id int, customerID string) *Order {
	// Create and return new order
	// Pre-allocate Items slice with reasonable capacity
	return &Order{
		ID:         id,
		CustomerID: customerID,
		Items:      make([]OrderItem, 0, 1),
		Status:     string(StatusPending),
		CreatedAt:  time.Now(),
	}
}

func (o *Order) AddItem(item OrderItem) {

	defer func() {
		o.CalculateTotal()
	}()

	for i := range o.Items {
		if o.Items[i].ProductID == item.ProductID {
			o.Items[i].Quantity += item.Quantity
			return
		}
	}

	o.Items = append(o.Items, item)
}

func (o *Order) CalculateTotal() {
	// Calculate and return total (read-only, use value receiver)
	// *** calculate total price and total items in order is better

	for _, i := range o.Items {
		o.TotalItem += i.Quantity
		o.TotalPrice += i.Price * float64(i.Quantity)
	}
}

func (o *Order) UpdateStatus(status string) {
	if o.Status != string(StatusCancelled) && o.Status != status {
		o.Status = status
	}
}

func (o Order) PrintSummary() {
	// Print order summary
	fmt.Printf("Order ID is: %d\n", o.ID)
	fmt.Printf("CustomerID is: %s\n", o.CustomerID)
	fmt.Printf("Order TotalPrice is: %.2f\n", o.TotalPrice)
	fmt.Printf("Order Status is: %s\n", o.Status)
	fmt.Printf("Order CreatedAt is: %s\n", o.CreatedAt)
	fmt.Println("Order Items are:")
	for index, data := range o.Items {
		fmt.Printf("%d- product id: %d, product name: %s, price is: %.2f, quantity is: %d\n",
			index+1, data.ProductID, data.Name, data.Price, data.Quantity)
	}
}
