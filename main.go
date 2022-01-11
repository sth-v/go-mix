package main

import (
	"fmt"
	"mixgo/algs"
)

var a = 450
var b = 81
var step = -2

func lambda001(n int) int {
	return 2*n - 1
}
func lambda002(n int) int {
	return 2*n + 1
}

func main() {

	messageE := algs.E(a, b)
	messageI := algs.I(step, lambda001, lambda002)

	fmt.Printf("E: %v\n", messageE)
	fmt.Printf("I: %v\n", messageI)
}
