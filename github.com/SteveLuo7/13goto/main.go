package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 10; i++ {
		for j := 'A'; j < 'Z'; j++ {
			if j == 'C' {
				goto xx //跳出到指定的标签
			}
			fmt.Printf("%v-%c\n", i, j)
		}
	}
xx: // 标签
	fmt.Println("over")
}
