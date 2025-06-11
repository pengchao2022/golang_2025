// 指针的练习
package main

import "fmt"

func main() {
	var x int
	x = 10
	var y = &x
	var z = *y
	x = 20
	fmt.Println(x)
	fmt.Println(*y)
	fmt.Println(z) //需要注意的是 z 的值为 10 而不是  20

	//练习2
	var a = 100
	var b = &a //var b *int
	var c = &b //var c * *int --- c的类型
	**c = 200
	fmt.Println(a)      //打印结果为 200
	fmt.Printf("%T", c) //打印c 的类型可以看到是 **int
}
