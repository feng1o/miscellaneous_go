package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", &myHandler{})
	mux.HandleFunc("/bye", sayBye) //这个直接到hadler fun

	log.Println("Starting server... v2")

	err := http.ListenAndServe(":8080", mux) //ga

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
