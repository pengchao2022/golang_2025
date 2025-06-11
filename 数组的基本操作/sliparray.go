package main

import (
	"fmt"
	"reflect"
	"strings"
)

func main() {
	var cities = [...]string{"北京", "上海", "南京", "杭州", "广州", "深圳"}
	//获取数组的长度len()
	lenth := len(cities)
	fmt.Println(lenth) //可以看到打印结果为 6

	//查看数组的类型
	fmt.Println(reflect.TypeOf(cities)) //打印结果为 [6]string
	//索引取数组里面的元素
	fmt.Println(cities[2]) //打印结果为：南京
	fmt.Println(cities[4]) //打印结果为：广州

	//数组的切片
	//[开始索引:结束索引]
	//切片是顾头不顾尾的
	fmt.Println(cities[1:4]) //打印出 [上海 南京 杭州] 切出来的结果是新的数据类型，叫作切片类型
	var ret1 = cities[1:4]
	fmt.Println(reflect.TypeOf(ret1)) // 打印结果为 []string 注意，没有长度的叫做切片类型，不是数组类型，数组类型需要带有长度

	// 缺省取最后或者最前端
	fmt.Println(cities[2:]) // 表示取到最后 打印结果为：[南京 杭州 广州 深圳]

	// 缺省前端，从最前端开始取
	fmt.Println(cities[:6]) // 表示从索引0开始取 打印结果为：[北京 上海 南京 杭州 广州 深圳]

	// 切片没有长度，切片是可伸缩的， 而数组有长度限制

	// 数组的循环

	for i := 0; i < len(cities); i++ {
		fmt.Println(cities[i])
	}

	// 遍历循环使用for range
	for i := range cities {
		fmt.Println(i, cities[i])
	}

	// 遍历使用两个变量i,v , v 为每一个元素的值，i为索引
	for i, v := range cities {
		fmt.Println(i, v)
	}

	// 匿名变量
	for _, v := range cities {
		fmt.Println(v)
	}
	// 只想打印上海这个元素
	for _, v := range cities {
		if v == "上海" {
			fmt.Println("只打印上海")
			fmt.Println(v)
		}
	}

	// 打印以”州“结尾的城市的名字
	fmt.Println("打印以“州”结尾的城市名字")
	for _, v := range cities {

		if strings.HasSuffix(v, "州") {

			fmt.Println(v)
		}
	}

	// 修改数组里面元素的值
	// 将姓名里面姓“李"的人全部改为匿名
	// 注意，这里使用names[i] = “匿名” 而不能使用 v = “匿名”

	var names = [...]string{"李飞", "李婷", "李娜", "周杰伦", "周华健", "李军", "李凡", "董卓", "吕布"}
	var count = 0
	for i, v := range names {
		if strings.HasPrefix(v, "李") {
			names[i] = "匿名"
			count++

		}
	}
	fmt.Printf("共找到 %d 个姓李的员工\n", count)
	fmt.Println(names) // 打印结果为：[匿名 匿名 匿名 周杰伦 周华健 匿名 匿名 董卓 吕布]

}
