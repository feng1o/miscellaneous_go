package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"time"
)

// array: is value copy
//
//

type arry struct {
}

func (*arry)InitArr() {
	var a [3]int                    // 定义长度为3的int型数组, 元素全部为0
	var b= [...]int{1, 2, 3}       // 定义长度为3的int型数组, 元素为 1, 2, 3
	var c= [...]int{2: 3, 1: 2}    // 定义长度为3的int型数组, 元素为 0, 2, 3
	var d= [...]int{1, 2, 4: 5, 6} // 定义长度为6的int型数组, 元素为 1, 2, 0, 0, 5, 6

	// print
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)

	// for range
	for i,v := range d {
		fmt.Println(i, v)
	}
}

// 字符串数组
var s1 = [2]string{"hello", "world"}
var s2 = [...]string{"你好", "世界"}
var s3 = [...]string{1: "世界", 0: "你好", }

// 结构体数组
var line1 [2]image.Point
var line2 = [...]image.Point{image.Point{X: 0, Y: 0}, image.Point{X: 1, Y: 1}}
var line3 = [...]image.Point{{0, 0}, {1, 1}}

// 图像解码器数组
var decoder1 [2]func(io.Reader) (image.Image, error)
var decoder2 = [...]func(io.Reader) (image.Image, error){
	png.Decode,
	jpeg.Decode,
}

// 接口数组
var unknown1 [2]interface{}
var unknown2 = [...]interface{}{123, "你好"}

// 管道数组
var chanList = [2]chan int{}

// 我们还可以定义一个空的数组：

var d [0]int       // 定义一个长度为0的数组
var e = [0]int{}   // 定义一个长度为0的数组
var f = [...]int{} // 定义一个长度为0的数组

// 长度为0的数组在内存中并不占用空间。空数组虽然很少直接使用，但是可以用于强调某种特有类型的操作时避免分配额外的内存空间，比如用于管道的同步操作：
func emptyChannel() {
	c1 := make(chan [0]int)
	go func() {
		fmt.Println("put empty to channel : c1")
		c1 <- [0]int{}
	}()
	<-c1
	//在这里，我们并不关心管道中传输数据的真实类型，其中管道接收和发送操作只是用于消息的同步。对于这种场景，用空数组来作为管道类型可以减少管道元素赋值时的开销。一般更倾向于用无类型的匿名结构体代替：

	c2 := make(chan struct{})
	go func() {
		fmt.Println("c2")
		c2 <- struct{}{} // struct{}部分是类型, {}表示对应的结构体值
	}()
	<-c2
}




func main()  {
	// 1. init
	var arr = arry{}
	arr.InitArr()

	// 2. empty array
	emptyChannel()

	time.Sleep(10*time.Second)

}