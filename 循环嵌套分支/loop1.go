// 循环嵌套语句
package main

import "fmt"

func main() {
	for true {
		var age int
		fmt.Println("请在此输入您的年龄：")
		fmt.Scan(&age)
		// 判断年龄，小于18岁不可以玩这种游戏
		if age >= 18 {
			fmt.Println("欢迎进入星球大战游戏！")
			info := `
			输入“1”：则会输出“攻击敌方坦克”
			输入“2”：则会输出“攻击敌方飞机”
			输入“3”：则会输出“攻击敌方弹药库”
			输入“4”：则会输出”补给我放弹药“
			输入“5”：则会输出“添加士兵血量”
			输入“6”：则会输出“占领敌方后勤部”

			`
			fmt.Println(info)
			var choice int
			fmt.Println("请在此输入您的选择：")
			fmt.Scan(&choice)
			switch choice {
			case 1:
				fmt.Println("攻击敌方坦克")
			case 2:
				fmt.Println("攻击敌方飞机")
			case 3:
				fmt.Println("攻击敌方弹药库")
			case 4:
				fmt.Println("补给我方弹药")
			case 5:
				fmt.Println("添加士兵血量")
			case 6:
				fmt.Println("占领敌方后勤部")
			default:
				fmt.Println("输入有误，请重新输入")
			}

		} else {
			fmt.Println("Sorry，未成年人不可以玩星球大战游戏哦！")
			break
		}
	}

}
