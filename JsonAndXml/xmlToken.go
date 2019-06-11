package main

import (
	"fmt"
	"encoding/xml"
	"os"
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
	t2()
}

func t2() {
	f, err := os.Create("d:/myxml.xml")
	if err != nil {
		fmt.Println("err: ", err.Error())
		return
	}
	defer f.Close()

	//实例化对象
	stu :=Student{23, Address{"guangzhou", "dongpu"}, []Email{Email{"home", "home@qq.com"}, Email{"work", "work@qq.com"}}, "junhua", "zheng"}
	fmt.Println("stu: ", stu)
	//序列化到文件里面
	encoder := xml.NewEncoder(f)
	err1 := encoder.Encode(stu)
	if err1 != nil {
		fmt.Println("err1: ", err1.Error())
		return
	}
	//重置文件指针
	f.Seek(0, os.SEEK_SET)
	decoder := xml.NewEncoder(f)
	var strName string
	for {
		token, err2 := decoder.Token()
		if err2 != nil {
			break
		}
		switch t:= token.(type) {
		case xml.StartElement:
			stelm := xml.StartElement(t)
			fmt.Println("start: ", stelm.Name.Local)
			strName = stelm.Name.Local
		case xml.EndElement:
			endelm := xml.EndElement(t)
			fmt.Println("end: ", endelm.Name.Local)
		case xml.CharData:
			data := xml.CharData(t)
			str := string(data)
			switch strName {
			case "City":
				fmt.Println("city: ", str)
			case "first":
				fmt.Println("first: ", str)
			}
		}
	}
}