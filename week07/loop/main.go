package main

import (
	"bufio"
	"fmt"
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
	i, _ := r.ReadString('\n')
	fmt.Println(i)
}
