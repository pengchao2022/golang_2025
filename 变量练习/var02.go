package main

import (
	"fmt"
	"reflect"
)

func main() {
	var a int8
	a = 127
	//简写 := 就可以
	b := 300
	c := 500.98

	//多个变量赋值 x,y,z
	var x, y, z = 109, 110, 120
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
	//逗号隔开打印多个变量的值
	fmt.Println(x, y, z)
	// 打印数据类型
	fmt.Println(reflect.TypeOf(x))
	fmt.Println(reflect.TypeOf(c))
}
