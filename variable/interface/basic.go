package main

import (
	"fmt"
)

func main() {
	var v1 interface{} = 1
	fmt.Println("type:", v1.(type))
}
