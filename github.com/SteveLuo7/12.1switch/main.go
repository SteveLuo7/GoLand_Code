package main

import "fmt"

// 12.1switch
// 简化大量的判断

func main() {
	var n = 1
	//if n == 1 {
	//	fmt.Println("big thumb")
	//} else if n == 2 {
	//	fmt.Println("mid thumb")
	//} else if n == 3 {
	//	fmt.Println("little thumb")
	//} else {
	//	fmt.Println("none")
	//}

	switch n {
	case 1:
		fmt.Println("big thumb")
	case 2:
		fmt.Println("mid thumb")
	case 3:
		fmt.Println("little thumb")
	case 4:
		fmt.Println("none")
	}
}
