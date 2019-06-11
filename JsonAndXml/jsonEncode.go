package main

import (
	"fmt"
	"encoding/json"
	"os"
)

type Address struct {
	City string `json:"city"`
	Area string `json:"area"`
}

type Email struct {
	Where string `json:"where"`
	Addr string `json:"addr"`
}

type Student struct {
	Id int `json:"id"`
	Address `json:"addresss"`
	Email []Email `json:"email`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
}

func main()  {
	f, err := os.Create("d:/myjson.txt")
	if err != nil {
		fmt.Println("err: ", err.Error())
		return
	}
	defer f.Close()

	//实例化对象
	stu := Student{23, Address{"guangzhou", "dongpu"}, []Email{Email{"home", "home@qq.com"}, Email{"work", "work@qq.com"}}, "junhua", "zheng"}
	fmt.Println("stu: ", stu)
	//序列化
	encoder := json.NewEncoder(f)
	err1 := encoder.Encode(stu)
	if err1 != nil {
		fmt.Println("err1: ", err1.Error())
		return
	}

	//重置文件指针
	f.Seek(0, os.SEER_SET)
	var newStu Student

	//反序列化
	decoder := json.NewDecoder(f)
	err2 := decoder.Decode(&newStu)
	if err2 != nil {
		fmt.Println("err2: ", err2.Error())
		return
	}
	fmt.Println("newStu: ", newStu)

}