package form

import (
	"io"
	"log"
	"net/http"
	"learnhttp/form"
)

func panicProcess() {
	http.HandleFunc("/test1", logPanics(SimpleServer))
	http.HandleFunc("/test2", logPanics(FormServer))
	if err := http.ListenAndServe(":8080", nil); err != nil {
	}
}

func logPanics(handle http.HandleFunc) http.HandleFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		defer func() {
			if x := recover(); x != nil {
				log.Printf("[%v] caught panic: %v", request.RemoteAddr, x)
			}
		}()
		handle(writer, request)
	}
}