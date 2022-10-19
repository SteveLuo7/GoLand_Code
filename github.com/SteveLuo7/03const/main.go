package main

import "fmt"

const pi = 3.1415926

const (
	statusOK = 200
	notFound = 404
)

const (
	n1 = iota
	n2
	n3
)

const (
	b1 = iota
	b2
	_
	b3
)

const (
	c1 = iota
	c2 = 100
	c3 = iota
	c4
)

const (
	d1, d2 = iota + 1, iota + 2
	d3, d4 = iota + 1, iota + 2
)

//批量声明常量时 如果某一行声明后没有赋值 默认就和上一行一致

// 定义数量级
const (
	_  = iota
	KB = 1 << (10 * iota)
	MB = 1 << (10 * iota)
	GB = 1 << (10 * iota)
	TB = 1 << (10 * iota)
	PB = 1 << (10 * iota)
)

func main() {
	//pi =

	fmt.Println(n1, n2, n3)
	fmt.Println(b1, b2, b3)
	fmt.Println(c1, c2, c3, c4)
	fmt.Println(d1, d2, d3, d4)
	fmt.Println(KB, MB, GB, TB, PB)

}
