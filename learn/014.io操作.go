package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type Data struct {
}

//没有初始化回叼用
func (self Data) String() string {
	return "data"
}

func ReadFrom(reader io.Reader, num int) ([]byte, error) {
	p := make([]byte, num)
	n, err := reader.Read(p)
	if n > 0 {
		return p[:n], nil
	}
	return p, err
}

//read from stdin
func sampleReadFromString() {
	data, _ := ReadFrom(strings.NewReader("from string"), 12)
	fmt.Println("read,data = %s", data)
}

//from file
func sampleReadFile() {
	file, _ := os.Open("000.hello.go")
	defer file.Close()

	data, _ := ReadFrom(file, 10)
	fmt.Println("read from file ", data)

}
func sampleReadStdin() {
	fmt.Println("input fromt stdin:")

	data, _ := ReadFrom(os.Stdin, 11)
	fmt.Print(data)
}

//缓存io， buffio
func main() {
	fmt.Fprintln(os.Stdout, "A\n")

	fmt.Printf("%v\n", Data{})

	fmt.Println("________________reader")
	sampleReadFromString()
	sampleReadFile()
	sampleReadStdin()
}
