package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// context 用来控制goroutine

func f1(ctx context.Context) {
loop:
	for {
		fmt.Println("王兰花")
		time.Sleep(500 * time.Millisecond)
		select {
		case <-ctx.Done():
			break loop
		default:
		}

	}
	wg.Done()
}

var wg sync.WaitGroup

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	wg.Add(1)
	go f1(ctx)
	time.Sleep(3 * time.Second)
	cancel() // 关闭子goroutine
	wg.Wait()
	fmt.Println("退出")

}
