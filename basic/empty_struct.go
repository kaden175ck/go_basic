package main

import (
	"fmt"
	"reflect"
	"time"
	"unsafe"
)

type ETS struct{} //ETS跟标准库的struct{}等价（可以互换）

// 所有的空结构体指向同一个地址(内核是完全一样的)
func allEmptyStructIsSame() {
	var a ETS
	var b ETS
	var c struct{}
	fmt.Printf("address of a %p b %p c %p\n", &a, &b, &c)
	fmt.Printf("size of a %d b %d c %d\n", unsafe.Sizeof(a), unsafe.Sizeof(b), unsafe.Sizeof(c))
	fmt.Printf("size of a %d b %d c %d\n", reflect.TypeOf(a).Size(), reflect.TypeOf(b).Size(), reflect.TypeOf(c).Size())
}

// 空结构体的应用场景
func scenariosOfEmptyStruct() {
	set := map[int]struct{}{
		1: {},
		4: struct{}{},
		7: struct{}{},
	}
	if _, exists := set[5]; exists {
		fmt.Println("5是存在的")
	} else {
		fmt.Println("5是不存在的")
	}

	// 通过这种方式可以实现主协程等待子协程的目的
	// 创建一个 channel，空的，容量为0，
	blocker := make(chan struct{})
	go func() {
		// 等待子协程2秒，然后打印done
		time.Sleep(2 * time.Second)
		fmt.Println("done")
		// 子协程完了之后，往channel写入这个空结构体，因为我们根本不关心是什么东西，使用空结构体还不占内存
		// 所以我们channel里面的数据类型就用struct{}这个空结构体
		blocker <- ETS{}
	}()
	<-blocker //主协程去读取channel，因为如果里面没有东西，读操作会阻塞，所以有东西的时候子协程run完了，主协程才会退出
	//通过这种方式就可以实现主协程等待子协程完成的目的
}

func main31() {
	allEmptyStructIsSame()
	scenariosOfEmptyStruct()
}
