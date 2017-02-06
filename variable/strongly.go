package main

import (
	"fmt"
)

func main() {
	var a int = 10
	var b int64 = 10
	c := 10 + b
	//if a == b { //错误
	//if a == int(b) { //正确
	//if a == c { //错误
	if a == int(c) {
		fmt.Println("Equal")
	} else {
		fmt.Println("Not Equal")
	}
}
