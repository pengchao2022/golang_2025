package main

import "fmt"

func main() {
	//切片的构建
	//1 通过数组切片生成
	var arr = [...]int{11, 33, 45, 8, 97, 66, 108, 209, 31}
	fmt.Printf("%T\n", arr) // 数组的类型 [9]int
	//对数组切片后得到切片对象
	s1 := arr[0:5]
	s2 := arr[4:]
	s3 := arr[:7]
	fmt.Println(s1, s2, s3)
	//查看切片类型
	fmt.Printf("%T,%T,%T\n", s1, s2, s3) // 打印结果 []int,[]int,[]int
	// 修改数组里面元素的值
	arr[2] = 500
	fmt.Println(s1) //修改数组里面元素后，切片也会跟着改变 [11 33 500 8 97]
	fmt.Println(cap(s1))
	fmt.Println(arr)

	/*
		切片原理
		切片存储三个重要数值：
		1，起始地址
		2，切的长度
		3，容量--- 原数组的长度
		注意：切片本身并不存储任何数组元素，存储的只是对原数组的引用信息
	*/

	// 切片容量cap 为数组的总长度减去起始的位置
	// 切片的长度为 末尾减去起始

	// 容量应为9-4 = 5
	// 代码打印容量
	fmt.Println("打印s2容量")
	fmt.Println(cap(s2))

	//练习计算切片容量和长度
	var my_arr = [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	a1 := my_arr[0:3] // 计算长度为3-0=3 , 容量为9-0=9
	a2 := my_arr[0:5] // 计算长度为5-0=5, 容量为9-0=9
	a3 := my_arr[1:5] // 长度为 5-1=4， 容量为9-1=8
	a4 := my_arr[1:]  // 长度为 9-1=8， 容量为9-1=8
	a5 := my_arr[:]   // 长度为9-0=9， 容量为9-0=9
	a6 := my_arr[1:2]

	fmt.Printf("a1的长度为:%d,容量为:%d\n", len(a1), cap(a1))
	fmt.Printf("a2的长度为:%d,容量为：%d\n", len(a2), cap(a2))
	fmt.Printf("a3的长度为:%d,容量为：%d\n", len(a3), cap(a3))
	fmt.Printf("a4的长度为：%d,容量为：%d\n", len((a4)), cap(a4))
	fmt.Printf("a5的长度为：%d,容量为：%d\n", len(a5), cap(a5))
	fmt.Printf("a6的长度为:%d,容量为：%d\n", len(a6), cap(a6))

}
