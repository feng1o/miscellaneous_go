package main // 代码包声明语句

import (
	"fmt" // 导入代码包fmt。
	"os"
	"reflect"
	"time"
	"unsafe"
)

type Apple struct {
	Price int8
	Weight int16
	Name string
}

func main() {
	var num uint64 = 65535
	var apple Apple
	apple = Apple{1, 9, "apple"}
	fmt.Fprintf(os.Stdout, "-----size of %v ", apple)
	fmt.Fprintf(os.Stdout, "-----size of %#v ", apple)
	fmt.Fprintf(os.Stdout, "-----size of %#v ", unsafe.Sizeof(num))
	xx := Apple{1,2,"xxx"}
	fmt.Println(reflect.TypeOf(xx))
	// 短变量声明语句，由变量名size、特殊标记:=，以及值（需要你来填写）组成。
	size := (8)
	fmt.Printf("类型为 uint64 的整数 %d 需占用的存储空间为 %d 个字节。\n", num, size)
	//fmt.Println("___________fun__________")
	main2()
	integerTrs()
	floatFun()
	complexFun()
	runeFun()
}

//1.数组
func integerTrs() {
	var n int = 3
	var num1 [3]int
	num1[1] = n
	num1[2] = 3
	num1[0] = 0
	fmt.Println(num1)
}

//3.切片
func main2() {
	type myslice []int // 这样的字面量与数组（值）的字面量的区别也只在于最左侧的类型字面量。
	var slice1 = myslice{0, 1, 2, 3, 4}
	var slice2 = slice1[1:2]
	var size2 = cap(slice2)
	var slice3 = slice1[1:4:4]
	fmt.Println("切片slice = ", slice2, slice3, size2)

	var numbers4 = [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	slice5 := numbers4[4:6:8] //tip cap len
	fmt.Println("---slice 5 info:----",slice5, cap(slice5), len(slice5))
	length := (2)
	capacity := (4)
	fmt.Printf("__slice__:%v, %v\n", length == len(slice5), capacity == cap(slice5))

	slice5 = slice5[:cap(slice5)] // LEN 4, CAP 4
	fmt.Println(slice5, len(slice5), cap(slice5))
	slice5 = append(slice5, 11, 12, 13) //LEN 7  8
	slice5 = append(slice5, 11, 12, 13) //LEN 7  16 2倍
	length = (7)
	fmt.Println(slice5, len(slice5), cap(slice5))
	fmt.Printf("__slice__:%v\n", length == len(slice5))
	slice6 := []int{0, 0, 0}
	copy(slice5, slice6) //6-5,复制最短的个数，也就是3个
	e2 := (2)
	e3 := (2)
	e4 := (2)
	fmt.Println(slice6, slice5)
	fmt.Printf("__slice_:%v, %v, %v\n", e2 == slice5[2], e3 == slice5[3], e4 == slice5[4])
}

//4.字典 map
func floatFun() {
	mm2 := map[string]int{"golang": 42, "java": 1, "python": 8}
	mm2["scala"] = 25
	mm2["erlang"] = 50
	delete(mm2, "python")

	fmt.Printf("%d, %d, %s \n", mm2["scala"], mm2["erlang"], mm2["python"])

	py, ok := mm2["python"]
	fmt.Printf("_______x___%d, %v\n", py, ok)
}

//5.通道类型chan 可用于在不同Goroutine之间传递类型化的数据，并且是并发安全的
func complexFun() {
	ch2 := make(chan string, 2)
	// 下面就是传说中的通过启用一个Goroutine来并发的执行代码块的方法。
	// 关键字 go 后跟的就是需要被并发执行的代码块，它由一个匿名函数代表。
	// 对于 go 关键字以及函数编写方法，我们后面再做专门介绍。
	// 在这里，我们只要知道在花括号中的就是将要被并发执行的代码就可以了。
	go func() {
		ch2 <- ("已经到达~！")
	}()
	var value string = "数据: "
	defer close(ch2) //关闭为什么蔑视？
	value = value + (<-ch2)
	fmt.Println(value)

	//value2, ok := <- ch2
	//fmt.Print("val2 = ", value2, ok)
}

//6.其他通道chan
type Sender chan<- int
type Receiver <-chan int

func runeFun() {
	var myChannel = make(chan int, (0))
	var number = 6
	go func() {
		var sender Sender = myChannel
		sender <- number
		fmt.Println("Sent!")
	}()
	go func() {
		var receiver Receiver = myChannel
		fmt.Println("Received!", <-receiver)
	}()
	// 让main函数执行结束的时间延迟1秒，
	// 以使上面两个代码块有机会被执行。
	time.Sleep(time.Second)
}
