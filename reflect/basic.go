package main

import (
	"fmt"
	"reflect"
)

func main() {
	m := map[string]int{}
	mt := reflect.TypeOf(m)
	nm := reflect.MakeMap(mt)
}
