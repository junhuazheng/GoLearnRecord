package client

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func client() {
	res, err := http.Get("https://www.baidu.com/")
	if err != nil {
		fmt.Println("get err", err)
		return
	}

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("get data err: ", err)
		return
	}

	fmt.Println(data)
}