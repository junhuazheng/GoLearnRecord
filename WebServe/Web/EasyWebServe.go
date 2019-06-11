package Web

import (
	"regexp"
	"html/template"
	"strings"
	"fmt"
	"net/http"
	"strconv"
	// "log"
)

func SayHelloName(w http.ResponseWriter, r *http.Request)  {
	r.ParseForm() //解析URL中的查询字符串，并将解析结果更新到r.Form字段。
	fmt.Println(r.Form)
	fmt.Println("path: ", r.URL.Path)
	fmt.Println("scheme: ", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key: ", k)
		fmt.Println("val: ", strings.Join(v, " "))
	}
	fmt.Fprintf(w, "hello chain!")
}

func login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() //解析form
	fmt.Println("method: ", r.Method)
	if r.Method == "GET" {
		t, _ := template.ParseFiles("./view/login.ctpl")
		t.Execute(w, nil)
	}else if r.Method == "POST" {
		if len(r.Form["username"][0]) == 0 {
			fmt.Fprintf(w, "username: null of empty \n")
		}
		age, err := strconv.Atoi(r.Form.Get("age"))
		if err != nil {
			fmt.Fprintf(w, "age: The format of the input is not correct \n")
		}
		if age < 18 {
			fmt.Fprintf(w, "age: Minors are not registered \n")
		}
		if m, _ := regexp.MatchString(`^([\w\.\_]{2,10})@(\w{1,}).([a-z]{2,4})$`, r.Form.Get("email")); !m {
		fmt.Fprintf(w, "eamil: The format of the input is not correct \n")
	    }
	}
}

// func Start() {
// 	http.HandleFunc("/", SayHelloName) //注册路由请求规则
// 	err := http.ListenAndServe(":9090", nil)
// 	if err != nil {
// 		log.Fatal("ListenAndServer: ", err)
// 	}
// }