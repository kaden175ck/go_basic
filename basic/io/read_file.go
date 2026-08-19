package io

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

// 读文件和用bufio读文件
func ReadFile() {
	if fin, err := os.Open("../data/verse.txt"); err != nil {
		fmt.Printf("open file faied: %v\n", err) //比如文件不存在
	} else {
		defer fin.Close()
		// 这样写说明长度和容量都是100
		bs := make([]byte, 100)
		fin.Read(bs)
		fmt.Println(string(bs))

		fin.Seek(0, 0)
		fin.Read(bs)
		fmt.Println(string(bs))

		// 再次回到文件起始位置再开始
		fin.Seek(0, 0)
		const BATCH = 10
		buffer := make([]byte, BATCH)
		for {
			n, err := fin.Read(buffer)
			if n > 0 {
				fmt.Println(buffer[0:n])
			}
			if err == io.EOF {
				break
			}
		}
	}
}

func ReadFileWithBuffer() {
	if fin, err := os.Open("../data/verse.txt"); err != nil {
		fmt.Printf("open file faied: %v\n", err) //比如文件不存在
	} else {
		defer fin.Close()
		reader := bufio.NewReader(fin)
		for {
			line, err := reader.ReadString('\n')
			if len(line) > 0 {
				line = strings.TrimRight(line, "\n")
				fmt.Println(line)
			}
			if err == io.EOF { // End Of File
				break
			}
		}
	}
}
