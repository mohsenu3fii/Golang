package main

import "fmt"

//Why do you use *Product for the value of your inventory?
func AddProduct(inventory map[int]*Product, product *Product) error {

	//You can choose a better name like data, exist
	if existing_data, is_exist := inventory[product.ID]; is_exist {

		// If product already exists, increase the quantity
		existing_data.Quantity += product.Quantity

		if existing_data.Price != product.Price {
			existing_data.Price = product.Price
		}
// You can return nil this line
	} else { // Remove this else
		// If product is new, add it
		inventory[product.ID] = product
	}

	//Why do you use error when a product added?
	return fmt.Errorf("product added")
}

//Why do you use pointer for product?
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

	//Instead of use "product not found" as a string, use const
	return fmt.Errorf("product not found")
}

func RemoveProduct(inventory map[int]Product, productID int) error {
	// If product doesn't exist, return an error
	if _, is_exist := inventory[productID]; is_exist {
		delete(inventory, productID)
		// return nil
	} else {//Remove this else
		return fmt.Errorf("product not fount")
	}

	//you must return nil and write a log as an info not return error
	return fmt.Errorf("product deleted")
}

//You can use a single response and it be int instead of (int, bool)
func CheckStock(inventory map[int]*Product, productID int) (int, bool) {
	// First value: product quantity
	// Second value: whether product exists or not
	if existing_data, is_exist := inventory[productID]; is_exist {

		//Just return the quantity
		return existing_data.Quantity, is_exist
	} else {// remove this else
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
