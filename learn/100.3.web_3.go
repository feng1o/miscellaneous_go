package main

import (
	"io"
	"log"
	"net/http"
	"time"
)

var mux = make(map[string]func(w http.ResponseWriter, r *http.Request))

func main() {
	server := &http.Server{
		Addr:              ":8080",
		WriteTimeout:      time.Second * 2, // 4
		Handler:           &myHandler{},
		ReadHeaderTimeout: time.Second * 4,
	}

	//var mux = make(map[string]func(w http.ResponseWriter, r *http.Request))
	mux["/hello"] = sayHello
	mux["/bye"] = sayBye

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
	/*
		quit := make(chan os.Signal)
		signal.Notify(quit, os.Interrupt)

		mux := http.NewServeMux()
		mux.Handle("/", &myHandler{})
		mux.HandleFunc("/bye", sayBye)

		go func() {
			<-quit

			if err := server.Close(); err != nil {
				log.Fatal("Close server:", err)
			}
		}()

		server.Handler = mux
		log.Print("Starting server... v3")
		err := server.ListenAndServe()
		if err != nil {
			if err == http.ErrServerClosed {
				log.Print("Server closed under request")
			} else {
				log.Fatal("Server closed unexpected")
			}
		}
	*/
}

type myHandler struct{}

/*
func (*myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "通过mux map自己路由 " + r.URL.String())
	io.WriteString(w, "\n")
	w.Write([]byte("Hello v3-1, the request URL is: " + r.URL.String()))
}*/

//这个才能根据mux自己路由
func (*myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if h, ok := mux[r.URL.String()]; ok {
		h(w, r) //h就应该是sayhello， 或者saybye
		return
	}
	//不存在就会走默认下面的
	io.WriteString(w, "通过mux map自己路由 "+r.URL.String())
	io.WriteString(w, "\n")
	w.Write([]byte("Hello v3-1, the request URL is: " + r.URL.String()))
}

func sayBye(w http.ResponseWriter, r *http.Request) {
	//time.Sleep(3 * time.Second)
	//w.Write([]byte("Bye bye this is version 3-1!"))
	io.WriteString(w, "Bye bye this is version 3-1!")
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	//time.Sleep(3 * time.Second) //sleep error
	//w.Write([]byte("Bye bye this is version 3-1!"))
	io.WriteString(w, "xxx hello   \n")
}
