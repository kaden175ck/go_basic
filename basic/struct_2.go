package main

import "fmt"

type Residence struct {
	Province string
	City     string
}

type User2 struct {
	Id            int
	Score         float64
	Name, Address string
	Residence     Residence //一个代表变量名称，一个代表变量类型
	// Father        User2这样写的话内存会陷入无限循环
	Father *User2 //这样写可以，因为指针类型初始值是nil，不指向任何内存空间，本身就是个地址，类似一个整数，只占8个字节
}

// 匿名结构体嵌套可以直接放到上面的residence字段代替，然后就可以把type的residence删除了
// Residence struct {
// 	Province string
// 	City     string
// }

// 成员方法
func (me User2) hello2() {
	fmt.Println("My Name is", me.Name)
}

//结构体嵌套

func main7() {
	var u User2
	u.Score = 100
	u = User2{Id: 32, Address: "中国西安", Name: "Shy"}
	fmt.Println(u.Name)
	u = User2{32, 99, "Kaden", "BJ", Residence{"shanxi", "xian"}, &User2{}}
	// father节点也可以是nil
	//然后这里也可以用匿名结构体，直接删除residence这个词，直接放struct{Province string，City string}{"shanxi", "xian"}
	// reisidence赋值你也可以u.Residence.Province="" 这样子也行

	fmt.Printf("%v\n", u)
	fmt.Printf("%+v\n", u)
	fmt.Printf("%#v\n", u)

	u.hello2()
}

// 2叉树结构体嵌套
type TreeNode struct {
	Data       int
	LeftChild  *TreeNode
	RightChild *TreeNode
}
