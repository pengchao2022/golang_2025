package main

import "fmt"

func main() {
	//基本数据类型都是 值 类型
	//值类型在声明未赋值之前都有一个默认值
	var x int
	fmt.Println("int default", x) // 打印结果为 0 整型默认值为 0
	x = 10

	//浮点型 默认值也为 0
	var y float64
	fmt.Println("float default", y) //打印结果为 0

	//字符串 string 默认值 为空
	var str1 string
	fmt.Println("string default", str1) //打印结果为空

	//布尔型 bool 默认值为 false
	var my_bo bool
	fmt.Println("bool default", my_bo) //打印结果为 false

	//指针变量 默认返回值 为 <nil>  空对象
	var p *int
	fmt.Println(p) //打印出空 nil
	//使用new()函数开辟存储空间
	p = new(int)
	//然后再给p 赋值
	*p = 100
	fmt.Println(*p) //打印出 100

}
