package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"
	"regexp"
	"github.com/lytics/logrus"
)

type Message struct {
	TimeLocal       time.Time
	ByteSends       int
	Content, Method string
	RequestTime     float64
}

type Reader interface {
	Read(rc chan []byte)
}

type Writer interface {
	Write(wc chan string)
}

type LogProcess struct {
	rc    chan []byte
	wc    chan string
	read  Reader
	write Writer
}

var line_end int = 0

type ReadFromFile struct {
	path string
}

func (r *ReadFromFile) Read(rc chan []byte) {
	//read file
	//f, err := os.OpenFile(r.path, os.O_RDWR, 0755)
	f, err := os.Open(r.path)
	if err != nil {
		//lg.Panic("open file err")
		panic(fmt.Sprintf("open file err"))
	}
	defer f.Close()

	//read line
	//f.Seek(0, 2)
	f.Seek(0,0)
	rd := bufio.NewReader(f)
	for {
		line, err := rd.ReadBytes('\n')
		if err == io.EOF {
			time.Sleep(1 * time.Second)
			line_end += 1
			fmt.Println("   %d    end of file", line_end)
		} else if err != nil {
			log.Panic("read file err")
		}
		fmt.Println(" line : ", (string)(line))
		//rc <- line[:len(line)-1]
		rc <- line
	}

}

type WriteToInfluxDb struct {
	influxDBsn string //influxdb dtat source
}

func (w *WriteToInfluxDb) Write(wc chan string) {
	for ch := range wc {
		fmt.Println(ch)
	}
}

func (l *LogProcess) Process() {
	/**
		_________2018/05/01 21:31:28 --- start

		([_]+)(\d{4}/\d{2}/\d{2}\s+[\d{2}:]+)\s+(---)\s([\w]+)
	**/
	loc, _ := time.LoadLocation("Asia/ShngHai")
	r := regexp.MustCompile(`([_]+)(\d{4}/\d{2}/\d{2}\s+[\d{2}:]+)\s+(---)\s([\w]+)`)
	for v := range l.rc {
		ret := r.FindStringSubmatch(string(v))
		log.Println("   ret size  %d", len(ret))
		if len(ret) != 4 {
			logrus.Println("find string sub mast fille", string(v))
			continue
		}

		message := Message{}
		t, err := time.ParseInLocation("2018/05/01 21:31:28", ret[1], loc)
		if err != nil {
			log.Println(" pars time fail")
		}
		message.TimeLocal = t
		message.Content = string(ret[3])

		l.wc <- strings.ToUpper(string(v))
	}
}

func main() {
	fmt.Println("-------------begin------------------F")
	/*
	var logger = lg.New(&bytes.Buffer{}, "\t", lg.Lshortfile)
	logger.SetOutput(os.Stdout)
	logger.Println(runtime.NumCPU())
	logger.Println("start process log files :F")
	*/
	log.New(&bytes.Buffer{}, " prefix:", log.Lshortfile)
	log.SetOutput(os.Stdout)
	log.Println("start process log files :F")

	w := &WriteToInfluxDb{
		influxDBsn: "joliu&feng123",
	}

	r := &ReadFromFile{
		path: "log.log",
	}
	lp := &LogProcess{
		rc:    make(chan []byte),
		wc:    make(chan string),
		read:  r,
		write: w,
	}

	go lp.read.Read(lp.rc)
	go lp.Process()
	go lp.write.Write(lp.wc)

	time.Sleep(1000 * time.Second)
	fmt.Println("----------------end---------------F")
}
