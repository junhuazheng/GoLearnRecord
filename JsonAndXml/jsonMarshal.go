package main

import (
	"fmt"
	"encoding/json"
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
	//实例化对象
	stu := Student{23, Address{"guangzhou", "dongpu"}, []Email{Email{"home", "home@qq.com"}, Email{"work", "work@qq.com"}}, "junhua", "zheng"}
	fmt.Println("stu: ", stu)
	//序列化对象到字符串
	buf, err := json.Marshal(stu)
	if err != nil {
		fmt.Println("err: ", err.Error())
		return
	}
	fmt.Println("json: ", string(buf))

	//反序列化字符串到对象
	var newStu Student
	err1 := json.Unmarshal(buf, &newStu)
	if err1 != nil {
		fmt.Println("err1: ", err1.Error())
		return
	}
	fmt.Println("newStu: ", newStu)
}
