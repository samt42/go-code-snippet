package main

import (
	"context"
	"time"
	"fmt"
	//"runtime"
)

func main() {
	c := context.Background()
	c, cancelFunc := context.WithTimeout(c, 4 * time.Second)
	defer cancelFunc()
	errChan := make(chan error, 1)
	go func(ctx context.Context, errChan chan error) {
		err :=  a()
		fmt.Println(err)
	}(c, errChan)
	select {
	case <-c.Done():
		switch c.Err() {
		case context.DeadlineExceeded:
			fmt.Println("DeadlineExceeded")
		case context.Canceled:
			fmt.Println("Canceled")
		}
	}
/*	time.Sleep(3*time.Second)
	fmt.Println(runtime.NumGoroutine())*/
}

func a() error {
	time.Sleep(3*time.Second)
	return nil
}
