// 输入两个数字，打印出较小的那个数字
package main

import "fmt"

func main() {
	var x, y int
	fmt.Println("请在此输入第一个数：")
	fmt.Scan(&x)
	fmt.Println("请输入第二个数：")
	fmt.Scan(&y)
	if x > y {
		x, y = y, x
	}
	fmt.Printf("较小的数为：%d\n", x)
}
