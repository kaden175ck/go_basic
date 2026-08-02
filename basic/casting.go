package main

import (
	"fmt"
	"math"
)

func main20() {
	cast()
}

func cast() {
	fmt.Println("1. 整数转整数：可能发生截断")
	var big uint64 = 300
	small := int8(big)
	fmt.Printf("uint64=%d -> int8=%d\n", big, small)

	big = math.MaxUint64
	signed := int64(big)
	unsigned32 := uint32(big)
	fmt.Printf("uint64=%d -> int64=%d\n", big, signed)
	fmt.Printf("uint64=%d -> uint32=%d\n", big, unsigned32)
	fmt.Println()

	fmt.Println("2. rune 和 int：字符本质上是整数")
	ch := '中'
	code := int(ch)
	fmt.Printf("%q -> %d\n", ch, code)
	fmt.Printf("%d -> %q\n", code, rune(code))
	fmt.Println()

	fmt.Println("3. byte 和 int：本质也是整数，只是范围不同")
	var b byte = 65
	n := int(b)
	fmt.Printf("byte=%d -> int=%d\n", b, n)
	fmt.Printf("int=%d -> byte=%d\n", n, byte(n))
	fmt.Println()

	fmt.Println("4. float 和 int：小数部分会被直接丢掉")
	var f float64 = 12.99
	i := int(f)
	fmt.Printf("float64=%.2f -> int=%d\n", f, i)
	f = 9.01
	fmt.Printf("float64=%.2f -> int=%d\n", f, int(f))
	fmt.Println()

	fmt.Println("5. 转换不会自动帮你做逻辑判断")
	fmt.Println("比如 int 和 bool 不能直接互转，必须自己写条件判断")
}
