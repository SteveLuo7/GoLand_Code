package main

import "fmt"

func main() {
	// 当i=5时候 跳出for循环
	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}
		fmt.Println(i)

	}
	fmt.Println("over")
	// 当i = 5时候，跳过此次for循环

	for i := 0; i < 10; i++ {
		if i == 5 {
			continue
		}
		fmt.Println(i)

	}
	fmt.Println("over")
}
