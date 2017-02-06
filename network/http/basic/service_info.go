package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func info(w http.ResponseWriter, r *http.Request) {
	var output string
	output += fmt.Sprintf("Method:%s\n", r.Method)
	output += fmt.Sprintf("Path:%s\n", r.URL.Path)
	output += fmt.Sprintf("Scheme:%s\n", r.URL.Scheme)
	output += fmt.Sprintf("Fragment:%s\n", r.URL.Fragment)
	output += fmt.Sprintf("RawQuery:%s\n", r.URL.RawQuery)
	u, _ := url.Parse(r.URL.RawQuery)
	output += fmt.Sprintf("Fragment:%s\n", u.Fragment)

	r.ParseForm()
	output += fmt.Sprintf("POST:\n")
	for k, v := range r.Form {
		output += fmt.Sprintf("key:%s value:%s\n", k, strings.Join(v, ""))
	}
	fmt.Println(output)
	fmt.Fprintf(w, output)
}

func main() {
	http.HandleFunc("/", info)               //设置访问的路由
	err := http.ListenAndServe(":9999", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
