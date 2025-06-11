package main

import "fmt"

func main() {
	var a = make([]int, 3, 10)
	fmt.Println(a) // 打印出 [0 0 0]
	b := append(a, 34, 55, 66)
	fmt.Println(b) // 打印出 [0 0 0 34 55 66] 记住，虽然a 的前三个元素为0 ，但使用append 追加的时候还是从第四个元素开始
	fmt.Println(a) // 打印结果为 [0 0 0] 切片a 没有改变
	a[0] = 100
	fmt.Println(a)
	fmt.Println(b) // 没有扩容 a, b 的前三个元素是共享的，即 改变a 的前三个元素，b 的前三个元素也会改变

}
