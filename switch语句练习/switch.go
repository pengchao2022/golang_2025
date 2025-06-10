// 练习书写switch 语句
package main

import "fmt"

func main() {
	var choice int
	info := `
	输入1：输出“普通攻击”
	输入2：输出“超级攻击”
	输入3：输出“使用道具”
	输入4：输出“逃跑”
	输入5：输出“法师装备”
	
	`
	fmt.Println(info)
	fmt.Println("请在此输入您的选择：")
	fmt.Scan(&choice)

	switch choice {
	case 1:
		fmt.Println("普通攻击")
	case 2:
		fmt.Println("超级攻击")
	case 3:
		fmt.Println("使用道具")
	case 4:
		fmt.Println("逃跑")
	case 5:
		fmt.Println("法师装备")
	default:
		fmt.Println("输入错误，请重新输入！")
	}
}
