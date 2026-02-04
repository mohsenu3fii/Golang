package main

import (
	"fmt"
)

func SumNumbers(numbers []int) int {
	var sum int
	if len(numbers) > 0 {
		for _, i := range numbers {
			sum += i
		}
		return sum
	} else {
		return 0
	}

}

func main() {
	test1 := []int{1, 2, 3, 4, 5, 6}
	fmt.Println("sum of test1 is: ", SumNumbers(test1))
}
