package main

// 常量是不能改变的，常量可以不使用，会有warning可以忽略
// var必须使用
func main3() {
	const PI float32 = 3.14
	_ = PI

	const (
		E = 2.61
		// PI = 3.14
	)

	// 常量必须赋值，但是这种情况会依次加一
	//奇葩写法，iota按顺序递增，也可以通过一些变体自定义
	const (
		a = iota //0
		b        //1
		c        //2
		d        //3
	)

	const (
		NOT_USE = iota // iota =0
		// NOT_USE = 1 << (10*iota) // iota = 0
		KB = 1 << (10 * iota) // iota=1
		MB = 1 << (10 * iota) //iota=2
		GB = 1 << (10 * iota) //iota =3
	)
	// 还可以这样写,利用iota的递增性
	// const (
	// 	NOT_USE = 1 << (10*iota)
	// 	KB
	// 	MB
	// 	GB
	// )

	const (
		ss, mm = iota + 1, iota + 2 //1,2  iota=0
		kk, qq = iota + 1, iota + 2 //2,3  iota=1                     //2,3  iota=1
		ee, ff                      //3,4  iota=2
	)
}
