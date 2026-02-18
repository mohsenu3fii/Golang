package main

import (
	"time"
)

type OrderItem struct {
	ProductID int
	Name      string
	Price     float64
	Quantity  int
}

type Order struct {
	ID         int
	CustomerID string
	Items      []OrderItem
	TotalPrice float64
	TotalItem  int
	Status     string
	CreatedAt  time.Time
}

type Status string

const (
	StatusPending    Status = "pending"
	StatusProcessing Status = "processing"
	StatusComplete   Status = "complete"
	StatusCancelled  Status = "cancelled"
)
