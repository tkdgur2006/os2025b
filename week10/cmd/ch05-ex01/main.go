package main

import (
	"fmt"
)

func main() {
	numbers := [3]int{-9, 11, 7}
	for i, numbers := range numbers {
		fmt.Println(i, numbers)
	}
}
