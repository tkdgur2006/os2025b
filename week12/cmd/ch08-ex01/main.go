package main

import "fmt"

type subcriber struct {
	name  string
	price int
}

func applyPrice(s *subcriber) {
	s.price = 10000
	s.name = "park sanghyuk"
}

func main() {
	var s1 subcriber
	//s1.name = "kim sanghyuk"
	applyPrice(&s1)
	fmt.Println(s1.name, s1.price)
}
