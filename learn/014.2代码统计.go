package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("input err args ")
		//return
	}

	fileName := os.Args[1];
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "open file err")
		return
	}
	defer  file.Close()

	var line int8
	reader := bufio.NewReader(file)
	for {
		ctn, prifix, _ := reader.ReadLine()

		if prifix == true {
			continue
		} else {
			line++
		}
		fmt.Fprintln(os.Stdout, "", ctn,line)
		if line > 100 {
			break
		}
	}
}

func main2() {

	if len(os.Args) < 2 {
		return
	}

	fileName := os.Args[1]
	fmt.Println(fileName)
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()
	var line int
	reader := bufio.NewReader(file)
	for {
		_, isPrefix, err := reader.ReadLine()

		if err != nil {
			break
		}
		if !isPrefix {
			line++
		}
	}
	fmt.Println(line)
}
