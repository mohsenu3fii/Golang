package main

import (
	"fmt"
	"time"
)

func main() {
	rateLimit := CreateRateLimit(5 * time.Second)

	for i := range 5 {
		test := rateLimit("mohsen")

		switch test {
		case true:
			fmt.Printf("Mohsen %d: ALLOWED\n", i+1)
		case false:
			fmt.Printf("Mohsen %d: DENIED\n", i+1)
		}

		time.Sleep(1 * time.Second)
	}

	for i := range 10 {
		test := rateLimit("amir")

		switch test {
		case true:
			fmt.Printf("amir %d: ALLOWED\n", i+1)
		case false:
			fmt.Printf("amir %d: DENIED\n", i+1)
		}

		time.Sleep(1 * time.Second)
	}
}
