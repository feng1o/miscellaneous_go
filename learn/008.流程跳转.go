package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var number int = 5
	if number += 4; 10 > number {
		number += 3
		fmt.Println(number)
	} else if 10 < number {
		number -= 2
		fmt.Println(number)
	}
	fmt.Println(number)
	fmt.Println("switch____________")
	switchTest()
	forTest()
	selectTest()
}

func switchTest() {
	ia := []interface{}{byte(6), 'a', uint(10), int32(-4)}
	switch v := ia[rand.Intn(4)%2]; interface{}(v).(type) {
	case int32:
		fmt.Printf("Case A.\n")
	case byte:
		fmt.Printf("Case B.\n")
	default:
		fmt.Println("Unknown!")
	}

	names := []string{"Golang", "Java", "Rust", "C"}
	switch name := names[0]; name {
	case "Golang":
		fmt.Println("A programming language from Google.")
	case "Rust":
		fmt.Println("A programming language from Mozilla.")
	default:
		fmt.Println("Unknown!")
	}
}

func forTest() {
	for i, v := range "Go语言" {
		fmt.Printf("%d: %c\n", i, v)
	}
}

func selectTest() {
	ch4 := make(chan int, 1)
	for i := 0; i < 4; i++ {
		select {
		case e, ok := <-ch4:
			if !ok {
				fmt.Println("End.")
				return
			}
			fmt.Println(e)

		default:
			fmt.Println("No Data!")

		}
	}

	ch3 := make(chan int, 100)
	// 省略若干条语句
	select {
	case ch4 <- 1:
		fmt.Printf("Sent %d\n", 1)
	case ch3 <- 2:
		fmt.Printf("Sent %d\n", 2)
	default:
		fmt.Println("Full channel!")
	}
}

/*
 select语句属于条件分支流程控制方法，不过它只能用于通道。它可以包含若干条case语句，并根据条件选择其中的一个执行。进一步说，select语句中的case关键字只能后跟用于通道的发送操作的表达式以及接收操作的表达式或语句。示例如下：

ch1 := make(chan int, 1)
ch2 := make(chan int, 1)
// 省略若干条语句
select {
case e1 := <-ch1:
    fmt.Printf("1th case is selected. e1=%v.\n", e1)
case e2 := <-ch2:
    fmt.Printf("2th case is selected. e2=%v.\n", e2)
default:
    fmt.Println("No data!")
}

    如果该select语句被执行时通道ch1和ch2中都没有任何数据，那么肯定只有default case会被执行。但是，只要有一个通道在当时有数据就不会轮到default case执行了。显然，对于包含通道接收操作的case来讲，其执行条件就是通道中存在数据（或者说通道未空）。如果在当时有数据的通道多于一个，那么Go语言会通过一种伪随机的算法来决定哪一个case将被执行。

   另一方面，对于包含通道发送操作的case来讲，其执行条件就是通道中至少还能缓冲一个数据（或者说通道未满）。类似的，当有多个case中的通道未满时，它们会被随机选择。请看下面的示例：

ch3 := make(chan int, 100)
// 省略若干条语句
select {
case ch3 <- 1:
    fmt.Printf("Sent %d\n", 1)
case ch3 <- 2:
    fmt.Printf("Sent %d\n", 2)
default:
    fmt.Println("Full channel!")
}

    该条select语句的两个case中包含的都是针对通道ch3的发送操作。如果我们把这条语句置于一个循环中，那么就相当于用有限范围的随机整数集合去填满一个通道。

    请注意，如果一条select语句中不存在default case， 并且在被执行时其中的所有case都不满足执行条件，那么它的执行将会被阻塞！当前流程的进行也会因此而停滞。直到其中一个case满足了执行条件，执行才会继续。我们一直在说case执行条件的满足与否取决于其操作的通道在当时的状态。这里特别强调一点，即：未被初始化的通道会使操作它的case永远满足不了执行条件。对于针对它的发送操作和接收操作来说都是如此。

    最后提一句，break语句也可以被包含在select语句中的case语句中。它的作用是立即结束当前的select语句的执行，不论其所属的case语句中是否还有未被执行的语句。*/
