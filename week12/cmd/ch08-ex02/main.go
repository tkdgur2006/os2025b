package main

import (
	"fmt"

	"github.com/headfirstgo/magazine"
)

func main() {
	var s1 magazine.subcriber
	var e1 magazine.Employee
	e1.Name = "박상혁"
	e1.Salary = 100000
	s1.Name = "이상혁"
	fmt.Println(e1.Name, e1.Salary)
	fmt.Println(s1.Name)
}
