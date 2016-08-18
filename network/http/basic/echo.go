package main

import (
	"io"
	"log"
	"net/http"
)

func echo(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world")
}
func main() {
	http.HandleFunc("/", echo)
	err := http.ListenAndServe(":9999", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
