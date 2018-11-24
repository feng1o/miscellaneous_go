package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	go fmt.Println("Go 1!")
	time.Sleep(100 * time.Millisecond)
	runtime.Gosched()
	main2()
	main3()
	main4()
}

func main2() {
	var wg sync.WaitGroup
	wg.Add(3)
	go func() {
		fmt.Println("Go 2!")
		wg.Done()
	}()
	go func() {
		fmt.Println("Go 3!")
		wg.Done()
	}()
	go func() {
		fmt.Println("Go 4!")
		wg.Done()
	}()
	wg.Wait()
}

func main3() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		fmt.Println("1")
		wg.Done()
	}()
	wg.Wait()
	wg.Add(1)
	go func() {
		fmt.Println("2")
		wg.Done()
	}()
	wg.Wait()
	wg.Add(1)
	go func() {
		fmt.Println("3")
		wg.Done()
	}()
	wg.Wait()
}

func main4() {
	ch1 := make(chan int, 1)
	ch2 := make(chan int, 1)
	//ch3 := make(chan int, 3)
	go func() {
		fmt.Println("1")
		ch1 <- 2
	}()
	go func() {
		x := <-ch1
		fmt.Println(x)
		ch2 <- 1
	}()
	go func() {
		<-ch2
		fmt.Println("3")
		//ch3 <- 1
	}()
	runtime.Gosched() //这样保证了起那么能运行，会被阻塞在ch1,2
}
