package main

import "fmt"

func main() {
	var a, b int
	fmt.Println("请在此输入第一个数字：")
	fmt.Scan(&a)
	fmt.Println("请在此输入第二个数字：")
	fmt.Scan(&b)
	// 打印出a,b 的值
	fmt.Printf("a=%d\n", a)
	fmt.Printf("b=%d\n", b)

}
