package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// 定义按照字节读取的函数
func readBytes(file *os.File) {
	//声明一个切片
	data := make([]byte, 18) //18代表6个字符 3的倍数
	file.Read(data)
	fmt.Println(data)
	fmt.Println(string(data))

}

// 定义按照行来读取 读出来的需要转换为string
func readByLine(file *os.File) {
	reader := bufio.NewReader(file)
	data, _, _ := reader.ReadLine()
	fmt.Println(string(data)) // 只能打印一行

}

// 定义按照行来读取 读出来的不需要转换为string
func readByline_high(file *os.File) {
	reader := bufio.NewReader(file)
	data, _ := reader.ReadString('\n')
	fmt.Println(data)
}

// 定义按照行来读取整个文件内容
func readByline_whole(file *os.File) {
	reader := bufio.NewReader(file)
	for true {
		data, err := reader.ReadString('\n')
		fmt.Print(data)
		if err == io.EOF {
			break

		}

	}
}
func readFile() {
	data, _ := ioutil.ReadFile("/Users/pengchaoma/golang/golang_2025/文件读操作/七里香.txt")
	fmt.Println(string(data))
}

func main() {
	//打开文件
	file, err := os.Open("/Users/pengchaoma/golang/golang_2025/文件读操作/七里香.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(file) // 就可以拿到文件句柄 &{0xc000068120}
	// 读文件
	// 按照字节读
	//readBytes(file)
	// 按照行来读取
	//readByLine(file)
	// 按照行读取不用转换string
	//readByline_high(file)
	// 按行读取整个文件内容
	//readByline_whole(file)
	//适用于小文件的整个读取
	readFile() // 不需要用到句柄

}
