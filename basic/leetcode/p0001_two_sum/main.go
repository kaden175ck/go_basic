package main

import (
	"bufio"
	"fmt"
	"os"
)

// findTwoSum 在 numbers 中寻找两个数，使它们的和等于 target。
//
// 返回值说明：
//   - firstIndex：第一个数字的下标
//   - secondIndex：第二个数字的下标
//   - found：是否找到了符合条件的两个数字
func findTwoSum(numbers []int, target int) (firstIndex int, secondIndex int, found bool) {
	// numberToIndex 用来记录已经遍历过的数字及其下标。
	// map 的键是数组中的数字，值是这个数字对应的下标。
	numberToIndex := make(map[int]int)

	for currentIndex, currentNumber := range numbers {
		// 如果当前数字是 currentNumber，那么还需要 neededNumber，
		// 才能使两个数字的和等于 target。
		neededNumber := target - currentNumber

		// 查找 neededNumber 之前是否已经出现过。
		previousIndex, exists := numberToIndex[neededNumber]

		if exists {
			// 找到了符合条件的两个数字，返回它们的下标。
			return previousIndex, currentIndex, true
		}

		// 当前数字暂时没有找到配对，将它和它的下标保存到 map 中，
		// 供后面的数字查找。
		numberToIndex[currentNumber] = currentIndex
	}

	// 遍历完整个数组后仍未找到答案。
	return -1, -1, false
}

func main() {
	// 本程序约定的输入格式如下：
	// 第一行输入两个整数：数组长度 numberCount 和目标值 target。
	// 第二行输入 numberCount 个整数，表示数组中的元素。
	//
	// 示例输入：
	// 4 9
	// 2 7 11 15

	// 使用带缓冲的读取器读取标准输入。
	inputReader := bufio.NewReader(os.Stdin)

	// 使用带缓冲的写入器写入标准输出。
	outputWriter := bufio.NewWriter(os.Stdout)

	var numberCount int
	var target int

	// 读取数组长度和目标值。
	_, readError := fmt.Fscan(inputReader, &numberCount, &target)
	if readError != nil {
		// 输入为空或格式不正确时，直接结束程序。
		return
	}

	// 数组长度不能是负数。
	if numberCount < 0 {
		return
	}

	// 根据输入的数组长度创建切片。
	numbers := make([]int, numberCount)

	// 依次读取数组中的每一个数字。
	for index := 0; index < numberCount; index++ {
		_, readError = fmt.Fscan(inputReader, &numbers[index])
		if readError != nil {
			// 实际输入的数字数量不足时，直接结束程序。
			return
		}
	}

	// 调用算法函数寻找答案。
	firstIndex, secondIndex, found := findTwoSum(numbers, target)

	if found {
		// 找到答案时，输出两个数字的下标。
		fmt.Fprintln(outputWriter, firstIndex, secondIndex)
	} else {
		// 没有找到答案时，输出 -1 -1。
		fmt.Fprintln(outputWriter, -1, -1)
	}

	// 将缓冲区中的内容真正写入标准输出。
	outputWriter.Flush()
}
