package main

func main() {
	/*
		go 支持的类型
		bool
		数字类型：
			int8, int16, int32, int64, int
			uint8, uint16, uint32, uint64, uint
			float32, float64
			complex64, complex128
			byte
			rune
		string
	*/

	//bool 类型表示一个布尔值，值为 true 或者 false。 布尔值不参与运算 也 无法与其他类型进行转换 默认是false
	//a := false
	//b := true
	//fmt.Println("a:", a, "b:", b)
	//c := a && b
	//fmt.Println("c:", c)
	//d := a || b
	//fmt.Println("d:", d)

	/*
		有符号整型
		int8：表示 8 位有符号整型
		大小：8 位
		范围：-128～127

		int16：表示 16 位有符号整型
		大小：16 位
		范围：-32768～32767

		int32：表示 32 位有符号整型
		大小：32 位
		范围：-2147483648～2147483647

		int64：表示 64 位有符号整型
		大小：64 位
		范围：-9223372036854775808～9223372036854775807

		int：根据不同的底层平台（Underlying Platform），表示 32 或 64 位整型。除非对整型的大小有特定的需求，否则你通常应该使用 int 表示整型。
		大小：在 32 位系统下是 32 位，而在 64 位系统下是 64 位。
		范围：在 32 位系统下是 -2147483648～2147483647，而在 64 位系统是 -9223372036854775808～9223372036854775807。
	*/
	//var a = 89
	//b := 95
	//fmt.Println("value of a is", a, "and b is", b)
	////在 Printf 方法中，使用 %T 格式说明符（Format Specifier），可以打印出变量的类型。Go 的 unsafe 包提供了一个 Sizeof 函数，该函数接收变量并返回它的字节大小。unsafe 包应该小心使用，因为使用 unsafe 包可能会带来可移植性问题。
	//fmt.Printf("type of a is %T，size of a is %d",a,unsafe.Sizeof(a))

	/*无符号整型
	uint8：表示 8 位无符号整型
	大小：8 位
	范围：0～255

	uint16：表示 16 位无符号整型
	大小：16 位
	范围：0～65535

	uint32：表示 32 位无符号整型
	大小：32 位
	范围：0～4294967295

	uint64：表示 64 位无符号整型
	大小：64 位
	范围：0～18446744073709551615

	uint：根据不同的底层平台，表示 32 或 64 位无符号整型。
	大小：在 32 位系统下是 32 位，而在 64 位系统下是 64 位。
	范围：在 32 位系统下是 0～4294967295，而在 64 位系统是 0～18446744073709551615。
	*/
	/*
		浮点型
		float32：32 位浮点数
		float64：64 位浮点数
	*/

	//a, b := 5.67, 8.97
	//fmt.Printf("type of a %T b %T\n", a, b)
	//sum := a + b
	//diff := a - b
	//fmt.Println("sum", sum, "diff", diff)
	//
	//no1, no2 := 56, 89
	//fmt.Println("sum", no1+no2, "diff", no1-no2)
	//f1:=1.23
	//fmt.Printf("%T\n",f1) //默认是float64
	//f2:=float32(1.23) //显示声明float32
	////f1=f2 会报错 go语言中类型不一就不能赋值 没有隐式转换
	//fmt.Printf("%T",f2)

}
