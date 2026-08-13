package main

import (
	"fmt"
	"math/rand/v2"
	"strings"
)

var letterCollection = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ/_+石头")

// 生成长度为n的随机字符串
// 注释掉的是使用切片的实现，理解简单，但转 string 时会多一次内存拷贝
// strings.Builder （当前用的）：更推荐，性能更好，内存分配更少
func RandString(n int) string {
	// rect := make([]rune, 0, n)
	sb := strings.Builder{}
	for i := 0; i < n; i++ {
		index := rand.IntN(len(letterCollection))
		// rect = append(rect, letterCollection[index])
		sb.WriteRune(letterCollection[index])
	}
	// return string(rect)
	return sb.String()
}

func main40() {
	fmt.Println(RandString(10))
	fmt.Println(RandString(10))
	fmt.Println(RandString(10))
	fmt.Println(RandString(10))
}
