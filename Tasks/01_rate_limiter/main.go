package main

import (
	"fmt"
	"time"
)

func main() {
	rateLimit := CreateRateLimit(5 * time.Second)
	rateLimit2 := CreateRateLimit(3 * time.Second)

	for i := range 10 {
		test1 := rateLimit("mohsen")
		test2 := rateLimit2("amir")

		switch test1 {
		case true:
			fmt.Printf("Mohsen %d: ALLOWED\n", i+1)
		case false:
			fmt.Printf("Mohsen %d: DENIED\n", i+1)
		}

		switch test2 {
		case true:
			fmt.Printf("Akbar %d: ALLOWED\n", i+1)
		case false:
			fmt.Printf("Akbar %d DENIED\n", i+1)
		}

		fmt.Println("---")
		time.Sleep(1 * time.Second)
	}
}
