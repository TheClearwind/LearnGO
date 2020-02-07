package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	//fmt.Println(now)
	//fmt.Println(now.Year())
	//fmt.Println(now.Month())
	//fmt.Println(now.Day())
	//
	//// 时间戳
	//
	//fmt.Println(now.Unix())
	//
	//// 将一个时间戳转换成时间
	//fmt.Println(time.Unix(12346487, 0))
	//
	////时间相加
	//fmt.Println(now.Add(time.Minute * 40))

	//定时器
	//timer:=time.Tick(time.Second)
	//for t:=range timer{
	//	fmt.Println(t)
	//}

	// 自恋的格式化方式 以go语言诞生时期作为格式化 标准 2006-1-2 15-04-05
	fmt.Println(now.Format("2006-1-2 15-4-5"))

	//parse 以及时区
	time_str := "2020-02-03 17:51:00"
	fmt.Println(time.Parse("2006-01-02 15:4:5", time_str)) // 解析的是UTC
	loc, _ := time.LoadLocation("Asia/Shanghai")
	fmt.Println(time.ParseInLocation("2006-01-02 15:4:5", time_str, loc))
}
