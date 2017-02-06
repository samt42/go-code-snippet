package main

import (
	"fmt"
	"time"
)

//初始化通道
func main() {
	c := make(chan string) //channel
	fix := func(man string) {
		fmt.Println(man + "修复Bug")
		time.Sleep(2 * time.Second)
		c <- man + "修复完成"
	}
	fmt.Println("Coding")
	fmt.Println("Coding")
	fmt.Println("收到一个Bug")
	go fix("A")
	for {
		select {
		case msg := <-c:
			fmt.Println(msg)
		default:
			fmt.Println("Coding")
			time.Sleep(1 * time.Second)
		}
	}
}
