package main

import (
	"fmt"
	"log"
	"os"
)

type Usb interface {
	Name() string
	Connect()
	//Del()
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
}
