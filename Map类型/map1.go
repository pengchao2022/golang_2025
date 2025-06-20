package main

import "fmt"

func main() {
	// map 是go 语言唯一的一个映射数据类型
	//案例1 ，切片 和 Map 做对比
	// 切片实现
	customer_info := []interface{}{"allen", 22, 180, "male", 18510656167} //interface 表示接口，可代表任意数据类型
	fmt.Println(customer_info[0])                                         // 打印结果为：allen

	// 遍历切片中的索引和值
	for i, v := range customer_info {
		fmt.Println(i, v)
	}

	// map 实现
	var customer_info_1 = make(map[string]string) // key 为 string 类型， 值为 string 类型 需要使用make 函数初始化

	// 先声明 后初始化
	customer_info_1["name"] = "allen"
	customer_info_1["age"] = "22"
	customer_info_1["height"] = "180"
	customer_info_1["gender"] = "male"
	customer_info_1["mobile"] = "18510656167"

	// 打印出来
	fmt.Println(customer_info_1)
	// 打印姓名
	fmt.Println(customer_info_1["name"])

	// 声名并初始化
	var customer_info_2 = map[string]string{"name": "kate", "age": "21", "height": "175", "gender": "female", "mobile": "15010045761"}
	fmt.Println(customer_info_2)
	fmt.Println(customer_info_2["gender"])

	// 再练习一个声明并初始化
	var customer_info_3 = map[string]string{"name": "lily", "age": "26", "height": "181", "gender": "female", "mobile": "17278890861"}
	fmt.Println(customer_info_3)
	fmt.Println(customer_info_3["mobile"])

	// 声明一个新的切片，将三个map 变量放在这个切片里
	var my_customer_slice = []map[string]string{customer_info_1, customer_info_2, customer_info_3}
	fmt.Println(my_customer_slice)
	// 打印出lily 的年龄
	fmt.Println(my_customer_slice[2]["name"])
	fmt.Println(my_customer_slice[2]["age"])
	// 获取kate 的手机号
	fmt.Println(my_customer_slice[1]["mobile"])
	// 获取allen 的身高
	fmt.Println(my_customer_slice[0]["height"])

	// 打印出每个客户的信息
	for i, value := range my_customer_slice {
		fmt.Printf("客户ID：%d,客户姓名：%s,客户年龄：%s,客户身高：%s,客户性别：%s,客户手机号码：%s\n", i, value["name"], value["age"], value["height"], value["gender"], value["mobile"])
	}

	//添加一个客户的信息 比如 Sophia
	customer_info_4 := map[string]string{"name": "Sophia", "age": "35", "height": "181", "gender": "female", "mobile": "15710045261"}
	my_customer_slice = append(my_customer_slice, customer_info_4)
	fmt.Println(my_customer_slice)

	// 打印出每个客户的信息
	for _, value := range my_customer_slice {
		fmt.Printf("客户姓名：%s,客户年龄：%s,客户身高：%s,客户性别：%s,客户手机号码：%s\n", value["name"], value["age"], value["height"], value["gender"], value["mobile"])
	}

	// 遍历map 中的key ,value
	for key, value := range customer_info_3 {
		fmt.Println(key, value)
	}
	// Map 属于容量型数据，主要是增 删 改 查
	// 添加 key,value
	var cars = map[string]string{"brand": "BMW", "color": "yellow", "country": "Germany"}
	fmt.Println(cars)
	// 添加key 为 price
	cars["price"] = "320000"
	fmt.Println(cars)
	// 修改country 为China
	cars["country"] = "China"
	fmt.Println(cars)

	// 删除 key
	delete(cars, "country")
	fmt.Println(cars)
}
