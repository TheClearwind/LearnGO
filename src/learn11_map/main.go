package main

import "fmt"

func main() {
	//定义 map[keytype]valuetype map和slice 不初始化不能用
	var m1 = make(map[string]int, 5) // 尽量避免它动态扩容
	m1["lambda"] = 18
	m1["王兰花"] = 25
	fmt.Println(m1)

	//访问内容
	fmt.Println(m1["王兰花"])

	//访问不存在的值
	fmt.Println(m1["李美丽"]) // 返回对应类型的0值

	//更加健壮的写法
	v, ok := m1["三狗"]
	if !ok {
		fmt.Println("key error")
	} else {
		fmt.Println(v)
	}

	//map 的遍历
	for k, v := range m1 {
		fmt.Println(k, ":", v)
	}

	// 只遍历 key
	for k := range m1 {
		fmt.Println(k)
	}

	// 只遍历 v
	for _, v := range m1 {
		fmt.Println(v)
	}

	// 删除  若删除的key不存在 则什么也不做
	delete(m1, "王兰花")
	fmt.Println(m1)
}
