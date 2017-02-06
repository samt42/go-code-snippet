package main

import (
	"encoding/hex"
	"fmt"
)

func main() {
	a := []byte("中国")
	b := make([]byte, hex.EncodedLen(len(a)))
	len := hex.Encode(b, a)
	fmt.Printf("len: %d\n", len)
	for _, v := range b {
		fmt.Printf("%x,", v)
	}
}
