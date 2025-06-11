// 指针可以简单理解为 内存地址
package main

import "fmt"

func main() {
	var a int
	a = 200
	fmt.Println(a)
	// & 操作符 为获取变量地址的操作符
	fmt.Println(&a) // 打印结果为 a 变量对应的内存地址 0xc000104020 , 0x 代表16进制

	//将 a 对应的内存地址存放在变量 p 中
	var p *int // 因为 a 是int 类型，所以 p 的类型是 *int , 如果 a 是string 那么 p 的类型是 *string
	p = &a     // &a 为指针 或内存地址，p 为指针变量
	fmt.Println(p)

	//练习字符串型
	var str1 string
	str1 = "I love kate"
	fmt.Printf("str1对应的内存地址为:%p\n", &str1)
	fmt.Println(&str1)

	var p1 *string // 因为str1 是string 型，所以用于存放内存地址的变量p1 类型会是 *string
	p1 = &str1     // &str1 为指针， p1为指针变量
	fmt.Println(p1)

	// 指针变量的取值操作
	// 根据指针变量获取值

	fmt.Println(*p)  // 打印结果为 200 即为之前存放的整型变量
	fmt.Println(*p1) // 打印结果为 I love kate 即为之前存放的字符串变量

	// 值拷贝案例
	var num1 int
	num1 = 800
	num2 := num1
	num1 = 1200
	fmt.Printf("num2的值为：%d\n", num2) // 打印结果num2 依然为 800 ， num1 和 num2 具有完全不同的独立的内存地址，所以num2 保留了num1 之前的值
	fmt.Printf("num1的值为：%d\n", num1)
	fmt.Println(&num1) //num1 的内存地址为 0xc000098040
	fmt.Println(&num2) //num2 的内存地址为 0xc000098048

	// 指针拷贝案例
	var num3 int
	num3 = 300
	num4 := &num3      //num4 拿到的是num3对应的内存地址
	fmt.Println(*num4) //打印结果为 300
	num3 = 789
	fmt.Println(*num4) //打印结果为 789
	num3 = 9700
	fmt.Println(*num4) //打印结果为 9700 可以看到指针copy 的值会随着原始值的改变而改变

	// 通过指针地址改变变量存储的值
	var num5 int
	num5 = 99
	num6 := &num5 //num6存储的是num5的内存地址
	*num6 = 888
	fmt.Println(num5) //输出结果为 888

}
