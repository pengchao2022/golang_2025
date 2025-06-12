package main

import "fmt"

func main() {
	// Map嵌套 Map
	var student_1 = map[string]string{"name": "alen", "age": "24", "gender": "female", "height": "180"}
	var student_2 = map[string]string{"name": "kate", "age": "26", "gender": "female", "height": "175"}
	var student_3 = map[string]string{"name": "lily", "age": "27", "gender": "female", "height": "181"}
	var students = map[int]map[string]string{1001: student_1, 1002: student_2, 1003: student_3}
	//查询1003 学生的姓名
	//students 是map 类型，所以遵守 Key value 模式
	fmt.Println(students[1003]["name"]) // 打印结果为 lily
	//查询1002 学生的身高
	fmt.Println(students[1002]["height"]) //打印结果为 175
	//查询1001 学生的年龄
	fmt.Println(students[1001]["age"]) //打印结果为 24

	//Map嵌套切片slice
	var cars_1 = map[string]interface{}{"brand": "BMW", "color": "yellow", "country": "Germany", "series": []string{"BMW320", "BMW425", "BMW535"}}
	var cars = []map[string]interface{}{cars_1}
	//取出第一个车的第二个系列
	fmt.Println(cars[0]["series"].([]string)[1]) // 需要断言下

}
