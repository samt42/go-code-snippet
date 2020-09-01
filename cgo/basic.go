package main

import (
	"fmt"
)

func main() {
	call(2, 1, Add)
	call(2, 1, Sub)
}

func Add(a, b int) {
	fmt.Printf("The sum of %d and %d is: %d\n", a, b, a+b)
}

func Sub(a, b int) {
	fmt.Printf("The sum of %d and %d is: %d\n", a, b, a-b)
}

func call(x int, y int, f func(int, int)) {
	f(x, y)
}
