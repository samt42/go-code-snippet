package main

import (
	"fmt"
	"reflect"
)

func main() {
	var resolved interface{}

	m := map[string]int{}
	mt := reflect.TypeOf(m)
	nm := reflect.MakeMap(mt)
}
