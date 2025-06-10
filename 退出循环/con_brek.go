// 练习退出循环
// break 退出整个循环体
package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
		if i == 7 {
			break // 7 以后的数字将不再打印，循环终止
		}
	}

	for count := 100; count > 0; count-- {
		if count == 77 {
			continue // 77 这儿会跳过去，不会打印77
		}
		fmt.Println(count)
	}

	// 计算1到100 除了能整除13 以外的和
	sum := 0
	for k := 1; k <= 100; k++ {
		if k%13 == 0 {
			continue
		}
		sum = sum + k

	}

	fmt.Printf("越过能整除13以外的和为：%d", sum)
}
