package csp

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"
)

// TestCSP CSP 通信模型
func TestCSP(t *testing.T) {
	stop := make(chan bool)
	var wg sync.WaitGroup
	// Spawn example consumers
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(stop <-chan bool) {
			defer wg.Done()
			consumer(stop)
		}(stop)
	}
	// 等待 5s 后发出结束信号
	time.Sleep(5 * time.Second)
	close(stop)
	fmt.Println("stopping all jobs!")
	// 等待所有子协程退出
	wg.Wait()
}

func TestContextControl(t *testing.T) {
	wg := &sync.WaitGroup{}
	values := []string{"https://www.baidu.com/", "https://www.zhihu.com/"}
	ctx, cancel := context.WithCancel(context.Background())

	for _, url := range values {
		wg.Add(1)
		subCtx := context.WithValue(ctx, favContextKey("url"), url)
		go reqURL(subCtx, wg)
	}

	go func() {
		time.Sleep(time.Second * 3)
		cancel()
	}()

	wg.Wait()
	fmt.Println("exit main goroutine")
}
