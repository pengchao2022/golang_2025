package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	// 整型数字字符串转换成数字
	var a = "1"
	var b = "2"
	c := a + b
	fmt.Println(c) // 结果为字符串拼接 12

	// 当字符串转换为数字时，会返回2个值 分别为value 和 error 错误信息
	value_a, err_a := strconv.Atoi(a)
	// 下来需要进行判断，如果 error 不为空，则打印出错误信息，如果 error 为空，则只打印转换后的value 值, nil 表示空
	if err_a != nil {
		fmt.Println(err_a)
	}
	num_a := value_a
	fmt.Println(num_a)

	value_b, err_b := strconv.Atoi(b)
	if err_b != nil {
		fmt.Println(err_b)
	}
	num_b := value_b
	fmt.Println(num_b)

	num_c := num_a + num_b
	fmt.Printf("a + b 的值为：%d\n", num_c)

	// 不使用if 语句，使用匿名变量存储 error 信息
	myvalue_a, _ := strconv.Atoi(a)
	mynum_a := myvalue_a
	fmt.Println(mynum_a)
	myvalue_b, _ := strconv.Atoi(b)
	mynum_b := myvalue_b
	fmt.Println(mynum_b)

	mynum_c := mynum_a + mynum_b
	fmt.Println(mynum_c, reflect.TypeOf(mynum_c))

	// 将数字转换成字符串
	// 数字转换成字符串只有一个返回值string
	age := 23
	my_age := strconv.Itoa(age)
	fmt.Println(my_age, reflect.TypeOf(my_age)) // 打印结果为： 23 string

	// 浮点型字符串转换成数字

	my_circle := "3.1415926789"
	my_circle_num, _ := strconv.ParseFloat(my_circle, 64)
	fmt.Println(my_circle_num, reflect.TypeOf(my_circle_num))

}
