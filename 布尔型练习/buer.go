package main

import "fmt"

func main() {

	a := 10
	b := 20

	// 比较运算
	fmt.Println("a == b:", a == b) // false
	fmt.Println("a != b:", a != b) // true
	fmt.Println("a < b:", a < b)   // true
	fmt.Println("a > b:", a > b)   // false
	fmt.Println("a <= b:", a <= b) // true
	fmt.Println("a >= b:", a >= b) // false

	// 字符串比较
	s1 := "hello"
	s2 := "world"
	fmt.Println("s1 == s2:", s1 == s2) // false

}
