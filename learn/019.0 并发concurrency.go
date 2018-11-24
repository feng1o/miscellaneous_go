package main

import (
	"fmt"
	"os"
	"runtime"
	"time"
)

import (
	"bytes"
	lg "log"
)

func main() {
	var loger = lg.New(&bytes.Buffer{}, "logger: ", lg.Lshortfile)
	loger.Println(runtime.NumCPU())
	/*go Go()*/
	loger.SetOutput(os.Stdout)
	time.Sleep(1 * time.Second)
	c := make(chan bool) //双向通道
	go func() {
		loger.Println("将要放入chan: go ---test ")
		c <- true
		close(c) //必须明确要关闭，否则会都在等
	}()
	//<- c
	for v := range c {
		loger.Println(v)
	}
}

func Go() {
	fmt.Println("go go---!\n")
}
