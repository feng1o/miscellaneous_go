package main

import (
	"fmt"
	"reflect"
)

const name = "test"
//const (
//	a = iota
//	_      //跳值使用
//	b = iota
//	c = "inset value" //插值
//	d = iota
//)

type  ByteSize int64

const (
	_           = iota                   // ignore first value by assigning to blank identifier
	KB ByteSize = 1 << (10 * iota) // 1 << (10*1)
	MB                                   // 1 << (10*2)
	GB                                   // 1 << (10*3)
	TB                                   // 1 << (10*4)
	PB                                   // 1 << (10*5)
	EB                                   // 1 << (10*6)
	ZB                                   // 1 << (10*7)
	YB                                   // 1 << (10*8)
)

const (
	a,b = iota, iota+3
	c,d
	e = iota
)
func main() {
	fmt.Println(name)
	fmt.Println(reflect.TypeOf(name))
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)

}
