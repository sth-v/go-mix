package algs

import (
	"fmt"
)

func FloorDiv(a int, b int) (div int, mod int) {
	div = a / b
	mod = a % b
	fmt.Printf("%v / %v: div: %v, mod: %v\n", a, b, div, mod)
	return div, mod
}

func Euclid(a int, b int) int {
	for {
		if b == 0 {
			break
		}
		fmt.Println(a, b)
		return Euclid(b, a%b)
	}
	return a
}
