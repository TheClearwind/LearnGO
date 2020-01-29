package main

import "fmt"

func 阶乘(n int) int {
	if n == 1 {
		return 1
	} else {
		return n * 阶乘(n-1)
	}

}
func main() {
	fmt.Println(阶乘(8))
}
