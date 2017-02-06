package main

import (
	"fmt"
)

type MyInt int64

func (i MyInt) Abs() int64 {
	if i < 0 {
		return int64(-i)
	}
	return int64(i)
}

func main() {
	var num MyInt = -1
	fmt.Println(num.Abs())
}
