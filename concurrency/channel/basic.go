package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		for {
			fmt.Println("学PHP")
			time.Sleep(1 * time.Second)
		}
	}()

	go func() {
		for {
			fmt.Println("学Go")
			time.Sleep(1 * time.Second)
		}
	}()

	select {
	case <-time.After(2 * time.Second):
		fmt.Println("休息一下")
	}
}
