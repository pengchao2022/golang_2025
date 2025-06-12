package main

import "fmt"

// end 为形参，并需要指定形参类型为 int
func print_num(end int) {
	for i := 1; i <= end; i++ {
		fmt.Println(i)
	}
}

func main() {
	// 10 为实参 打印1-10  实参可以认为是变量赋值，将值赋给变量end
	print_num(10)
	// 100 为实参 打印1-100
	print_num(100)

	//1000 为实参 打印1-1000
	print_num(1000)

}
