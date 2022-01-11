package algs

import (
	"fmt"
)

//types func simple

type Inty func(n int) int
type Floaty func(n float64) float64

//types func compose

type Induct func(n int, step int, lambda Inty) (array []int, mean int)

//math

//math expressions

func MassAddictINT(array []int) (res int) {
	res = 0
	for i, val := range array {
		res += val
		println(i, res, "\n")
	}
	return
}

func MassAddictFLT(array []float64) (res float64) {
	res = 0.0
	for i, val := range array {
		res += val
		println(i, res)
	}
	return
}

func FloorDiv(a int, b int) (div int, mod int) {
	div = a / b
	mod = a % b
	fmt.Printf("%v / %v: div: %v, mod: %v\n", a, b, div, mod)
	return div, mod
}

//math lambdas

//math seq

func SeqINT(start int, step int) func() int {
	start -= step
	return func() int {
		start += step
		fmt.Println(start)
		return start
	}
}

//math induction

func Induction(n int, step int, lambda Inty) (array []int, mean int) {
	mean = n * n
	array = make([]int, n)

	last := lambda(n)
	i := 0
	next := SeqINT(last, step)
	for {
		array[n-1-i] = next()
		i += 1

		if i == n {
			break
		}
	}
	fmt.Printf("Induction : Induction(n=%v), %v = %v^2(%v)\n", n, array, n, mean)
	return
}

//math induction proofs

func IProof(lambdaI Induct, n int, step int, lambdaA Inty) (A int, B int) {
	arr, m := lambdaI(n, step, lambdaA)
	A = MassAddictINT(arr)
	B = m
	fmt.Printf("P(n=%v), {%v} == {%v}\n", n, A, B)
	return A, B
}

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

func I(step int, lambdaA Inty, lambdaB Inty) bool {
	var n int
	n = 1
	A, B := IProof(Induction, n, step, lambdaA)

	if A == B {
		fmt.Println("P(n=1) : true\n")

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
