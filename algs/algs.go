package algs

import (
	"fmt"
	"mixgo/math_tools"
)

//algorithms

func E(a int, b int) int {
	for {
		if b == 0 {
			break
		}
		fmt.Println(a, b)
		return E(b, a%b)
	}
	return a
}

func I(step int, lambdaA math_tools.Inti, lambdaB math_tools.Inti) bool {
	var n int
	n = 1
	A, B := math_tools.IProof(math_tools.Induction, n, step, lambdaA)

	if A == B {

		fmt.Println("P(n=1) : true")

		A := A + lambdaB(n)
		B := B + lambdaB(n)
		n := n + 1
		C := n * n

		fmt.Printf(" {%v}; n^2 + 2n + 1 = {%v}; (n+1)^2 = {%v}\n ", A, B, C)

		if A == B&C {
			fmt.Println("P(n+1) : true")
			return true
		} else {
			fmt.Println("P(n+1) : false")
			return false
		}
	} else {
		fmt.Println("P(n=1) : false")
		return false
	}
}
