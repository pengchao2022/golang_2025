package main

import "fmt"

func main() {
	fmt.Println("程序启动")
	age := 12
	if age >= 18 {
		fmt.Println("可以购买万宝路香烟")
	} else {
		fmt.Println("不可以购买香烟")
	}
	fmt.Println("程序关闭")

}
