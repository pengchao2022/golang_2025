package main

import "fmt"

func main() {

	// 取余数运算符 %
	a := 7
	b := 2
	c := a % b
	fmt.Println(c)

	// 关系运算符 返回一个bool布尔值
	fmt.Println(a > b)
	fmt.Println(a != b)
	fmt.Println(a <= b)

	//逻辑运算符
	username := "allen"
	passwd := "12378965"

	fmt.Println(username == "allen" && passwd == "12378965")
	fmt.Println("可以登录")
	// 打印结果 true 可以登录
}
