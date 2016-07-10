package main

import (
	"fmt"
)

func main() {
	var a string
	a = "你好"
	fmt.Printf("%s", a)
	fmt.Printf("\nByte:\n")
	for i := 0; i < len(a); i++ {
		fmt.Printf("%d, %x\n", a[i], a[i])
	}
	fmt.Printf("\nRune:\n")
	for _, b := range a {
		fmt.Printf("%x\n", b)
	}
}
