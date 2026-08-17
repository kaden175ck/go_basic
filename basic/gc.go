package main

import (
	"fmt"
	"os"
	"runtime/debug"
	"runtime/trace"
	"sync"
	"time"
)

// 值类型（包含数组）的全局变量分配在全局数据段上（生命周期一直到程序结束）
var arr [1000]int

// 引用类型的全局变量内存分配在堆上
var slc []int = make([]int, 1000)

func stack_heap() {
	var brr [1000]byte //函数的入参、出参、局部变量一般在栈上，体量很大时分配在堆上

	var crr [128 << 10]byte
	var drr [128<<10 + 1]byte //数组超过128K就会分配到堆上(moved to heap)

	err := make([]byte, 64<<10)   //函数的入参、出参、局部变量一般在栈上
	frr := make([]byte, 64<<10+1) //切片超过64K就会分配到堆上(escapes to heap)

	_ = arr
	_ = brr
	_ = crr
	_ = drr
	_ = err
	_ = frr
	_ = slc
}

const (
	NumWorkers    = 4     // Number of workers.
	NumTasks      = 500   // Number of tasks.
	MemoryIntense = 10000 // Size of memory-intensive task (number of elements).
)

// 所有的垃圾回收都是针对堆的
func gc() {
	// Write to the trace file.
	// 创建 data 文件夹。
	// 如果文件夹已经存在，MkdirAll 也不会报错。
	createDirectoryError := os.MkdirAll("data", 0755)
	if createDirectoryError != nil {
		fmt.Println("创建 data 文件夹失败：", createDirectoryError)
		return
	}

	// 在 data 文件夹中创建 trace.out 文件。
	traceFile, createFileError := os.Create("data/trace.out")
	if createFileError != nil {
		fmt.Println("创建 trace 文件失败：", createFileError)
		return
	}

	// 函数结束前关闭文件。
	defer traceFile.Close()

	// 开始记录程序运行信息。
	startTraceError := trace.Start(traceFile)
	if startTraceError != nil {
		fmt.Println("启动 trace 失败：", startTraceError)
		return
	}

	// gc 函数结束前停止记录。
	defer trace.Stop()

	// Set the target percentage for the garbage collector. Default is 100%.
	debug.SetGCPercent(100) //触顶时再执行GC

	// Task queue and result queue.
	taskQueue := make(chan int, NumTasks)
	resultQueue := make(chan int, NumTasks)

	// Start workers.
	var wg sync.WaitGroup
	wg.Add(NumWorkers)
	for i := 0; i < NumWorkers; i++ {
		go worker(taskQueue, resultQueue, &wg)
	}

	// Send tasks to the queue.
	for i := 0; i < NumTasks; i++ {
		taskQueue <- i
	}
	close(taskQueue)

	// Retrieve results from the queue.
	go func() {
		wg.Wait()
		close(resultQueue)
	}()

	// Process the results.
	for result := range resultQueue {
		fmt.Println("Result:", result)
	}

	fmt.Println("Done!")
}

// Worker function.
func worker(tasks <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	for task := range tasks {
		result := performMemoryIntensiveTask(task)
		results <- result
	}
}

// performMemoryIntensiveTask is a memory-intensive function.
func performMemoryIntensiveTask(task int) int {
	// Create a large-sized slice.
	data := make([]int, MemoryIntense) //申请一块大内存，虽是局部变量，但会逃逸到堆上，长时间累积会触发GC
	for i := 0; i < MemoryIntense; i++ {
		data[i] = i + task
	}

	// Latency imitation.
	time.Sleep(10 * time.Millisecond)

	// Calculate the result.
	result := 0
	for _, value := range data {
		result += value
	}
	return result
}

func main41() {
	stack_heap()
	gc() // 程序运行完之后生成一个文件data/trace.out, 然后执行 go tool trace data/trace.out
}

// go run -gcflags=-m ./basic/gc.go
// go run -gcflags=-m ./gc.go
// go run ./basic/gc.go
// go tool trace data/trace.out
// 可以看到打开浏览器可以看到运行过程触发了很多次gc，不好，需要avoid
