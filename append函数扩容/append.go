package main

import "fmt"

func main() {
	var cities = make([]string, 3, 5) // 初始化切片
	cities[0] = "上海"
	cities[1] = "南京"
	cities[2] = "北京"
	fmt.Println(cities)
	// 添加一个新城市杭州
	new_cities := append(cities, "杭州", "深圳")
	fmt.Println(new_cities)
	fmt.Printf("new_cities的类型：%T\n", new_cities) // 依然是切片类型 []string

	fmt.Println(cities) //cities 切片并没有改变
	cities[0] = "莫斯科"
	fmt.Println(cities)
	fmt.Println(new_cities)

	new_cities_1 := append(cities, "巴黎", "米兰", "柏林", "伦敦", "伯明翰", "斯图加特", "旧金山", "洛杉矶")
	fmt.Printf("打印超出容量的新切片\n")
	fmt.Println(new_cities_1)

	// 修改原始切片的值
	cities[1] = "夏威夷"
	fmt.Println("修改后的原始切片")
	fmt.Println(cities) // [莫斯科 夏威夷 北京]
	fmt.Println("打印超出容量的切片")
	fmt.Println(new_cities_1) // [莫斯科 南京 北京 巴黎 米兰 柏林 伦敦 伯明翰 斯图加特 旧金山 洛杉矶]
	fmt.Println("打印没有超出容量的切片")
	fmt.Println(new_cities) // [莫斯科 夏威夷 北京 杭州 深圳]

	// 得出的结论
	// 如果添加的数据超出切片的容量，则会开发完全独立的存储空间，成为完全独立的两个切片
	// 如果添加的数据没有超出切片的容量，则会和之前的切片共享存储空间，修改之前的元素，新的切片也会改变

}
