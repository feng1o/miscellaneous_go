package main

import (
	"log"
	"bytes"
	"os"
	"fmt"
	"time"
)

type IFile interface {
	OpenFile()
	ReadFile()
}

type LFile struct {
	path string
	fd   os.File
	len  int64
}

func (lf *LFile) OpenFile() {
	fmt.Println("1.start open file")
	fd, err := os.Open(lf.path)
	if err != nil {
		panic("open file err \n");
	}
	lf.fd = *fd
	len, err := fd.Seek(0, 2)
	log.Println("size = %d", len)

	lf.len = len
}

func (lf *LFile) ReadFile() {
	fmt.Println("2.read file")
	buff := make([]byte, lf.len)
	lf.fd.Seek(0, 0)
	defer lf.fd.Close()
	n, err := lf.fd.Read(buff)
	if err != nil {
		fmt.Println("size = %d , len =%d ", n, lf.len)
	}
	fmt.Printf("file content = %s", string(buff))
	time.Sleep(1 * time.Second)
}

func main() {
	var logger = log.New(&bytes.Buffer{}, "", log.Lshortfile)
	logger.SetOutput(os.Stdout)
	logger.Println("----create file-----")

	lf := &LFile{
		path: "log.log",
	}

	lf.OpenFile()
	lf.ReadFile()

}
