package main

import (
	"flag"
	"fmt"
)

var (
	Eshost  *string = flag.String("host", "localhost", "Elasticsearch Server Host Address")
	Eshost2 *string = flag.String("test", "test", "xxxxxxxxxxxxxx")
)

//go run 013.1测试flag.go  -host 122.2.23.3
func main() {
	flag.Parse()
	fmt.Println(*Eshost)
	fmt.Println(*Eshost2)
}
