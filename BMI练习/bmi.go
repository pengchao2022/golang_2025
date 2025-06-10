// BMI 为体重指数 计算方式是 体重KG 除以身高的平方
package main

import (
	"fmt"
	"math"
)

func main() {
	var weight, height float64
	fmt.Println("请输入您的体重(kg):")
	fmt.Scan(&weight)
	fmt.Println("请输入您的身高(m):")
	fmt.Scan(&height)

	BMI := weight / math.Pow(height, 2)

	if BMI < 18.5 {
		fmt.Println("体重过轻")
	} else if BMI >= 18.5 && BMI < 25 {
		fmt.Println("正常")
	} else if BMI >= 25 && BMI < 30 {
		fmt.Println("偏胖")
	} else if BMI >= 30 && BMI < 35 {
		fmt.Println("肥胖")
	} else if BMI >= 35 {
		fmt.Println("重度肥胖")
	} else {
		fmt.Println("输入有误！请重新输入")
	}

}
