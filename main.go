package main

import (
	"fmt"
	"mixgo/algs"
	"mixgo/math_tools"
)

var a = 450
var b = 81
var step = -2
var arrFlt = []float64{0.20, 0.33, 8.65}

func lambda001(n int) int {
	return 2*n - 1
}
func lambda002(n int) int {
	return 2*n + 1
}

func main() {

	math_tools.FloorDiv(a, b)
	math_tools.MassAddictFLT(arrFlt)

	messageE := algs.E(a, b)
	messageI := algs.I(step, lambda001, lambda002)

	fmt.Printf("E: %v\n", messageE)
	fmt.Printf("I: %v\n", messageI)
}
