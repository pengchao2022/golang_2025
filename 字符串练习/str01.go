package main

import "fmt"

func main() {
	// go 语言对于字符串只能用双引号
	str1 := "welcome to the world of go language!"
	fmt.Println(str1)
	/*
	   字符串索引
	*/
	fmt.Println(string(str1[0:6]))
	//前面缺省表示从0开始
	fmt.Println((string(str1[:7])))
	//后面缺省表示从末尾结束
	fmt.Println(string(str1[8:]))
	/*
	   字符串拼接
	*/
	name1 := "allen"
	name2 := " ma"
	name := name1 + name2
	fmt.Println(name)

	/*
		转义符练习
	*/

	name_list := "allen\nkate\nlily\n"
	fmt.Println(name_list)

	/*
		双引号的打印
	*/
	qu_str := "Hello, this is \"jack\""
	fmt.Println(qu_str)

	/*
		文件路径的打印
	*/
	my_path := "\\d\\work\\next"
	fmt.Println(my_path)

	/*
		多行内容的打印 使用反引号
	*/
	my_words := `
	春眠不觉晓
	   处处闻啼鸟
	      夜来风雨声
		     花落知多少
	
	`
	fmt.Println(my_words)

	customer_choose := `
	1,添加客户信息
	2,删除客户信息
	3,查询客户信息
	4,修改客户信息
	5,退出系统
	
	`
	fmt.Println(customer_choose)

}
