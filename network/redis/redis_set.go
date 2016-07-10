package main
import (
 //   "fmt"
    "github.com/alphazero/Go-Redis"
    "log"
   // "strconv"
)
//go get github.com/alphazero/Go-Redis
func main() {
    // 连接Redis服务器 127.0.0.1:6379
    spec := redis.DefaultSpec().Host("172.30.13.124").Port(6379)
    client, e := redis.NewSynchClientWithSpec(spec)
    // 是否连接出错
    if e != nil {
        log.Println("error on connect redis server")
        return
    }
    name := "http://lianjia.com/a?a=123&b=456!1234"
    ts := "123456"
    err := client.Set(name, []byte(ts))
    if err != nil {
    	log.Panic(err)
    }
}
