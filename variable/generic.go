package main

import (
	"fmt"
	//"reflect"
)

func test(a interface{}) {
	switch a.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	}
	//v := reflect.ValueOf(a)
	//fmt.Println(v.Type())
}

func main() {
	test("1")
	test(1)
}
