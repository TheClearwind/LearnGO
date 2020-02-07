package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func readFile1(file *os.File) {
	temp := make([]byte, 128) // 设置缓存区大小
	for {
		n, err := file.Read(temp)
		if err != nil {
			fmt.Printf("读取数据失败:%v\n", err)
		} else {
			fmt.Println(string(temp[:n]))
			if n < 128 {
				break
			}
		}
	}
}
func readFileByBufio(file *os.File) {
	//利用bufio 读取文件
	reader := bufio.NewReader(file)
	for {
		data, err := reader.ReadString('\n')
		if err == io.EOF {
			return
		}
		if err != nil {
			fmt.Printf("读取数据失败:%v\n", err)
			return
		}
		fmt.Print(data)
	}
}
func readFileByIOUtil(path string) {
	// 一次性全部读取完毕
	data, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Printf("文件读取失败:%v\n", err)
	} else {
		fmt.Print(string(data))
	}
}
func writeFile(path string) {
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 1) //第三个参数表示文件权限 windows下无效
	if err != nil {
		fmt.Printf("文件打开失败:%v\n", err)
		return
	}
	file.Write([]byte("王兰花\n")) //byte写入
	file.WriteString("哈哈哈哈哈")   //字符串写入
	defer file.Close()
}
func writeFileByBufio(path string) {
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 1) //第三个参数表示文件权限 windows下无效
	if err != nil {
		fmt.Printf("文件打开失败:%v\n", err)
		return
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	writer.WriteString("哈哈哈哈哈")
	defer writer.Flush() // 需要清空缓存区 否则不会写入硬盘

}

func writeFileByIoutil(path string) {
	var content = "哈哈哈哈哈"
	ioutil.WriteFile(path, []byte(content), 1)
}
func main() {
	//file, err := os.Open("./main.go")
	//if err != nil {
	//	fmt.Printf("文件读取失败: %v\n", err)
	//} else {
	//	readFileByIOUtil("./main.go")
	//}
	//defer file.Close() //关闭文件

	//文件写入
	//writeFile("aaa.txt")
	//writeFileByBufio("bbb.txt")
	writeFileByIoutil("ccc.txt")
}
