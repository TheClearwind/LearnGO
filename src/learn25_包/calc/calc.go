package calc

func Add(x,y int) int {
	return x+y
}
var a=10 // 只有首字母大写的才能被外面的包访问到

func init() {
	// 包导入时自动执行  无参数 也 没有 返回值 在全局变量声明之后 main调用之前 执行
}
