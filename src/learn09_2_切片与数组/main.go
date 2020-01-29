package main

import "fmt"

func main() {
	a := [...]int{1, 2, 3, 4, 5, 6, 7, 8}
	//b:=a[2:6]
	//c:=b[1:4]
	//fmt.Printf("b:%v len:%d cap:%d\n",b,len(b),cap(b)) //容量就是从始地址一直到底层数组最后的大小
	//fmt.Printf("c:%v len:%d cap:%d\n",c,len(c),cap(c))

	//恶心的地址问题
	//b := a[2:6]
	//c := b[:4]
	//fmt.Printf("b:%v len:%d cap:%d ptr:%p\n", b, len(b), cap(b), b) //b 和 c 地址相同
	//fmt.Printf("c:%v len:%d cap:%d ptr:%p\n", c, len(c), cap(c), b)
	b := a[:]
	a[5] = 500
	fmt.Println(a)
	fmt.Println(b)
}
