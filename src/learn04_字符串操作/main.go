package main

import (
	"fmt"
	"strings"
)

func main() {
	//go 语言中 字符串使用“”，字符使用‘’ 使用的是UTF-8编码 一个中文也算一个字符
	//c := 'a'
	//c_1 := "中"
	//s := "I am handsome"
	//fmt.Println(c)
	//fmt.Println(c_1)
	//fmt.Println(s)

	// 多行字符串 使用反引号 就是ESC下面哪个键
	s2 := `世情薄
人情恶
雨送黄昏花易落`
	fmt.Println(s2)

	//字符串相关操作
	//1.长度
	fmt.Println(len(s2))
	//2.字符串拼接
	name := "lambda"
	desc := " dsb"
	fmt.Println(name + desc)
	s := fmt.Sprintf("%s%s", name, desc)
	fmt.Println(s)
	//3.字符串分割
	file_path := `https:/www.bilibili.com/video/av73381776?p=14`
	ret := strings.Split(file_path, "/")
	fmt.Println(ret)
	//4.包含
	fmt.Println(strings.Contains(s2, "花"))
	fmt.Println(strings.Contains(s2, "草"))
	//5.前后缀判断
	fmt.Println(strings.HasPrefix(s2, "世")) // 前缀
	fmt.Println(strings.HasSuffix(s2, "世")) // 后缀

	//6. 字串出现的位置
	fmt.Println(strings.Index("lambda", "a"))
	fmt.Println(strings.LastIndex("lambda", "a"))

	//7.拼接
	fmt.Println(strings.Join(ret, ""))
}
