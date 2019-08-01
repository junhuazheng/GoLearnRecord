package learnhttp

import (
	"fmt"
	"net/http"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("handle hello")
	fmt.Fprintf(w, "hello")
}

func http1() {
	http.HandleFunc("/", Hello)
	err := http.ListenAndServe("localhost:8880", nil)
	if err != nil {
		fmt.Println("http listen failed")
	}
}