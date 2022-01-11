package math_tools

import (
	"fmt"
)

//types func simple

type Inti func(n int) int
type Floaty func(n float64) float64

//types func compose

type Induct func(n int, step int, lambda Inti) (array []int, mean int)

//math_tools

//math_tools expressions

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

//math_tools lambdas

//math_tools seq

func SeqINT(start int, step int) func() int {
	start -= step
	return func() int {
		start += step
		fmt.Println(start)
		return start
	}
}

//math_tools induction

func Induction(n int, step int, lambda Inti) (array []int, mean int) {
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

//math_tools induction proofs

func IProof(lambdaI Induct, n int, step int, lambdaA Inti) (A int, B int) {
	arr, m := lambdaI(n, step, lambdaA)
	A = MassAddictINT(arr)
	B = m
	fmt.Printf("P(n=%v), {%v} == {%v}\n", n, A, B)
	return A, B
}
