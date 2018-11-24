package main

import (
	"encoding/binary"
	"fmt"
	"os"
)

type BitMapInfoHeader struct {
	Size           uint32
	width          int32
	height         int32
	places         uint16
	bitcount       uint16
	compressiion   uint32
	sizeimage      uint32
	xperlspermeter int32
	yperlspermeter int32
	clsruse        int32
	clsrimport     int32
}

func main() {
	file, err := os.Open("image.bmp")
	if err != nil {
		fmt.Println("error")
		panic(fmt.Sprintf("err----open file"))
	}

	var heada, headb byte
	binary.Read(file, binary.LittleEndian, &heada)
	binary.Read(file, binary.LittleEndian, &headb)

	var sizea uint32
	binary.Read(file, binary.LittleEndian, &sizea)

	var reserveda, reservedb uint16
	binary.Read(file, binary.LittleEndian, &reserveda)
	binary.Read(file, binary.LittleEndian, &reservedb)

	fmt.Printf("_____%c,%c\n", heada, headb)

	//用结构体，让其自动填充
	infoheader := new(BitMapInfoHeader)
	binary.Read(file, binary.LittleEndian, &infoheader)

	defer file.Close()
	fmt.Println(infoheader)
}
