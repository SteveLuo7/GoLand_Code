package main

import "fmt"

//数组

//存放元素的容器
//必须指定存放的元素的类型和容量
//数组的长度是数组类型的一部分

// var 数组变量名 数组长度 数组类型

func main() {
	var a1 [3]bool
	var a2 [4]bool

	fmt.Printf("a1:%T a2:%T\n", a1, a2)

	//数组的初始化
	//如果不初始化：默认元素都是零值 （布尔值 false 整数和浮点型都是0 字符串”“）
	fmt.Println(a1, a2)
	a1 = [3]bool{true, true, true}
	fmt.Println(a1)

	//a100 := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	a10 := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9} //根据数组的长度自定计算数组的长度
	fmt.Println(a10)

	a3 := [5]int{0: 1, 4: 2} // 根据索引进行初始化
	fmt.Println(a3)

	//数组的遍历
	citys := [...]string{"beijing", "shanghai", "shenzhen"}
	for i := 0; i < len(citys); i++ {
		fmt.Println(citys[i])
	}

	//for range 数组遍历
	for i, v := range citys {
		fmt.Println(i, v)
	}

	//多堆数组
	//[[1,2] [3,4] [5,6]]
	var a11 [3][2]int
	a11 = [3][2]int{
		[2]int{1, 2},
		[2]int{3, 4},
		[2]int{5, 6},
	}
	fmt.Println(a11)

	//多维数组的遍历
	for _, v1 := range a11 {
		fmt.Println(v1)
		for _, v2 := range v1 {
			fmt.Println(v2)
		}
	}

	// 数组是值类型
	b1 := [3]int{1, 2, 3}
	b2 := b1
	b2[0] = 100
	fmt.Println(b1, b2)
}
