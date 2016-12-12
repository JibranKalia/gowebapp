package main

import "fmt"

func main() {

	easyArray := [2][4]int{{3, 3, 22, 3}, {23, 3, 222, 3}}

	var ar = [10]byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j'}

	fmt.Println(easyArray[0][3])
	fmt.Printf("%s", ar[2])
}
