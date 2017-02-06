package main

import (
	"fmt"
)

func say(s string) {
	fmt.Println(s)
}

func main() {
	go say("hello")
	go say("world")
	go func() {
		fmt.Println("!")
	}()
}
