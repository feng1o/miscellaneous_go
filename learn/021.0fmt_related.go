package main

import (
	"fmt"
)

type Project struct {
	Name string
	Sn [5]byte
	SnUne [5]rune
}

func (p Project)String() string {
	return fmt.Sprintf("%v, %v, %v", p.Name, p.Sn, p.SnUne);
}

type IPAddr [4]byte

func (ip IPAddr) String() string {
	return fmt.Sprintf("%v.%v.%v.%v", ip[0], ip[1], ip[2], ip[3])
}

func Test(){
	addrs := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for n, a := range addrs {
		fmt.Printf("%v: %v\n", n, a)
	}
}

/*
1、Go语言支持两种字符类型：byte代表UTF-8，rune代表Unicode
2、根据字符串下标取字符，类型为byte（中文在UTF-8中占3个字节，而不是一个）
 */
func BytesLenRange()  {
	var str string = "1234五"
	for i := 0; i < len(str); i++ {
		fmt.Println(i, str[i])
	}
	fmt.Println("-----------------------------")
	for i, ch :=  range(str) {
		fmt.Println(i, ch)
	}
}

func main()  {
	var p Project
	p.Name = "fk"
	p.Sn[0] = '1'
	p.SnUne[0] = '五'
	fmt.Println("%v", p)

	Test()
	BytesLenRange()
}

