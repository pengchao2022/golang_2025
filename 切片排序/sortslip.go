package main

import (
	"fmt"
	"sort"
)

func main() {
	// 整型排序
	var age = []int{22, 45, 12, 98, 34, 56, 8}
	sort.Ints(age)
	fmt.Println(age) // 打印结果为 [8 12 22 34 45 56 98]

	// 字符串型排序
	var fruits = []string{"orange", "apple", "kiwi", "cherry", "banana", "tomato"}
	sort.Strings(fruits)
	fmt.Println(fruits) // 打印结果为 [apple banana cherry kiwi orange tomato]

	// 浮点型排序
	var heights = []float64{8.99, 4.56, 1.34, 4.89, 5.34, 1.78}
	sort.Float64s(heights)
	fmt.Println(heights) //打印结果为 [1.34 1.78 4.56 4.89 5.34 8.99]

}
