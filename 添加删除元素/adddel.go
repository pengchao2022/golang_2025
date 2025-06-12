package main

import "fmt"

func main() {
	// 在切片头部插入元素
	var s = []int{1, 2, 3, 4, 5, 6}
	new_s := append([]int{100}, s...)
	fmt.Println(new_s) // 打印结果为 [100 1 2 3 4 5 6]
	fmt.Println(s)     // 打印结果为 [1 2 3 4 5 6]

	// 首位置插入一个切片
	new_s_1 := append([]int{101, 102, 103, 104}, s...)
	fmt.Println(new_s_1) // 打印结果为 [101 102 103 104 1 2 3 4 5 6]

	// 删除某索引值
	var s1 = []int{100, 200, 300, 400, 500}
	// 删除索引为 2 的值
	new_s1 := append(s1[:2], s1[3:]...) // 索引为2 之前的取出来，索引为3 之后的追加
	fmt.Println(new_s1)

	// 删除索引为 3 的值
	new_s1_1 := append(s1[:3], s1[4:]...) // 索引为3 之前的取出来，索引为4 之后的追加
	fmt.Println(new_s1_1)

	// 删除索引为 0 的值
	new_s1_2 := s1[1:]
	fmt.Println(new_s1_2)

	// 删除索引为 4 的值
	new_s1_3 := s1[:4]
	fmt.Println(new_s1_3)

	//按照元素的值来删除
	var s3 = []int{1000, 2000, 3000, 4000}
	var delete_value = 2000
	// 新建一个空切片用于存储s3的元素
	var s4 []int
	for _, v := range s3 {
		if v != delete_value {
			s4 = append(s4, v)
		}
	}
	fmt.Println("删除后的切片为：")
	fmt.Println(s4)

	// 双指针原地删除
	i := 0
	for _, v := range s3 {
		if v != 2000 {
			s3[i] = v
			i++

		}

	}
	s3 = s3[:i]
	fmt.Println("删除后的切片为：")
	fmt.Println(s3)
}
