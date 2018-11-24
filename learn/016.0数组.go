package main

import (
	"fmt"
)

func main() {
	fmt.Println("_____________________begin_x________________")
	var a [2]int
	var b [1]int
	//a = b
	for i := 0; i < len(a); i++ {
		a[i] = i
		fmt.Println(b)
	}

	c := [...]int{100: 100}
	fmt.Println(c)

	d := [...]int{0: 1, 1: 2}
	fmt.Println(d)

	x := [...]int{1, 2, 3, 45}
	fmt.Println(x)

	var p *[4]int = &x
	fmt.Println(p)

	str := [5]string{4: "4444"}
	fmt.Println(len(str), str[4])

	e := [...][3]int{
		{1, 2, 3},
		{0: 3, 1: 2, 2: 1},
	}
	fmt.Println(e)

	var k = [...]int{5, 4, 3, 2, 1}
	fmt.Println(k)
	for i := 0; i < len(k); i++ {
		for j := i; j < len(k); j++ {
			if k[i] > k[j] {
				tmp := k[i]
				k[i] = k[j]
				k[j] = tmp
			}
		}
	}
	fmt.Println(k)
}
