package main

import (
	"fmt"
	"os"
)

func main() {
	// 定义一个用户界面
	var panel = `
	******************客户信息管理系统*********************
	1. 查看客户
	2. 添加客户
	3. 删除客户
	4. 修改客户
	5. 退出程序
	*******************************************************
	`
	//定义一个切片
	var my_customers []string
	// 定义一个循环
	for {
		fmt.Println(panel)
		var choice int
		fmt.Println("请输入您的选择编号：")
		fmt.Scan(&choice)

		//编写switch 分支语句
		switch choice {
		case 1:
			fmt.Println("开始查看客户信息")
			fmt.Println("客户姓名：")
			for i, v := range my_customers {
				fmt.Println(i+1, v)
			}
			fmt.Println()

		case 2:
			fmt.Println("开始添加客户信息")
			var customer_name string
			fmt.Println("请输入客户姓名：")
			fmt.Scan(&customer_name)
			// 输入的用户名保存在切片中
			my_customers = append(my_customers, customer_name)
			fmt.Println("添加成功")
			fmt.Println()
			fmt.Println()

		case 3:
			fmt.Println("开始删除客户信息")
			var delName string
			fmt.Println("请输入要删除客户的姓名：")
			fmt.Scan(&delName)
			for i, v := range my_customers {
				if v == delName {
					my_customers = append(my_customers[:i], my_customers[i+1:]...)
				}
			}
			fmt.Println("删除成功")

		case 4:
			fmt.Println("开始修改客户信息")

		case 5:
			fmt.Println("程序即将退出,欢迎再次使用！")
			os.Exit(0)

		default:
			fmt.Println("输入有误，请重新输入")

		}

	}

}
