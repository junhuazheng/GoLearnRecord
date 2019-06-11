package main

import (
	"fmt"
	"encoding/xml"
)

type Address struct {
	City string
	Area string
}

type Email struct {
	Where string `xml:"where,attr"`
	Addr string
}

type Student struct {
	Id int `xml:"id,attr"`
	Address
	Email []Email
	FirstName string`xml:"name>first"`
	LastName string `xml:"name>last"`
}

func main()  {
	//实例化独享
	stu :=Student{23, Address{"guangzhou", "dongpu"}, []Email{Email{"home", "home@qq.com"}, Email{"work", "work@qq.com"}}, "junhua", "zheng"}
	fmt.Println("stu: ", stu)
	//序列化
	buf, err := xml.Marshal(stu)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("xml: ", string(buf))

	var newStu Student
	//反序列化
	err1 := xml.Unmarshal(buf, &newStu)
	if err1 != nil {
		fmt.Println(err1.Error())
		return
	}
	fmt.Println("newStu: ", newStu)
}