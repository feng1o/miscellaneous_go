package main

//https://leileiluoluo.com/posts/golang-context.html

import (
"context"
"fmt"
)

type ctxKey string

func main() {
	ctx := context.WithValue(context.Background(), ctxKey("a"), "a")

	get := func(ctx context.Context, k ctxKey) {
		if v, ok := ctx.Value(k).(string); ok {
			fmt.Printf("-----key = %v, value = %2s", k, v)
		}
	}
	get(ctx, ctxKey("a"))
	get(ctx, ctxKey("b"))
}
