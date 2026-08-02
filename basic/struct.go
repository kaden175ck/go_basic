package main

import "fmt"

type User struct {
	Id            int
	Score         float64
	Name, Address string
}

func main5() {
	var u User
	u = User{Id: 32, Score: 99, Name: "shy-1", Address: "BJ"}

	fmt.Println(u.Name)
	// u = User{32, 99, "shy", "BJ"}
	// 也可以直接这样赋值
	// 这种声明，需要按结构体当时定义的顺序，并且必须给全员赋值

	u.hello()

	//匿名结构体---结构体类型的变量，和上面的u一样
	//匿名结构体通常只在这个函数里使用一次
	var student struct {
		Name string
		Age  int
	}
	student.Name = "shy-2"
	student.Age = 26
	_ = student.Age

	u.Name = student.Name
	fmt.Println(u.Name)

	u2 := User{} //都没有赋值
	u2.Address = "上海"
	fmt.Println(u2.Address)

	u3 := &u2
	fmt.Println(u3.Address)

	u4 := new(User) //u4也是*user的指针类型,new先创建空的结构体再返回指针
	u4.Name = "shihaoyang"
	fmt.Println(u4.Name)

}

// 成员方法
func (me User) hello() {
	fmt.Println("My Name is", me.Name)
}
