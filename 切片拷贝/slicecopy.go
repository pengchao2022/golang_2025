package main

import "fmt"

func main() {
	var num_list = []int{1, 2, 3, 4, 5}
	var my_nums = num_list
	fmt.Println(my_nums) //打印结果为 [1 2 3 4 5]
	num_list[1] = 4000
	fmt.Println(num_list) //打印结果为：[1 4000 3 4 5]
	fmt.Println(my_nums)  // 打印结果为：[1 4000 3 4 5]

	// 现在对 num_list 进行扩容
	num_list = append(num_list, 3000, 4000)
	fmt.Println(num_list) // 打印结果为 [1 4000 3 4 5 3000 4000]
	fmt.Println(my_nums)  // 打印结果为 [1 4000 3 4 5] 扩容后 num_list 和 my_num 将相互独立

}
