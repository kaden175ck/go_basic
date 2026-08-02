package main

import "fmt"

// 切片内存共享
func slice2() {
	arr := make([]int, 3, 5)
	arr[0], arr[1], arr[2] = 2, 9, 7

	brr := arr
	brr[0] = 4

	fmt.Printf("arr=%v\n", arr)
	fmt.Printf("brr=%v\n", brr)
	fmt.Printf("arr[0]=%d, brr[0]=%d\n", arr[0], brr[0])
	fmt.Printf("len(arr)=%d cap(arr)=%d\n", len(arr), cap(arr))
	fmt.Printf("len(brr)=%d cap(brr)=%d\n", len(brr), cap(brr))
	fmt.Println("两个切片共享同一底层数组，所以修改 brr[0] 会影响 arr[0]")
}

func main18() {
	slice2()
}
