package main

import (
	"fmt"
)

func main() {
	//for i := 0; i < 10; i++ {
	//	fmt.Println(i)
	//}

	//var i = 5
	//for ; i < 10; i++ {
	//	fmt.Println(i)
	//}

	//var i = 5
	//for i < 10 {
	//	fmt.Println(i)
	//	i++
	//}

	//for {
	//	fmt.Println("123")
	//}
	//
	//s := "hello沙河"
	//for i, v := range s {
	//	fmt.Printf("%d %c\n", i, v)
	//}

	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d*%d=%d\t", j, i, j*i)

		}
		fmt.Println()
	}

}
