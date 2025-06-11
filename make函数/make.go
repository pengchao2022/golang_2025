package main

import "fmt"

func main() {
	var x [3]int
	fmt.Println(x)
	var y []int
	fmt.Println(y) // 无法初始化
	//y[0] = 100
	//fmt.Println(y) 切片没有长度和容量无法插入数值，即无法初始化
	//所以这种情况就引入了make 函数

	var z = make([]int, 4, 6) //初始化切片 长度为4，容量为6
	fmt.Println(z)            //初始化为 [0 0 0 0]
	z[0] = 100
	fmt.Println(z) //打印结果为 [100 0 0 0]

}
