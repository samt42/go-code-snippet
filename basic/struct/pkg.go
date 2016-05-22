package main

import (
	"fmt"
	"go-code-snippet"
)

func main() {
	type struct_a struct {
		A string
		B string
	}
	var a = new(struct_a)
	var b *struct_a
	a = &struct_a{"Hello", "World"}
	b = a
	fmt.Println(a.A, a.B)
	fmt.Println(b.A, b.B)

	var c = &struct_pkg.T{1, 2}
	fmt.Println(c.X, c.Y)

	type struct_b struct {
		struct_pkg
	}
	var d = &struct_b.T{1, 2}
	fmt.Println(c.X, c.Y)
}
