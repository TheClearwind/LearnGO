package main

import "fmt"

import (
	"fmt"
	sss "learn25_包/calc" // 起别名 需要配置go path
)
func main() {
	fmt.Println(sss.Add(10,20))
}
