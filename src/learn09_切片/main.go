package main

import "fmt"

func main() {
	//切片三要素:始地址 大小 容量

	//var s = []int{1,2,3,45,6}
	//fmt.Printf("%T\n",s)
	////从数组中得到切片
	//var a =[3]int{1,2,3}
	//var s1 = a[:2]
	//fmt.Println(s1)
	//切片相关
	//x := [...]int{1, 2, 3, 4, 5, 6, 4}
	//y := x[2:5]
	//fmt.Println("y的长度：", len(y)) // 切片的长度 目前有多少元素
	//fmt.Println("y的容量：", cap(y)) // 切片的容量，底层数组的长度 从idx=2的位置一直数到最后

	// 切片的容量机制
	//3种：
	// 1.每次只追加一个元素，每一次都是上一次的2倍
	// 2.追加的超过原来容量的1倍，就等于原来的容量+扩容元素个数的最接近的偶数
	// 3.如果切片的容量大于了1024，后续就每次扩容0.25倍
	//var x []int
	//fmt.Printf("x:%v len:%d  cap:%d  ptr:%p\n", x, len(x), cap(x), x)
	//x = append(x, 1)
	//fmt.Printf("x:%v len:%d  cap:%d  ptr:%p\n", x, len(x), cap(x), x)
	//x = append(x, 1)
	//fmt.Printf("x:%v len:%d  cap:%d  ptr:%p\n", x, len(x), cap(x), x)
	//x = append(x, 1)
	//fmt.Printf("x:%v len:%d  cap:%d  ptr:%p\n", x, len(x), cap(x), x)
	//x = append(x, 1)
	//fmt.Printf("x:%v len:%d  cap:%d  ptr:%p\n", x, len(x), cap(x), x)

	////切片是引用类型
	//x:=[]int{1,2,3}
	//fmt.Println(x)
	//y:=x
	//y[0]=200
	//fmt.Println(y)
	//fmt.Println(x)

	////一个坑爹的地方 内存问题
	//a:=[]int{1,2,3}
	//var c []int // 这里只是完成的声明 c无法容纳任何东西
	//num := copy(c,a)
	//fmt.Println(num)
	//
	////解决方案
	////1. 使用make(type,len,cap)
	//c = make([]int,3,3)
	//num = copy(c,a)
	//fmt.Println(num)
	//
	////2. 定义的时候 初始化 c:=[]int{0,0,0}

	// 最坑爹的地方 切片删除 没有内置方法 只能自己实现
	var s = []string{"北京", "上海", "广州", "深圳"}
	//删除 上海
	//步骤：1.把上海之前的元素取出来
	//		2.把上海之后的元素添加进去 就这样绕过了上海 等价于删除了上海
	s = append(s[:1], s[2:]...) //...操作把切片打散成一个个的元素
	fmt.Println(s)
}
