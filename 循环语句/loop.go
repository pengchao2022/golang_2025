// 循环语句练习
package main

import "fmt"

func main() {
	// for 循环语句
	// 无限循环
	/*for true {
		fmt.Println("Hello kitty")
	}
	*/

	count := 0
	for count < 10 {
		fmt.Println("循环练习")
		count++
	}
	fmt.Println(count)

	// 打印从1 到 100
	num := 1
	for num <= 100 {
		fmt.Println(num)
		num++
	}
	// 打印从100 到 1
	num1 := 100
	for num1 >= 1 {
		fmt.Println(num1)
		num1--
	}

	// 循环主体 for 循环三要素
	for count := 0; count < 100; count++ {
		fmt.Println("循环主体打印")
		fmt.Println(count)

	}
	// 计算1加到100 的和
	sum := 0
	for i := 0; i <= 100; i++ {
		sum += i
	}
	fmt.Println(sum)
}
