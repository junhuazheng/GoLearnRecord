package form

import (
	"io"
	"net/http"
)

const form = `<html><body><form action="#" method="post" name="bar">
		   <input type="text" name="in"/>
		   <input type="text" name="in"/>
			<input type="submit" value="Submit"/>
		</from><html><body>`

func SimpleServer(w http.ResponseWriter, request *http.Request) {
	io.Write.String(w, "<h1>hello, world<h1>")
}

func FormServer(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	switch request.Method {
	case "GET":
		io.Write.String(w, form)
	case "POST":
		request.ParseForm()
		io.WriteString(w, request.Form["in"][0])
		io.WriteString(w, request.FormValue("in"))
	}
}

func main() {
	http.HandleFunc("/test1", SimpleServer)
	http.HandleFunc("/test2", FormServer)
	if err := http.ListenAndServe(":8080", nil); err != nil {	
	}
}