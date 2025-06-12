package main

import "fmt"

func main() {
	var num_list = []int{1, 2, 3, 4, 5}
	var my_nums = num_list
	fmt.Println(my_nums) //打印结果为 [1 2 3 4 5]
	num_list[1] = 4000
	fmt.Println(num_list) //打印结果为：[1 4000 3 4 5]
	fmt.Println(my_nums)  // 打印结果为：[1 4000 3 4 5]
	// 扩容前查看切片num_list 的内存地址
	fmt.Printf("%p\n", num_list) // 打印结果为 0xc000118030
	// 查看my_nums 的内存地址
	fmt.Printf("%p\n", my_nums) // 打印结果为 0xc000118030 可以看到扩容前两个切片的地址是相同的

	// 现在对 num_list 进行扩容
	num_list = append(num_list, 3000, 4000)
	fmt.Println(num_list) // 打印结果为 [1 4000 3 4 5 3000 4000]
	fmt.Println(my_nums)  // 打印结果为 [1 4000 3 4 5] 扩容后 num_list 和 my_num 将相互独立
	// 查看 扩容后的num_list 地址
	fmt.Printf("%p\n", num_list) // 打印结果为 0xc00001c0a0
	// 查看 my_num 的地址
	fmt.Printf("%p\n", my_nums) // 打印结果为 0xc000016240 可以看到扩容后两个切片的地址是不同的

	// 使用copy 函数
	var fruits = []string{"apple", "orange", "banana"}
	//var my_fruits []string 切片没有初始化 将无法copy
	var my_fruits = make([]string, 4, 8) //make 函数初始化
	copy(my_fruits, fruits)
	fmt.Println(my_fruits) // 打印结果为 [apple orange banana ]
	//打印内存地址
	fmt.Printf("fruits 的内存地址为：%p\n", fruits)      //0xc00007e0c0
	fmt.Printf("my_fruits的内存地址为：%p\n", my_fruits) //0xc00010e000 可以看到是完全不同的内存地址

	//修改fruits 的元素
	fruits[1] = "kiwi"
	fmt.Println(fruits)    //[apple kiwi banana]
	fmt.Println(my_fruits) // [apple orange banana ] 可以看到my_fruits 是不变的

	// copy 练习
	var num1 = []int{1, 2, 3, 4, 5, 6}
	var num2 = []int{2000, 3000, 4000}
	copy(num1, num2)
	fmt.Println(num1) // 打印结果为 [2000 3000 4000 4 5 6]
	copy(num2, num1)
	fmt.Println(num2)

}
