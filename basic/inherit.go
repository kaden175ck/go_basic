package main

import "fmt"

func main6() {
	type User struct {
		Name string
		Age  int
	}

	type NewUser struct {
		NewAge int
	}

	type Video struct {
		Length  int
		Name    string
		Author  User //用变量类型来充当变量名称
		NewUser      //匿名成员，没有名称只有类型
	}

	u := User{Name: "shihaoyang", Age: 25}
	u2 := NewUser{NewAge: 26}
	v := Video{
		Length: 120,
		Name:   "go语言教程",
		Author: u, //注意：行尾一定要加逗号
		// NewMe:  u2,
		NewUser: u2,
	}

	fmt.Println(v.Length)
	fmt.Println(v.Name)        //访问自己的Name
	fmt.Println(v.Author.Name) //访问“父类”的Name
	fmt.Println(v.Author.Age)  // 从User里“继承”了Age
	fmt.Println(v.NewAge)
}
