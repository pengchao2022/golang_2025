// 数组的练习
package main

import "fmt"

func main() {
	// 案例1 先声明再赋值
	var fruits [5]string
	fruits[0] = "apple"
	fruits[1] = "orange"
	fruits[2] = "watermalen"
	fruits[3] = "kiwi"
	fruits[4] = "banana"

	fmt.Println(fruits)

	//案例2 声明的同时可以直接赋值
	var animals = [5]string{"dog", "chicken", "duck", "pig", "cat"}
	fmt.Println(animals)

	var names = [6]string{"kate", "jim", "lily", "sophia", "allen", "tom"}
	fmt.Println(names)

	var height = [5]float64{1.87, 1.34, 1.67, 1.80, 1.89}
	fmt.Println(height)

	var age = [6]int{24, 34, 56, 78, 38, 90}
	fmt.Println(age)

	// 省略数组长度
	cities := []string{"sahnghai", "jinan", "beijing", "jilin", "dalian"}
	fmt.Println(cities)

	months := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	fmt.Println(months)

	// 数组的空间地址都是连续的
	// 数组是支持索引和切片操作的
	// 数组的切片arr[开始索引:结束索引]
	var name_list = [6]string{"kate", "jim", "lily", "sophia", "allen", "tom"}
	//拿出allen
	fmt.Println(name_list[4])
	// 数组切片
	fmt.Println(name_list[1:3])
	// 检查数组里面元素之间内存是否连续
	fmt.Println(&name_list[0]) //一个字符串占 16个字节 ，所以地址之间相隔 16个字节
	fmt.Println(&name_list[1])
	fmt.Println(&name_list[2])
	fmt.Println(&name_list[3])
	fmt.Println(&name_list[4])
	fmt.Println(&name_list[5])
	fmt.Println()

	// 用int8 来举例，int8 是占一个字节，int16 占2个字节， int32 占4个字节， int64 占8个字节
	var age_list = [4]int8{34, 36, 37, 40}
	fmt.Println(&age_list[0])
	fmt.Println(&age_list[1])
	fmt.Println(&age_list[2])
	fmt.Println(&age_list[3])
	/*
		0xc000116080
		0xc000116081
		0xc000116082
		0xc000116083
		打印结果是连续的
	*/

}
