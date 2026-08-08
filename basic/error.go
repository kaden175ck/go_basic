package main

import (
	"errors"
	"fmt"
)

var (
	ErrNotFound = errors.New("not found error") //error变量名一般以Err开头
	ErrServer   = errors.New("server error")
)

// 自定义Error
type MyError struct {
	Name string
	Code int
	Desc string
}

// 构造函数，返回指针节省内存空间
func NewMyError(name string, code int, desc string) *MyError {
	return &MyError{
		Name: name,
		Code: code,
		Desc: desc,
	}

	//另外一种构造方式和上面一样效果
	// err:=new(NewMyError)
	// err.Name=name
	// err.Code=code
	// err.Desc=desc
	// return err
}

// 拥有Error() string即实现了error接口。Print error时默认会调用error的Error()方法
// 如果把这里加一个指针符号*MyError，那就是结构体指针实现了error方法，不是结构体本身，所以到时候第51行要加取地址符号
// 否则就是第58行，不用加&符号，直接传实例就好
func (e MyError) Error() string {
	return fmt.Sprintf("[%d]%s: %s", e.Code, e.Name, e.Desc)
}

// Print对象时默认会调用对象的String()方法。error是个例外
// 他会找实现error接口的方法跑
// func (e MyError) String() string {
// 	return e.Name
// }

// 函数有多个返回值时，error通常是最后一个
func divide(a, b int) (int, error) {
	if b == 0 {
		// return 0, &MyError{
		// 	Name: "math",
		// 	Code: 101,
		// 	Desc: "divide by zero",
		// }

		// return 0, MyError{}
		return 0, NewMyError("math", 101, "divide by zero")
		// 这个函数本身就返回的是一个指针，所以如果第39行当时加了指针func (e *MyError) Error() string {}用59行也可以
		// return 0, ErrNotFound  这两个很好理解
		// return 0, ErrServer 这两个很好理解
		// return 0, fmt.Errorf("divide error %d %d", a, b)
	} else {
		return a / b, nil
	}
}

func main32() {
	err2 := NewMyError("math", 102, "divide by zero")
	fmt.Printf("%s\n", err2)

	c, err := divide(5, 0)
	if err != nil {
		fmt.Printf("出错 %s\n", err) //默认会调用err.Error()
		// fmt.Printf("出错 %s\n", err.Error())
	} else {
		fmt.Printf("结果 %d\n", c)
	}
}
