package main

import (
	"fmt"
	"strings"
)

func main() {

	path := "\"D:\\Go_Code\""
	fmt.Println(path)

	s := "'i am ok'"
	fmt.Println(s)

	s2 := `
		世情薄
		人情恶
		雨送黄昏花易落
`

	fmt.Println(s2)

	s3 := `D:\Go_Code`
	fmt.Println(s3)

	fmt.Println(len(s3))

	name := "罗仕斌"
	world := "dsb"

	ss := name + world
	fmt.Println(ss)
	ss1 := fmt.Sprintf("%s%s", name, world)
	fmt.Printf("%s%s", name, world)

	fmt.Println(ss1)

	ret := strings.Split(s3, "\\")
	fmt.Println(ret)
}
