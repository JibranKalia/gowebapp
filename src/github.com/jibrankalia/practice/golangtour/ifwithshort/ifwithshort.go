package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		fmt.Printf("%g < %g. V is smaller than Lim. \n", v, lim)
	} else {
		fmt.Printf("%g >= %g V is greater than Lim\n ", v, lim)
	}
}

func main() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}
