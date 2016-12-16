package main

import (
	"fmt"
)

func add (x int, y int) int {
	return x + y
}


func main(){
	i, j := 4, "nice"
	l := 558

	fmt.Println(add(i, l))
	fmt.Printf("%s \n", j)
}
