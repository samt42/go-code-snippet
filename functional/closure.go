package main

import (
	"fmt"
)

func main() {
	var factorial func(int) int
	factorial = func(n int) int {
		if n == 1 {
			return 1
		}
		return n * factorial(n-1)
	}
	f := factorial
	fmt.Println(f(3))
}
