package main

import (
	"fmt"
	"reflect"
)

func main() {
	var s = 'a'
	fmt.Println(reflect.TypeOf(s)) // 打印结果为int32 即int 类型 存一个字符，会是一个数字
	fmt.Println(s)                 // 打印结果为 97
	// 可以转码 转换为字符类型 string
	fmt.Println(string(s)) // 打印结果为 a

	// byte 等于 unit8 , rune 等于 int32

	var s1 byte // 声明s1 为字节型
	s1 = 'a'
	fmt.Println(s1) // 打印出来是数字

	var s2 uint8 // 声明是 数字型
	s2 = 97
	fmt.Println(s2) // 打印出来是数字
	fmt.Println(string(s2))

	var s3 rune // 存储中文用rune
	s3 = '马'
	fmt.Println(s3)         // 打印出的也是数字 39532
	fmt.Println(string(s3)) // utf-8 int32 使用 4 个字节 编码汉字

}
