package Web

import (
	"fmt"
	"net/http"
	"log"
)

type MyMux struct {
}

func (p *MyMux)ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		SayHelloName(w, r)
		return
	}
	if r.URL.Path =="/about" {
		about(w, r)
		return
	}
	if r.URL.Path == "/login" {
		login(w, r)
		return
	}
	http.NotFound(w, r)
	return
}

func about(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "I am chain, from guangzhou")
}

func Start() {
	mux := &MyMux{}
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}