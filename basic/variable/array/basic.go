package main

import (
	"fmt"
)

func main() {
	var a []int = []int{97, 98}
	fmt.Println(a)
	var b []int = make([]int, 2, 5)
	fmt.Println(b, " len:", len(b), " cap:", cap(b))
	b = []int{99, 100}
	fmt.Println(b, " len:", len(b), " cap:", cap(b))
	b = []int{99, 100, 101, 102}
	fmt.Println(b, " len:", len(b), " cap:", cap(b))
	var c = []string{"Hello", "World"}
	for i, v := range c {
		fmt.Printf("%d %s\n", i, v)
	}
}
