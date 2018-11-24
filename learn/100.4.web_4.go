package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	os.Open("xxxx.log")
	f2, _ := os.Create("create.txt")

	log.SetOutput(f2)
	log.SetPrefix("prefix  ") // prefix

	mux := http.NewServeMux()

	mux.Handle("/", &myHandler{})
	mux.HandleFunc("/bye", sayBye) //这个直接到hadler fun

	log.Println("Starting server... v2")

	wd, err := os.Getwd() //路径
	if err != nil {
		log.Fatal(err)
	}

	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer((http.Dir(wd))))) //这个能获得静态文件列表

	err = http.ListenAndServe(":8080", mux)
	log.Fatal(err)
}

func sayBye(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Bye bye this is version 2!"))
}

type myHandler struct{}

func (*myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "URL"+r.URL.String()) //
	//w.Write([]byte("Hello v2, the request URL is: " + r.URL.String()))
}
