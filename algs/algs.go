package algs

import (
	"fmt"
	"mixgo/math_tools"
)

//algorithms

func E(m int, n int) int {
	for {
		if n == 0 {
			break
		}
		fmt.Println(m, n)
		return E(n, m%n)
	}
	return m
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

func Eo(m int, n int) (A int, B int, D int) {
	if m&n > 0 {
		fmt.Printf("A.1: correct. %v & %v > 0 \n", m, n)
		var (
			a  int
			a_ int
			b  int
			b_ int
			c  int
			d  int
			q  int
			r  int
		)
		b = 1
		a_ = b
		b_ = 0
		a = b_
		c = m
		d = n
		fmt.Printf("Eo.1:  %v  %v  %v  %v  %v  %v\n", a_, a, b_, b, c, d)
		//fmt.Printf("A.2:  c %v = m %v > 0,  d %v = n %v > 0,\n  a %v = b_ %v = 0,  a_ %v = b %v = 1\n", c, m, d, n, a, b_, a_, b)
		for {
			q, r = math_tools.FloorDiv(c, d)
			fmt.Printf("Eo.2:  %v , %v = %v / %v\n", q, r, c, d)

			fmt.Printf("Eo.3:  r = %v\n", r)
			if r == 0 {
				break
			}
			c = d
			d = r
			t := a_
			a_ = a
			a = t - q*a
			t = b_
			b_ = b
			b = t - q*b
			fmt.Printf("Eo.4:  a_  a   b_  b   c   d   q   r\n")
			fmt.Printf("      %v  %v  %v  %v  %v  %v  %v  %v\n", a_, a, b_, b, c, d, q, r)
			A = a
			B = b
			D = d
		}
		if A*m+B*n == D {
			fmt.Printf("Eo.End: correct. %v * %v + %v * %v = %v + %v = %v == gsd: %v\n", A, m, B, n, A*m, B*n, A*m+B*n, D)
		} else {
			fmt.Printf("Eo.End: error. %v * %v + %v * %v = %v + %v = %v || gsd: %v\n", A, m, B, n, A*m, B*n, A*m+B*n, D)
		}

	} else {
		fmt.Printf("A.1: error. %v & %v <= 0 \n", m, n)
	}
	return A, B, D
}
