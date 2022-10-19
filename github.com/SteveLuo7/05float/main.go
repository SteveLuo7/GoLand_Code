package main

import "fmt"

func main() {
	//math.MaxFloat32
	f1 := 1.23456
	fmt.Printf("%T\n", f1) // 默认go语言中 小数都是float64类型的
	f2 := float32(1.23456)
	fmt.Printf("%T\n", f2)

	//f1 = f2  类型不一样 是不能进行强制转换
}
