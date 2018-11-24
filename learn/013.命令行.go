package main

import (
	//"errors"
	"fmt"
	//"os"
	"flag"
	//"os"
)

func main() {
	style1()
	//fmt.Println(os.Args)
	methodptr := flag.String("method", "default", "method of sample")
	valueptr := flag.Int("value", -1, "valude of sample")
	flag.Parse()
	fmt.Println(valueptr, methodptr)
}

func style1() {
	var method string
	var value int

	flag.StringVar(&method, "method", "default", "method of sample")
	flag.IntVar(&value, "value", -1, "value of sample")

	flag.Parse()

	fmt.Println(method, value)
}
