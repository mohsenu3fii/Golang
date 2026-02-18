package main

func NewCatalog() *Catalog {
	// Initialize and return catalog
	return &Catalog{
		Products: make(map[int]Product),
	}
}

func (c *Catalog) AddProduct(p Product) {
	// Add product to catalog
	if index, ok := c.Products[p.ID]; ok {
		index.Stock += p.Stock
		index.Price = p.Price
		c.Products[p.ID] = index
	} else {
		c.Products[p.ID] = p
	}
}

func (c *Catalog) GetProduct(id int) (Product, bool) {
	// Retrieve product by ID
	// Return product and true if found, zero value and false if not found
	if index, ok := c.Products[id]; ok {
		return index, true
	}
	return Product{}, false
}

func (c *Catalog) ListProducts() []Product {
	// Return all products as a slice
	data := make([]Product, 0, len(c.Products))
	for _, i := range c.Products {
		data = append(data, i)
	}
	return data
}
