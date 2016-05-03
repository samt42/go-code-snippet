package main

import (
	"fmt"
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
}
