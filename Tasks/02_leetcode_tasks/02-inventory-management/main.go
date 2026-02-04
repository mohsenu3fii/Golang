package main

import "fmt"

type Product struct {
	ID       int
	Name     string
	Quantity int
	Price    float64
}

func main() {
	inventory := make(map[int]*Product)

	product1 := &Product{ID: 1, Name: "Laptop", Quantity: 10, Price: 1500.0}
	result := AddProduct(inventory, product1)
	fmt.Println(result)

	product2 := &Product{ID: 2, Name: "Mouse", Quantity: 50, Price: 25.0}
	result = AddProduct(inventory, product2)
	fmt.Println(result)

	product3 := &Product{ID: 2, Name: "Mouse", Quantity: 30, Price: 23.0}
	result = AddProduct(inventory, product3)
	fmt.Println(result)

	quantity, exists := CheckStock(inventory, 3)
	fmt.Printf("quantity is : %d, and exists status is: %v\n", quantity, exists)

	totalValue := CalculateTotalValue(inventory)
	fmt.Println("total value is:", totalValue)
}
