package main

import "fmt"

type Person struct {
	Name    string
	Gender  string
	Age     uint8
	Address string
}

/*
1.等给结构体一个方法，
2.必须指针
3.move方法依附于Person
*/
func (per *Person) Move(newAddress string) string {
	oldAddress := per.Address
	per.Address = newAddress
	return oldAddress
}

func main() {
	p := Person{"Robert", "Male", 33, "Beijing"}
	oldAddress := p.Move("San Francisco")
	fmt.Printf("%s moved from %s to %s.\n", p.Name, oldAddress, p.Address)
}
