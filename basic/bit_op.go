package main

import "fmt"

func main12() {
	var a uint64 = 200
	fmt.Printf("a=%b\n", a)
	//%b这种格式不输出前面的一堆零，我们后面的函数把他64位都打印出来了

	binaryFormat(a)
}

func binaryFormat(a uint64) {
	var c uint64 = 1 << 63 //最高位是1，其他位都是0
	for i := 0; i < 64; i++ {
		if c&a == c {
			fmt.Print("1")
		} else {
			fmt.Printf("0")
		}
		c = c >> 1
	}
}

// 所以整个函数就是从最高位到最低位，一个一个检查：
// 如果当前位是 1，打印 "1"
// 如果当前位是 0，打印 "0"
