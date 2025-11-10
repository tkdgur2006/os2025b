package main

import "fmt"

func main() {
	subjects := []string{"Go", "Javascript", "Python", "Linux"}
	subjectsSlice := subjects[:3]
	subjectsSlice[0] = "Java"
	for _, subjects := range subjects {
		fmt.Println(subjects)
	}
	fmt.Println("===========")
	for i := 0; i < len(subjectsSlice); i++ {
		fmt.Println(subjectsSlice[i])
	}
}
