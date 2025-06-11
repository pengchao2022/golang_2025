package main

import "fmt"

func main() {
	// 直接生命方式产生切片
	// 切片和数组的直接区别在于有没有声明长度
	var my_slip = []string{"kate", "lily", "jim", "sophia", "emily"}
	fmt.Printf("打印类型：%T\n", my_slip) // 打印出 []string 为切片，没有长度的不是数组类型

	// 练习1
	s1 := []int{1, 2, 3} // 声明一个切片
	s2 := s1[1:]         // s1 从索引1开始切到最后得到新的切片s2 [2,3]
	s2[1] = 4            // s2 索引1 为3 ，就是把 3改为 4
	fmt.Println(s1)      // 修改后 s1 为[1,2,4]

	// 练习2
	var a = []int{1, 2, 3} // 切片会自动构造底层数组[1,2,3]
	b := a
	a[0] = 100
	fmt.Println(b) // 打印结果为[100 2 3]

}
