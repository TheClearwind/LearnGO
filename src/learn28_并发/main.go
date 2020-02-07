package main

import (
	"fmt"
	"sync"
)

//goroutine

func say() {
	fmt.Println("hello world")
	wg.Done() //通知wg goroutine结束
}
func produce(ch chan<- int) {
	//chan <- 将ch1限定为只进不出 同理 <-chan 限定为只出不进
	for i := 0; i < 100; i++ {
		ch <- i
	}
	close(ch)
}
func consumer(ch1 <-chan int, ch2 chan<- int) {
	// 从通道中取值
	for {
		ret, ok := <-ch1 //当ch1为空且被关闭时 ok才为false 也就是说 即使ch1被关闭但是还有数据 ok也不是false
		if !ok {
			break
		} else {
			ch2 <- ret * ret
		}
	}
	close(ch2)
}

var wg sync.WaitGroup

func main() {
	//wg.Add(1) //添加一个
	//go say() // 开启一个goroutine 去执行 类似于线程
	//fmt.Println("hello go!")  // main goroutine 执行结束后整个就结束了 say goroutine 被强制关闭
	//// 使用 waitgroup 等待
	//wg.Wait()

	// channel 用于 goroutine 数据交互
	//var c chan int //var name chan 元素类型
	//c = make(chan int,1) // 初始化 make(type，缓存区大小)
	//c<-10 //向通道中发送数据
	//a:= <-c // 从通道中取值
	//fmt.Println(a)
	//close(c) //关闭通道

	// 生产者消费者模式
	//ch1 := make(chan int, 100)
	//ch2 := make(chan int, 100)
	//go produce(ch1)
	//go consumer(ch1, ch2)
	//
	//// 使用for range也可以从通道中取值
	//for ref := range ch2 {
	//	fmt.Println(ref)
	//}
	//
	//select 多路复用

	ch := make(chan int, 20)
	for i := 0; i < 10; i++ {
		select { // select 语句每次只有一个case会被执行 如果多个case都满足则随机选择一个case执行
		case data := <-ch: //case 后面是发送，赋值等操作 只有该操作能够执行才会运行case
			fmt.Println(data)
		case ch <- i:
		default:
			fmt.Println("哈哈哈哈哈")
		}
	}
}
