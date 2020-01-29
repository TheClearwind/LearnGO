package main

import "fmt"

func main() {
	////声明数组
	//var a [3]int     // [3]int 是一种类型
	//a = [3]int{1, 2} //初始化
	//var b = [3]float32{1, 2, 3}
	//fmt.Println(a)
	//fmt.Println(b)
	//
	////推断长度
	//var c = [...]int{1, 2, 3, 4, 5, 6} //[] 需要有...否则就是一个切片
	//fmt.Println(c)
	//
	////索引初始化
	//var d [10]int
	//d = [10]int{5: 100} //索引为5的置为100 其余为0
	//fmt.Println(d)
	//
	////利用索引取值
	//fmt.Println(d[5])
	//
	////遍历
	//for i:=0;i<len(d);i++{
	//	fmt.Println(d[i])
	//}
	//for idx,value:=range d{
	//	fmt.Println(idx,":",value)
	//}

	//多维数组
	var a = [3][2]int{ //长度为3 类型为[2]int 的数组
		{1, 2},
		{3, 4},
	}
	fmt.Println(a)
	fmt.Println("****************************")
	//fmt.Println(a[1][1])
	//遍历
	for _, value := range a {
		for _, v := range value {
			fmt.Print(v, " ")
		}
		fmt.Print("\n")
	}
	fmt.Println("****************************")
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[i]); j++ {
			fmt.Print(a[i][j], " ")
		}
		fmt.Println()
	}
}
