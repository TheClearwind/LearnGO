package main

import "fmt"

func main() {
	//1.一个uint8 或者byte 代表一个ASCII字符
	//2.一个rune 类型代表一个 UTF-8 字符
	//当处理混合字符串时就需要用到rune 类型
	s3 := "hello 中国"
	for i := 0; i < len(s3); i++ {
		fmt.Printf("%c\n", s3[i]) //中文会乱码 字符在go 中是一个整形
	}
	// for range 是按照rune 类型进行遍历的
	for index, value := range s3 {
		fmt.Printf("%d:%c\n", index, value)
	}

	//字符串是无法被修改的 如果想要修改字符串 必须把先把字符串强转成byte数组或rune数组
	bs := []rune(s3)
	bs[6] = '北'
	bs[7] = '京'
	fmt.Println(string(bs))
}
