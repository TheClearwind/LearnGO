package main

import (
	"fmt"
	"sync"
	"time"
)

//sync once
var once sync.Once // 确保某项操作只做一次 比如关闭通道 加载配置文件 等只执行一次的操作
// 互斥锁
var lock sync.Mutex
var x = 0
var rwlock sync.RWMutex // 读写互斥锁 读取的时候不会阻塞 修改会阻塞
func f1() {
	//互斥锁
	wg.Add(2)
	go add()
	go add()
	wg.Wait()
	fmt.Println(x)
}
func write() {
	rwlock.Lock()
	x += 1
	time.Sleep(time.Microsecond * 5)
	rwlock.Unlock()
	wg.Done()
}
func read() {
	rwlock.RLock()
	time.Sleep(time.Microsecond)
	fmt.Println(x)
	rwlock.RUnlock()
	wg.Done()
}
func f2() {
	//读写互斥锁
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go write()
	}
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go read()
	}
}
func add() {
	for i := 0; i < 50000; i++ {
		lock.Lock()
		x++
		lock.Unlock()
	}
	wg.Done()
}

var wg sync.WaitGroup

func main() {
	start := time.Now()
	f2()
	wg.Wait()
	end := time.Now().Sub(start)
	fmt.Println(end)
	fmt.Println(x)
}
