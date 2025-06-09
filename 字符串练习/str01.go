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
}
