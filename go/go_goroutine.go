package main
/*
	一个简单的goroutine的例子比较直观的说明其作用

*/
import (
    "fmt"
    "runtime"
)

func say(s string) {
    for i := 0; i < 5; i++ {
        runtime.Gosched() //用于让出CPU时间片。world运行了一会碰到runtime.Gosched()就把运行的权利给hello了，world歇着了，hello继续跑。
        fmt.Println(s)
    }
}

func main() {
    go say("world")
    say("hello")
}