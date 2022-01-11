package main

import (
	"fmt"
	"mixgo/algs"
)

var a = 450
var b = 81

func main() {
	div, mod := algs.FloorDiv(a, b)
	fmt.Printf("FloorDiv result : %v, %v\n", div, mod)
	r := algs.Euclid(a, b)
	fmt.Printf("Euclid result : %v\n", r)
}
