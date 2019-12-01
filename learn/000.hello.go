package main

import (
	"fmt"
	"time"
)

func main() {
	i  := 1
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	var ut int64 = time.Now().Unix();
	fmt.Println(time.Unix(ut,0).Format("2006-01-02 15:04:05"))
	fmt.Println("type = %T\n", i)
	fmt.Printf("hello, world\n") //kang
}
