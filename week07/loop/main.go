package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

/*
func main() {
    var now time.Time = time.Now()
    var day int = now.Day()
    fmt.Println(day)


    univ := "Go$ Inha$"
    chager := strings.NewReplacer("$", "!")
    fixed := chager.Replace(univ)
    fmt.Println(fixed)
}
*/

func main() {
	r := bufio.NewReader(os.Stdin)
	i, err := r.ReadString('\n')
	//fmt.Println(err)
	log.Fatal(err)
	fmt.Println(i)
}
