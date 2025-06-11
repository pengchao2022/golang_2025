package main

import "fmt"

func main() {
	//输出函数练习
	//打印多个
	fmt.Println("kate,lily,tim")
	fmt.Println("kate", "lily", "tim")

	// ln换行，不加ln 练习
	fmt.Print("kate", "lily", "tim") // 打印结果会连在一起 katelilytim

	// printf 格式化输出
	// 字符串拼接
	name := "allen"
	age := "24"
	age1 := 24
	height := 1.82
	fmt.Printf("姓名：" + name)
	fmt.Printf("年龄：" + age)

	//格式化输出
	fmt.Printf("姓名：%s, 年龄：%d, 身高: %.2f(m)", name, age1, height)
	fmt.Println("")

	//强制占位
	fmt.Printf("演示强制占位")
	fmt.Println()
	fmt.Printf("姓名：%v, 年龄：%v, 身高：%v(m)", name, age, height)

	// %T 显示变量的类型
	fmt.Println()
	fmt.Printf("显示身高变量的类型：%T", height) //输出结果 显示身高变量的类型：float64

	// 打印地址 %p
	// 需要用到取指符号 &
	fmt.Println()
	fmt.Printf("打印年龄age 变量的内存地址：%p", &age)
	fmt.Println()

	//Sprint 将输出结果保存到一个字符串里
	str := fmt.Sprintf("姓名：%v, 年龄：%v, 身高：%v(m)", name, age, height)
	fmt.Println(str)

	//***************************************************************************************
	// 输入函数 fmt.Scan()
	var x, y int
	fmt.Println("请输入两个数字：") // 注意输入多个数字的时候，默认以空格作为分割符
	fmt.Scan(&x, &y)
	fmt.Println(x, y)

	// scanln 遇到回车立马结束输入
	// scanf 格式化输入
	fmt.Println("下面演示scanf 的格式化输入举例")
	var num1, num2 int
	fmt.Println("请输入两数相加：")
	fmt.Scanf("%d+%d", &num1, &num2)
	my_num := num1 + num2
	fmt.Printf("结果为：%d\n", my_num)
}
