package main

import "fmt"

func main() {
	var a = 5
	var b = 2

	// 算术运算符
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)
	fmt.Println(a % b)

	a++
	b--
	fmt.Println(a)
	fmt.Println(b)

	// 关系运算符
	fmt.Println(a == b) //go语言是强类型，相同类型的变量才可以进行比较
	fmt.Println(a != b)
	fmt.Println(a >= b)
	fmt.Println(a <= b)
	fmt.Println(a < b)
	fmt.Println(a > b)

	//逻辑运算符
	//如果年龄大于18岁 并且年龄小于60岁

	age := 22
	if age > 18 && age < 60 {
		fmt.Println("working hard")
	} else {
		fmt.Println("Happy everyday")
	}

	//如果年龄小于18岁 并且年龄大于60岁

	if age > 60 || age < 18 {
		fmt.Println("No working")

	} else {
		fmt.Println("Continue working")
	}

	//取反
	isMarried := false
	fmt.Println(isMarried)  //false
	fmt.Println(!isMarried) //true

	//位运算符 针对的是二进制
	//5 的二进制  101
	//2的二进制   10

	//&：按位与 （两位均为1才为1）
	fmt.Println(5 & 2) // 000

	// |: 按位或（两位有1个为1就为1）
	fmt.Println(5 | 2) // 111

	// ^: 按位异或 （两位不一样则为1)
	fmt.Println(5 ^ 2) // 111

	// <<: 将二进制位左移指定位数
	fmt.Println(5 << 1) // 10
	fmt.Println(1 << 10)
	fmt.Println(5 >> 1)

	var m = int8(1)
	fmt.Println(m << 10)

	// 赋值运算符：用来给变量赋值
	var x int
	x = 10
	x += 1
	fmt.Println(x)

}
