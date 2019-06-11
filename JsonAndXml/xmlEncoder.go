package main

import (
	"os"
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

func main() {
	f, err := os.Create("d:/myxml.xml")
	if err != nil {
		fmt.Println("err: ", err.Error())
		return
	}
	defer f.Close()

	//实例化对象
	stu :=Student{23, Address{"guangzhou", "dongpu"}, []Email{Email{"home", "home@qq.com"}, Email{"work", "work@qq.com"}}, "junhua", "zheng"}
	fmt.Println("stu: ", stu)
	//序列化到文件中
	encoder := xml.NewEncoder(f)
	err1 := encoder.Encode(stu)
	if err1 != nil {
		fmt.Println("err1: ", err1.Error())
		return
	}
	//重置文件指针
	f.Seek(0, os.SEEK_SET)
	var newStu Student
	//反序列化到newStu对象
	decoder := xml.NewDecoder(f)
	err2 := decoder.Decode(&newStu)
	if err2 != nil {
		fmt.Println("err2: ", err2.Error())
		return
	}
	fmt.Println("newStu: ", newStu)
}
