package main

import (
	"fmt"
	"time"
)
/*
0.faq
	csp的goroutine和线程区别？ 170

1.关闭chan发送导致宕机；
  关闭chann接受操作有无cach会立马发挥默认值
  关闭一个已关闭的chan导致宕机  180
2.chan不是必须关闭的，和fd必须关闭不同；关闭chan是为了告诉接收方，已发送完毕；

3.单向通道: 只能接收的，不能主动关闭，倒是compile err;
           双向chan可以转单向

4.缓冲通道：

 */

func counter(out chan<- int)  {
	for x := 0;  x < 5; x++ {
		out <- x
	}
	close(out)
}

func square(out chan<-int ,in <-chan int)  {
	for x := range in {
		out <- x*x
	}
	close(out)
}

func printer(in <-chan int) {
	for v := range in {
		fmt.Println(v)
	}
}
func main()  {
	naturals := make(chan int)
	squares := make(chan int)

	go counter(naturals)
	go square(squares, naturals)
	printer(squares)
}
func main1()  {
	naturals := make(chan int)
	squares := make(chan int)

	//counter
	go func() {
		for x :=0; ; x++ {
			naturals <- x
			if x > 5 {
				close(naturals)
				break
			}
		}
	}()

	//square
	go func() {
		for {
			x, ok := <- naturals
			if !ok {
				break
			}
			squares <- x*x
		}
	}()

	// print
	for {
		time.Sleep(1*time.Second)
		fmt.Println(<-squares)
	}
}