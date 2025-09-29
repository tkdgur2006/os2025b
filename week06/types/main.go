package main

import (
	"fmt"
	"reflect"
)

func main() {
	/* 같은 의미
	var name string = "Kim inha"
	var id int = 1000

	fmt.Println(name, reflect.TypeOf(name))
	fmt.Println(id, reflect.TypeOf(id))
	*/
	var gpa float32 = 3.99 // 원하는 타입으로 만들수 있음 (메모리 절약)
	name := "Kim inha"
	id := 1000
	//gpa := 3.99 double 로 선언됨 (쓸대없이 큰 메모리 차지함)

	fmt.Println(name, reflect.TypeOf(name))
	fmt.Println(id, reflect.TypeOf(id))
	fmt.Println(gpa, reflect.TypeOf(gpa))
}
