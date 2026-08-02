package main

// 切片的内存共享和内存分离的情况
import "fmt"

func main19() {
	arr := [3]int{1, 2, 3}
	fmt.Printf("arr: %v len=%d cap=%d\n", arr, len(arr), cap(arr))
	fmt.Printf("&arr[0]=%p &arr[1]=%p &arr[2]=%p\n", &arr[0], &arr[1], &arr[2])

	brr := arr[1:2]
	fmt.Printf("before append:\n")
	fmt.Printf("arr: %v\n", arr)
	fmt.Printf("brr: %v len=%d cap=%d\n", brr, len(brr), cap(brr))
	fmt.Printf("&brr[0]=%p\n", &brr[0])

	brr = append(brr, 8)
	fmt.Printf("after first append:\n")
	fmt.Printf("arr: %v\n", arr)
	fmt.Printf("brr: %v len=%d cap=%d\n", brr, len(brr), cap(brr))
	fmt.Printf("&brr[0]=%p &brr[1]=%p\n", &brr[0], &brr[1])

	brr = append(brr, 7)
	fmt.Printf("after second append (may detach):\n")
	fmt.Printf("arr: %v\n", arr)
	fmt.Printf("brr: %v len=%d cap=%d\n", brr, len(brr), cap(brr))
	fmt.Printf("&brr[0]=%p &brr[1]=%p &brr[2]=%p\n", &brr[0], &brr[1], &brr[2])
}
