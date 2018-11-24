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
