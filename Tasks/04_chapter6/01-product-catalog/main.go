package main

import "fmt"

func main() {
	catalog := NewCatalog()

	// Add products
	catalog.AddProduct(Product{ID: 1001, Name: "Laptop", Price: 999.99, Stock: 10})
	catalog.AddProduct(Product{ID: 1002, Name: "Mouse", Price: 29.99, Stock: 50})
	catalog.AddProduct(Product{ID: 1003, Name: "Keyboard", Price: 79.99, Stock: 25})
	catalog.AddProduct(Product{ID: 1003, Name: "Keyboard", Price: 88.99, Stock: 15})

	// Get specific product
	if product, found := catalog.GetProduct(1001); found {
		fmt.Printf("Found: %s - $%.2f (Stock: %d)\n",
			product.Name, product.Price, product.Stock)
	}

	// List all products
	fmt.Println("\nAll Products:")
	for _, product := range catalog.ListProducts() {
		fmt.Printf("  [%d] %s - $%.2f (Stock: %d)\n",
			product.ID, product.Name, product.Price, product.Stock)
	}
}
