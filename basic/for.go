package main

import "fmt"

func main9() {
	var sum int
	a, b := 6, 30
	for a >= 0 && b > a {
		sum += a
		if a == 4 {
			break
		}
		a, b = a-1, b/2

		for i := 0; i < 10; i++ {
			for j := 8; j > 0; j -= 2 {
				sum += i * j
			}
		}
	}
	fmt.Println(sum, a, b)

	for {
		fmt.Println("hello")
		a--
		if a <= 3 {
			break
		}
	}
}
