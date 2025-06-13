package main

/*
重要知识点！！！
字节byte 转换为 字符串 使用 string 进行转换 ---- 对应 读操作
字符串string 转换为 字节 使用 []byte 进行转换 ---- 对应 写操作

*/

import (
	"os"
)

func writeBytesOrStr(file *os.File) {
	var s = `打开窗子吧，让自由的空气进来，呼吸一下英雄们的气息。
	    人生是艰苦的，对于不甘于平凡的人，那是一场无日无夜的斗争，贫穷，日常的烦闷。。。深深的压着他们的心灵，没有一道欢乐之光。。。`
	// 当还行写操作时，file.write 只能写字节类型byte, 不能直接写字符串string,所以要将字符串转换为字节再写入
	file.Write([]byte(s))
	//当然也可以使用file.WriteString(s) 来写入

}
func writeFile() {
	//定义你的文件内容
	data := ` 今天我，寒夜里看雪飘过，怀着冷却了的心窝想远方，风雨里追赶，雾里看不清影踪，天空海阔你与我，可会变？
	
	
	`
	// 直接写整个文件
	os.WriteFile("allen_new.txt", []byte(data), 0666)

}
func main() {
	//写文件操作
	/*
		file, err := os.OpenFile("allen.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0777)
		if err != nil {
			fmt.Println(err)
		}

		writeBytesOrStr(file)
	*/
	writeFile()
}
