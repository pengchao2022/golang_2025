package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	// 输入多个客户信息并存储字切片中
	control_panel := `

	********************** 客户信息管理系统 *******************
	1.查看用户信息                     2.添加用户信息
	3.删除用户信息                     4.修改用户信息
	5.退出系统

	*********************************************************
	`
	var my_customer_slice = []map[string]string{}
	count := 0
	for {
		fmt.Println(control_panel)
		var choice int
		fmt.Println("请在此输入你的选择(1-5):")
		fmt.Scan(&choice)
		switch choice {
		case 1:
			fmt.Println("查询用户信息")
			for i, value := range my_customer_slice {
				fmt.Printf("客户ID：%d    ,客户姓名：%s   ,客户年龄：%s   ,客户性别：%s   ,客户身高：%s   ,客户联系方式：%s   \n", i+1, value["name"], value["age"], value["gender"], value["height"], value["contact"])
			}

		case 2:
			fmt.Println("添加用户信息")
			var customer_name string
			fmt.Println("请输入客户姓名：")
			fmt.Scan(&customer_name)
			var customer_age string
			fmt.Println("请输入客户年龄：")
			fmt.Scan(&customer_age)
			var customer_gender string
			fmt.Println("请输入客户性别：")
			fmt.Scan(&customer_gender)
			var customer_height string
			fmt.Println("请输入客户身高：")
			fmt.Scan(&customer_height)
			var customer_mobile string
			fmt.Println("请输入客户联系方式：")
			fmt.Scan(&customer_mobile)
			var customer_add = map[string]string{"name": customer_name, "age": customer_age, "gender": customer_gender, "height": customer_height, "contact": customer_mobile}
			my_customer_slice = append(my_customer_slice, customer_add)
			count++
			fmt.Printf("添加成功！共添加%d条记录\n", count)
			fmt.Println()
		case 3:
			fmt.Println("删除用户信息")
		case 4:
			fmt.Println("修改用户信息")
		case 5:
			fmt.Printf("系统将在5秒后退出...欢迎再次使用！\n")
			time.Sleep(5 * time.Second)
			os.Exit(200)
		default:
			fmt.Println("输入有误！请重新输入！")

		}

	}

}
