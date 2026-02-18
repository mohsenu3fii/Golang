package main

type Product struct {
	ID    int
	Name  string
	Price float64
	Stock int
}

type Catalog struct {
	Products map[int]Product
}
