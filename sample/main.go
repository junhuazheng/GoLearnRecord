package main

import (
	"log"
	"os"

	 _"github.com/goinaction/code/chapter2/sample/matchers"
	 "github.com/goinactin/code/chapter2/sample/search"
)

func init() {
	//将日志输出到标准输出
	log.SetOutput(os.Stdout)
}

func main() {
	//使用特定的项做搜索
	search.Run("president")
}