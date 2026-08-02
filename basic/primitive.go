package main

import (
	"fmt"
)

func main2() {
	var MyName int //变量声明

	fmt.Print(MyName)   //变量使用
	fmt.Println(MyName) //变量使用

	var a int = 9
	// fmt.Println(a)

	var b = a //变量推断自动类型
	_ = b     //占位符，你不用他，但必须用变量
	c := b    //:相当于声明变量，相当于var的作用，不用写var了
	a = c

	var (
		d uint16
		e int8
		f float32
		g float64
	)

	a = -5
	d = 05   //前缀0表示8进制
	a = 0o57 //0o也表示8进制，8进制只到7
	// a = 0o58 到8就报错
	a = 0xab3      // 0x前缀表示16进制0-9,a-f
	a = 5_0_123_2  //可以这样写数字
	a = 13_000_000 //13M
	f = 1.43321321321
	g = 34
	m := 43. //没声明就默认浮点64

	_, _, _, _ = d, e, f, m

	var n bool = true // 默认值是false

	fmt.Printf("a=%d, g=%f, n=%t, g=%.2f\n", a, g, n, g)
	fmt.Printf("f=%g, f=%e\n", f, f)

	fmt.Printf("f=%[1]f, g=%[2]f, g=%[2]g, f=%[1]e\n", f, g)
	//表达浮点数，可以用g，可以用f，可以用e科学计数法

}
