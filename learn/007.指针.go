package main

import "fmt"

type MyInt struct {
	n int
}

func (myInt *MyInt) Increase() {
	myInt.n++
}

func (myInt *MyInt) Decrease() {
	myInt.n--
}

func main() {
	mi := MyInt{}
	mi.Increase()
	mi.Increase()
	mi.Decrease()
	mi.Decrease()
	mi.Increase()
	fmt.Printf("%v\n", mi.n == 1)
	main2()
}

type Pet interface {
	Name() string
	Age() uint8
}
type Dog struct {
	name string
	age  uint8
}

func (d Dog) Name() string {
	return d.name
}
func (d Dog) Age() uint8 {
	return d.age
}
func main2() {
	myDog := Dog{"Little D", 3}
	_, ok1 := interface{}(&myDog).(Pet)
	_, ok2 := interface{}(myDog).(Pet)
	fmt.Printf("%v, %v\n", ok1, ok2)
}
