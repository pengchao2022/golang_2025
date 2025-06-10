//编写一个猜字游戏

package main

import "fmt"

func main() {
	my_num := 66
	count := 0
	for true {
		var guess_nu int
		fmt.Println("请输入您猜的数字1-100 之间：")
		fmt.Scan(&guess_nu)
		count++
		if guess_nu > my_num {
			fmt.Println("太大了")
		} else if guess_nu < my_num {
			fmt.Println("太小了")

		} else if guess_nu == my_num {
			fmt.Println("恭喜！猜对了！")
			break

		}

	}
	fmt.Printf("总共猜了%d 次", count)

}
