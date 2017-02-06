package main

import (
	"net"
)

func main() {
	conn, err := net.Dial("tcp", ":8000")
	if err != nil {
		return
	}
	defer conn.Close()
	conn.Write([]byte("Hello World"))
}
