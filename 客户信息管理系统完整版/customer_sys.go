/*
author: Allen Ma
date: 2025/06/12
location: Shanghai
*/
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
			//创建一个新的Map对象
			var customer_add = map[string]string{"name": customer_name, "age": customer_age, "gender": customer_gender, "height": customer_height, "contact": customer_mobile}
			//将Map对象放入大切片里
			my_customer_slice = append(my_customer_slice, customer_add)
			count++
			fmt.Printf("添加成功！共添加%d条记录\n", count)
			fmt.Println()
		case 3:
			fmt.Println("删除用户信息")
			var delName string
			fmt.Println("请输入要删除的用户的姓名：")
			fmt.Scan(&delName)
			for i := 0; i < len(my_customer_slice); i++ {
				if my_customer_slice[i]["name"] == delName {
					my_customer_slice = append(my_customer_slice[:i], my_customer_slice[i+1:]...)
					i--
				}
			}
			fmt.Println("删除成功！")
		case 4:
			fmt.Println("修改用户信息")
			mini_panel := `
			1. 修改用户姓名
			2. 修改用户性别
			3. 修改用户年龄
			4. 修改用户身高
			5. 修改用户联系方式
			6. 完成操作
			`
			for {
				fmt.Println(mini_panel)
				var mini_choice int
				fmt.Println("请输入您的选择(1-5):")
				fmt.Scan(&mini_choice)
				switch mini_choice {
				case 1:
					fmt.Println("修改用户姓名")
					var modName string
					fmt.Println("请输入要修改的姓名：")
					fmt.Scan(&modName)
					var newName string
					fmt.Println("请输入新姓名：")
					fmt.Scan(&newName)
					for i := 0; i < len(my_customer_slice); i++ {
						if my_customer_slice[i]["name"] == modName {
							my_customer_slice[i]["name"] = newName

						}
					}
					fmt.Println("正在修改，请稍等...")
					time.Sleep(3 * time.Second)
					fmt.Println("修改完成！")

				case 2:
					fmt.Println("修改用户性别")
					var gen_name string
					fmt.Println("请输入需要修改性别的用户姓名：")
					fmt.Scan(&gen_name)
					var mod_gender string
					fmt.Println("请输入新的性别：")
					fmt.Scan(&mod_gender)
					for i := 0; i < len(my_customer_slice); i++ {
						if my_customer_slice[i]["name"] == gen_name {
							my_customer_slice[i]["gender"] = mod_gender
						}
					}
					fmt.Println("正在修改，请稍等...")
					time.Sleep(3 * time.Second)
					fmt.Println("修改完成！")
				case 3:
					fmt.Println("修改用户年龄")
					var age_name string
					fmt.Println("请输入需要修改年龄的用户姓名：")
					fmt.Scan(&age_name)
					var newAge string
					fmt.Println("请输入新的年龄：")
					fmt.Scan(&newAge)
					for i := 0; i < len(my_customer_slice); i++ {
						if my_customer_slice[i]["name"] == age_name {
							my_customer_slice[i]["age"] = newAge
						}
					}
					fmt.Println("正在修改，请稍等...")
					time.Sleep(3 * time.Second)
					fmt.Println("修改完成！")

				case 4:
					fmt.Println("修改用户身高")
					var hei_name string
					fmt.Println("请输入需要修改身高的用户姓名：")
					fmt.Scan(&hei_name)
					var newHeight string
					fmt.Println("请输入新的身高：")
					fmt.Scan(&newHeight)
					for i := 0; i < len(my_customer_slice); i++ {
						if my_customer_slice[i]["name"] == hei_name {
							my_customer_slice[i]["height"] = newHeight
						}
					}
					fmt.Println("正在修改，请稍等...")
					time.Sleep(3 * time.Second)
					fmt.Println("修改完成！")
				case 5:
					fmt.Println("修改用户联系方式")
					var phone_name string
					fmt.Println("请输入需要修改电话号码的用户姓名：")
					fmt.Scan(&phone_name)
					var newPhone string
					fmt.Println("请输入用户新的电话号码：")
					fmt.Scan(&newPhone)
					for i := 0; i < len(my_customer_slice); i++ {
						if my_customer_slice[i]["name"] == phone_name {
							my_customer_slice[i]["contact"] = newPhone
						}
					}
					fmt.Println("正在修改，请稍等...")
					time.Sleep(3 * time.Second)
					fmt.Println("修改完成！")

				}
				if mini_choice == 6 {
					break
				}
			}
		case 5:
			fmt.Printf("系统将在5秒后退出...欢迎再次使用！\n")
			time.Sleep(5 * time.Second)
			os.Exit(200)
		default:
			fmt.Println("输入有误！请重新输入！")

		}

	}

}
