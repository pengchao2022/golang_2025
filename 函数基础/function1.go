package main

import "fmt"

// 自定义自己的一个函数
// 函数声明不会执行内部的代码
// 函数作用是包裹 或者 封装
// 函数是代码的一种组织形式
func loop_print() {
	for i := 1; i <= 100; i++ {
		fmt.Println(i)
	}
}

func main() {
	//自定义的函数，需要放在main()函数里去执行
	// main 函数是程序的入口 对于go 语言必须要有 main 包 ，main函数
	loop_print()

}
