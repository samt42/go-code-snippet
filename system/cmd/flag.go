package main

import(
     "flag"
     "fmt"
)

func main(){
    var name string
    flag.StringVar(&name, "b", "123", "name")
 
    flag.Parse()
 
    fmt.Println("name:", name)
}

