package main

import "fmt"

// 函数在执行后，里面所有的变量都会被当做垃圾清理掉
func addNum(x, y int) int { //大括号前面的 int 为函数返回值类型
	ret := x + y
	return ret // 函数执行完后，所有变量都会被销毁，如果不想让变量被销毁，需要使用return 返回并保留该变量
}

// 函数有多个返回值
func query() (string, int) {
	return "yuan", 49
}

func main() {
	var ret1 = addNum(10, 12) //把返回值存入变量 ret1
	var ret2 = addNum(34, 90) //把返回值存入变量 ret2
	fmt.Println(ret1)
	fmt.Println(ret2)

	// 当然你也可以直接打印
	fmt.Println(addNum(4, 9))

	//有多个返回值的函数，需要有多个变量来接受
	ret5, ret6 := query()
	fmt.Println(ret5)
	fmt.Println(ret6)

}
