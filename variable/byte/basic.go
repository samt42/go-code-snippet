package main

import (
	"fmt"
)

func main() {
	var a byte
	var b []byte
	var c []byte
	a = 'a'
	fmt.Println(a)
	b = []byte("abc")
	fmt.Println(b)
	c = []byte("你好")
	fmt.Println(c)
}
