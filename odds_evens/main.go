package main

import (
	"fmt"
)

func main() {
	var soi = []int{
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
	}
	for _, num := range soi {
		newNum := num % 2
		if newNum == 0 {
			fmt.Printf("%v is an even number", num)
		} else {
			fmt.Printf("%v is an odd number", num)
		}
	}
}
