package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "Hello World 罗仕斌"

	n := len(s)
	fmt.Println(n)

	//for i := 0; i < len(s); i++ {
	//	fmt.Println(s[i])
	//
	//	fmt.Printf("%c\n", s[i])
	//
	//}

	s2 := "白萝卜"
	s3 := []rune(s2)
	s3[0] = '红'
	fmt.Println(string(s3)) // 把rune切片强制转换成字符串

	c1 := "红"
	c2 := '红'
	fmt.Printf("c1:%T\n c2:%T\n", c1, c2)

	c3 := "H"
	c4 := byte('H')
	fmt.Printf("c3:%T\n c4:%T\n", c3, c4)

	//双引号包裹的是字符串 string类型 '' 单引号是字符

	//类型转换

	a := 10
	var f float64
	f = float64(a)
	fmt.Println(f)
	fmt.Printf("%T\n", f)

	q := "hello 沙河小王子"
	q1 := strings.Split(q, "%c\n")
	fmt.Println(q1)

}
