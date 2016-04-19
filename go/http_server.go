package main

import (
    "fmt"
    "net/http"
    "net/url"
    "log"
)

func Test2(w http.ResponseWriter, r *http.Request){
    if r.Method == "GET" {
	fmt.Println("GET")
    	log.Println(r.URL.RawQuery)
	log.Println(r.URL.Path)
	m, _ := url.ParseQuery(r.URL.RawQuery)
	fmt.Println("ak = ", m["ak"][0], "sk = ", m["sk"][0])
      }
    if r.Method == "PUT" {
	fmt.Println("PUT")
	log.Println(r.URL.RawQuery)
        log.Println(r.URL.Path)
        m, _ := url.ParseQuery(r.URL.RawQuery)
        fmt.Println("ak = ", m["ak"][0], "sk = ", m["sk"][0])
      }
}
func main() {

    stat := http.StatusText(200)
    fmt.Println(stat) //状态码200对应的状态OK

    stringtype := http.DetectContentType([]byte("test"))
    fmt.Println(stringtype) 

    http.HandleFunc("/", Test2)
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        fmt.Println(err)
    }

}
