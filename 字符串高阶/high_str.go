package main

import (
	"fmt"
	"strings"
)

func main() {
	// 统计字符串的长度
	str1 := "This is the day I have been looking forward for two and three years!"
	fmt.Println(len(str1))
	fmt.Printf("字符串长度为：%d\n", len(str1))

	// 字母变为大写
	str1_upper := strings.ToUpper(str1)
	fmt.Println(str1_upper)

	// 字母变为小写
	str1_lower := strings.ToLower(str1)
	fmt.Println(str1_lower)

	// 判断字符串里以特定开头的字符
	name1 := "allen kate lily"
	test1 := strings.HasPrefix(name1, "allen")
	fmt.Println(test1)

	//判断以特定字母结尾的
	name2 := "I love China"
	test2 := strings.HasSuffix(name2, "China")
	fmt.Println(test2)

	//判断有没有包含特定字符的
	str2 := "this is ios26, designed by apple in california, assembled in China!"
	test3 := strings.Contains(str2, "2")
	fmt.Println(test3)

	// 字符串去掉空格
	username := "   allen.ma@gmail.com "
	fmt.Println(username)
	new_name := strings.Trim(username, " ")
	fmt.Println(new_name)

	//索引函数，找出字符出现在字符串中的位置
	str3 := "After a landslide election victory in 2024, President Donald J. Trump is returning to the White House to build upon his previous successes and use his mandate to reject the extremist policies of the radical left while providing tangible quality of life improvements for the American people."
	word_location := strings.Index(str3, "Trump")
	fmt.Println(word_location)
	no_word := strings.Index(str3, "allen")
	// 输出值为 -1 代表没有找到，及不存在该字符
	fmt.Println(no_word)

	// 分割字符串
	cities := "北京 上海 广州 成都 西安 拉萨 兰州 南京 深圳 石家庄"
	cut_city := strings.Split(cities, " ")
	fmt.Println(cut_city)
	// 统计列表里面的成市区个数
	fmt.Println(len(cut_city))
	// 打印结果 [北京 上海 广州 成都 西安 拉萨 兰州 南京 深圳 石家庄]
	// 统计结果 10

	//字符的拼接，将列表里面的字符按照逗号拼接起来
	new_string1 := strings.Join(cut_city, ",")
	fmt.Println(new_string1)
	// 打印结果 北京,上海,广州,成都,西安,拉萨,兰州,南京,深圳,石家庄

}
