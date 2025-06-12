package main

import "fmt"

// 位置参数
func cal_add(x int, y int) { // 形参要是类型一样，也可以简写为：x,y int
	result := x + y
	println(result)

}

// 不定长参数
func cal_add_es(nums ...int) {
	ret := 0
	for _, num := range nums {
		ret += num

	}
	fmt.Println(ret)
}

// 写一个计算器函数
func calculater(operater string, nums ...int) {

	switch operater {

	// 做一个累加
	case "+":
		ret := 0
		for _, num := range nums {
			ret += num
		}
		fmt.Println(ret)

	// 做一个累乘
	case "*":
		ret := 1
		for _, num := range nums {
			ret *= num
		}
		fmt.Println(ret)
	}

}
func main() {
	cal_add(8, 7) // 位置参数 一一对应

	cal_add_es(900, 4, 5, 6, 7, 800)

	calculater("+", 23, 45, 67)
	calculater("*", 4, 4, 4)

}
