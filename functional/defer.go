package main

import (
	"context"
	"fmt"
)

func main() {
	var ctx context.Context
	defer func() {
		if e := recover(); e != nil {
			fmt.Println(ctx)
            fmt.Println(ctx.Value("a"))
		}
	}()
	A(ctx)(ctx)
}
type X func(ctx context.Context) (context.Context)

func A(ctx context.Context) X {
    ctx = context.Background()
	ctx = context.WithValue(ctx, "a", "A")
	panic("xxx")
	return func(c context.Context) context.Context {
        fmt.Println("yyy")
        fmt.Println(c)
        return c
    }
}


