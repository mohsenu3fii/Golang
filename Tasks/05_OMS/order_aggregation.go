package main

import "fmt"

func NewOrder(id int) *Order {
	return &Order{
		ID: id,
	}
}

func (o *Order) AddItem(product Product, qty int) error {
	newItem, err := NewOrderItem(product, qty)
	if err != nil {
		return err
	}

	// - Use append for items.
	o.Items = append(o.Items, newItem)
	fmt.Println("Item added successfully.")
	return nil
}

func (o *Order) ApplyVoucher(v Voucher) error {
	if o == nil {
		return fmt.Errorf("order is nil")
	}
	if o.State != Created {
		return fmt.Errorf("voucher cannot apply in this state")
	}

	// - Prevent applying voucher twice.
	if o.Voucher != nil {
		return fmt.Errorf("voucher already applied")
	}
	o.Voucher = v
	return nil
}

func (o *Order) TotalAmount() (int, error) {
	if len(o.Items) == 0 {
		return 0, fmt.Errorf("Order has no item")
	}

	total := 0
	for _, i := range o.Items {

		total += i.TotalPrice()
	}
	if o.Voucher != nil {
		total = o.Voucher.Apply(total)
	}
	// - Prevent zero/negative final totals.
	if total <= 0 {
		return 0, fmt.Errorf("Invalid total amount")
	}

	return total, nil

}

func (o *Order) Pay() error {

	if o == nil {
		return fmt.Errorf("Order is nil")
	}
	// - Enforce business rules.
	if len(o.Items) <= 0 {
		return fmt.Errorf("order has no item")
	}
	if o.State != Created {
		return fmt.Errorf("order cannot be paid in this status %v", o.State)
	}

	totalAmount, err := o.TotalAmount()
	if err != nil {
		return err
	}

	if err := o.ChangeState(Paid); err != nil {
		return err
	}

	defer func(id int, totalAmount int) {
		fmt.Printf("orderID: %d with price: %d paid successfully\n", id, totalAmount)

	}(o.ID, totalAmount)

	return nil

}

func (o *Order) Cancel() error {
	if o == nil {
		return fmt.Errorf("Order is nil")
	}
	if err := o.ChangeState(Cancelled); err != nil {
		return err
	}
	fmt.Println("Order cancelled successfully")
	return nil
}

func (o *Order) SnapshotItems() []OrderItem {
	snapshot := make([]OrderItem, len(o.Items))
	copy(snapshot, o.Items)
	return snapshot
}
