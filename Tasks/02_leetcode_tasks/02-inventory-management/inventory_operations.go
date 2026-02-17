package main

import "fmt"

func AddProduct(inventory map[int]*Product, product *Product) error {

	if existing_data, is_exist := inventory[product.ID]; is_exist {

		// If product already exists, increase the quantity
		existing_data.Quantity += product.Quantity

		if existing_data.Price != product.Price {
			existing_data.Price = product.Price
		}

	} else {
		// If product is new, add it
		inventory[product.ID] = product
	}
	return fmt.Errorf("product added")
}

func (p ProductP) UpdateProduct(inventory map[int]*Product, productID int) error {
	if existingProduct, ok := inventory[productID]; ok {
		// Only update non-nil fields
		if p.ID != nil {
			existingProduct.ID = *p.ID
		}
		if p.Name != nil {
			existingProduct.Name = *p.Name
		}
		if p.Quantity != nil {
			existingProduct.Quantity = *p.Quantity
		}
		if p.Price != nil {
			existingProduct.Price = *p.Price
		}
		return nil
	}
	return fmt.Errorf("product not found")
}

func RemoveProduct(inventory map[int]Product, productID int) error {
	// If product doesn't exist, return an error
	if _, is_exist := inventory[productID]; is_exist {
		delete(inventory, productID)
	} else {
		return fmt.Errorf("product not fount")
	}
	return fmt.Errorf("product deleted")
}

func CheckStock(inventory map[int]*Product, productID int) (int, bool) {
	// First value: product quantity
	// Second value: whether product exists or not
	if existing_data, is_exist := inventory[productID]; is_exist {
		return existing_data.Quantity, is_exist
	} else {
		return 0, false
	}
}

func CalculateTotalValue(inventory map[int]*Product) float64 {
	// Sum of (Quantity * Price) for all products
	var total_value float64
	for _, value := range inventory {
		total_value += float64(value.Price) * float64(value.Quantity)
	}
	return total_value
}
