package main

import "fmt"

func niming() (int, int) {
	return 100, 200
}
func main() {

	// 匿名变量 用下划线表示 _ , 当只使用函数中一个返回值时，不使用的返回值 用下划线表示 _
	var x, _ = niming()
	fmt.Println(x)

}
