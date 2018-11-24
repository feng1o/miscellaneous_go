package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	strReader := strings.NewReader("hello world")

	buffReader := bufio.NewReader(strReader)

	//peek不减少
	data, _ := buffReader.Peek(5) //peek后，buffer还在里面，
	fmt.Println(data)
	fmt.Println((string)(data))
	fmt.Println("统计多少个字符", (string)(data), buffReader.Buffered()) //缓存了多少 11个，全部

	//抓取，回少
	str, _ := buffReader.ReadString('r') //取了buffer就少了S
	fmt.Println(str, buffReader.Buffered())

	//写入
	w := bufio.NewWriter(os.Stdout)
	fmt.Fprint(w, "\t写入：hello ")
	fmt.Fprint(w, "world")
	w.Flush() //让输出S
}
