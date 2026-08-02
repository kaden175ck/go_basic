package main

import "fmt"

func basic_goto() {
	var i int = 4
MYLABEL:
	i += 3
	i *= 2
	fmt.Println(i)
	if i > 200 {
		return
	}
	goto MYLABEL

}

func if_goto() {
	var i int = 4
	if i%2 == 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	i += 3
	fmt.Println(i)
L2:
	i *= 3
	fmt.Println(i)
}

func for_goto() {
	const SIZE = 5
L1:
	for i := 0; i < SIZE; i++ {
	L2:
		fmt.Printf("开始检查第%d行\n", i)
	L3:
		if i%2 == 1 {
			for j := 0; j < SIZE; j++ {
				fmt.Printf("开始检查第%d列\n", j)
				if j%3 == 0 {
					goto L1 //i从0开始，运行一个全新的for循环。把goto换成break或continue不是开启一个新的for循环
				} else if j%3 == 1 {
					goto L2
				} else {
					goto L3
				}
			}
		}
	}

}

func continue_label() {
	const SIZE = 5
L1:
	for i := 0; i < SIZE; i++ {
	L2:
		fmt.Printf("开始检查第%d行\n", i)
	L3:
		if i%2 == 1 {
			for j := 0; j < SIZE; j++ {
				fmt.Printf("开始检查第%d列\n", j)
				if j%3 == 0 {
					continue L1 //此时i本来是1，回到L1就继续执行i=2，不从0重新开始
				} else if j%3 == 1 {
					goto L2
				} else {
					goto L3
				}
			}
		}
	}
}

func break_label() {
	const SIZE = 5
L1:
	for i := 0; i < SIZE; i++ {
	L2:
		fmt.Printf("开始检查第%d行\n", i)

		if i%2 == 1 {
		L3:
			for j := 0; j < SIZE; j++ {
				fmt.Printf("开始检查第%d列\n", j)
				if j%3 == 0 {
					break L1 //直接退出最外层的for循环就是L1的循环，整个函数结束
					//continue和break针对的Label, 这个label必须是个循环loop
					// 就是说如果你改成break L2这种print语句，break，continue就会报错
				} else if j%3 == 1 {
					goto L2
					//goto可以针对任意Label
				} else {
					break L3
				}
			}
		}
	}
}

func main10() {
	// basic_goto()
	// if_goto()
	// for_goto()
	// continue_label()
	break_label()
}
