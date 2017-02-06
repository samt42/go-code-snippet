package main

import (
	"fmt"
	"strconv"
	"time"
)

func enqueue(channel chan int) {
	i := 0
	for {
		channel <- i
		i++
		time.Sleep(1 * time.Second)
	}
}

func dequeue(channel chan int) {
	for {
		i := <-channel
		fmt.Println(strconv.Itoa(i) + ":" + strconv.Itoa(len(channel)))
		time.Sleep(2 * time.Second)
	}
}

func main() {
	channel := make(chan int, 5)
	go enqueue(channel)
	dequeue(channel)
}
