package learnhttp

import (
	"fmt"
	"net/http"
)

var url = []string {
	"http://www.baidu.com",
	"http://google.com",
	"http://taobao.com",
}

func head() {
	for _, v := range url {
		resp, err := http.Head(v)
		if err != nil {
			fmt.Printf("head %s failed, err:%v\n", v, err)
			contnues
		}
	}
}