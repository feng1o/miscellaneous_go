package main

import (
	"fmt"
	"log"
	"os"
)

type Usb interface {
	Name() string
	//Connect()
	Connect //这是一接口，被usb嵌套了
	//Del()
}

type Connect interface {
	Connect()
}

type PhoneConnect struct {
	name string
}

func (pc PhoneConnect) Name() string {
	fmt.Println(pc.name)
	return pc.name
}

func (pc PhoneConnect) Connect() {
	fmt.Println("connect : ", pc.name)
}

func Disconnect(usb Usb) {
	fmt.Println("disconnect : ", usb.Name())
	if pc, ok := usb.(PhoneConnect); ok {
		fmt.Println("是phone", pc.name)
	}
}

/*
 空interface，是任何的接口都可以作为其参数了
*/
func DisconnectAll(usb interface{}) {
	if pc, ok := usb.(PhoneConnect); ok {
		fmt.Println("是phone", pc.name)
	}
	//通过type来取其类型---
	switch v := usb.(type) {
	case PhoneConnect:
		fmt.Println("phone ----", v.name)
	default:
		fmt.Println("unkown")
	}
}

func main() {
	file, _ := os.OpenFile("log.log", os.O_WRONLY|os.O_APPEND|os.O_CREATE, os.ModePerm)
	defer file.Close()
	log.SetOutput(file)
	log.SetPrefix("_________")
	//fmt.Println(log.Prefix())
	//log.SetOutput(os.Stdout)
	log.Println("--- start")
	fmt.Println("------print-----")
	//fmt.Println(fmt.Sprintf("ReadRequest error: "))
	log.Println("--- end\n\n")

	var a Usb //这里是usb
	a = PhoneConnect{"phoneconnect"}
	a.Connect()
	Disconnect(a)
	DisconnectAll(a)

	fmt.Println("--------------------接口互相转换------------------")
	var b Connect
	b = Connect(a) //强行吧 phoneconnec转换为了COnnect
	b.Connect()
}
