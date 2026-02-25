package main

import "fmt"

func main() {
	order := NewOrder(1)
	fmt.Printf("Created Order ID: %d, State: %v\n", order.ID, order.State)

	prod1 := Product{ID: 1, Name: "Laptop", Price: 1000}
	prod2 := Product{ID: 2, Name: "Mouse", Price: 100}
	prod3 := Product{ID: 3, Name: "Mobile", Price: 700}
	// add item
	if err := order.AddItem(prod1, 1); err != nil {
		fmt.Println("AddItem Error:", err)
	}
	if err := order.AddItem(prod2, 3); err != nil {
		fmt.Println("AddItem Error:", err)
	}

	// add item failed
	if err := order.AddItem(prod3, 0); err != nil {
		fmt.Println("AddItem Error:", err)
	}

	// snapshot
	fmt.Println("Items added. Snapshot:", order.SnapshotItems())

	// Voucher
	percentVoucher, _ := NewPercentageVoucher("Percentage", 10)
	if err := order.ApplyVoucher(percentVoucher); err != nil {
		fmt.Println("ApplyVoucher Error:", err)
	} else {
		fmt.Println("Percentage voucher applied:", percentVoucher.Code())
	}
	// failed to create
	if _, err := NewFixedAmountVoucher("Fixed", 0); err != nil {
		fmt.Println("Create voucher error:", err)
	}
	// failed to apply
	fixedVoucher, _ := NewFixedAmountVoucher("Fixed", 200)
	if err := order.ApplyVoucher(fixedVoucher); err != nil {
		fmt.Println("ApplyVoucher Error:", err)
	}

	total, err := order.TotalAmount()
	if err != nil {
		fmt.Println("TotalAmount Error:", err)
	} else {
		fmt.Printf("TotalAmount with voucher applied: %d\n", total)
	}

	// failed to change status
	err = order.ChangeState(Shipped)
	if err != nil {
		fmt.Println("ChangeState Error:", err)
	}

	if err := order.Pay(); err != nil {
		fmt.Println("Pay Error:", err)
	} else {
		fmt.Printf("Order State after Pay: %v\n", order.State)
	}

	err = order.ChangeState(VendorAccepted)
	if err != nil {
		fmt.Println("ChangeState Error:", err)
	} else {
		fmt.Printf("Order State after change: %v\n", order.State)
	}
}
