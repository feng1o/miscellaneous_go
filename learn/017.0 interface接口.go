package main

import (
	"fmt"
	"log"
)

type Person struct {
	Name string
	Age  int
}

func (p *Person) GetName() string {
	fmt.Println(p.Name)
	return p.Name
}

func (p *Person) GetAge() int {
	fmt.Println(p.Age)
	return p.Age
}

func main() {
	log.Println("----")
	fmt.Println("test----\n")
	var p Person
	p.Name = "test0"
	p.Age = 100
	p.GetAge()
	p.GetName()
}
