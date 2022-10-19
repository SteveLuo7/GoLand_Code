package main

import "fmt"

//  Go语言中推荐使用驼峰式命名

// var studentName string

// var name string
// var age int
// var isOk bool

var (
	name string
	age  int
	isOk bool
)

func main() {
	name = "Luo"
	age = 16
	isOk = true
	// Go语言变量声明必须使用，不适用就编译不过去
	fmt.Printf("name:%s\n", name)
	fmt.Println(age)
	fmt.Print(isOk)
	fmt.Println()

	//声明变量同时赋值
	var s1 string = "luoshibin"
	fmt.Println(s1)

	//类型推导 （根据值判断该变量是什么类型）
	var s2 = "20"
	fmt.Println(s2)

	//简短变量声明
	s3 := "罗仕斌"
	fmt.Println(s3)
}
