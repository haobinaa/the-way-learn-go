package context

import (
	"context"
	"fmt"
	"time"
)

// 全局变量，用于存储上下文信息
var (
	deadline  time.Time
	requestID string
)

// NoContext 没有 context 使用全局变量来传递
func NoContext() {
	// 设置上下文信息
	deadline = time.Now().Add(5 * time.Second)
	requestID = "123456"
	// 启动一个 goroutine 来处理任务
	go func() {
		for {
			select {
			case <-time.After(1 * time.Second):
				// 模拟一些耗时的操作
				fmt.Println("goroutine 1: doing some work")
			default:
				// 检查上下文信息，如果已经超时或被取消了，就退出循环
				if time.Now().After(deadline) {
					fmt.Println("goroutine 1: context canceled")
					return
				}
			}
		}
	}()

	// 启动另一个 goroutine 来处理任务
	go func() {
		for {
			select {
			case <-time.After(1 * time.Second):
				// 模拟一些耗时的操作
				fmt.Println("goroutine 2: doing some work")
			default:
				// 检查上下文信息，如果已经超时或被取消了，就退出循环
				if time.Now().After(deadline) {
					fmt.Println("goroutine 2: context canceled")
					return
				}
			}
		}
	}()

	// 等待一段时间，然后取消上下文信息
	time.Sleep(3 * time.Second)
	fmt.Println("main: context canceled")
	deadline = time.Now()
	time.Sleep(1 * time.Second)
}

// WithContext 用 context 来传递
func WithContext() {
	// 创建一个带有截止时间的 context
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(5*time.Second))
	defer cancel()
	// 启动一个 goroutine 来处理任务
	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				// 如果 context 被取消了，就退出循环
				fmt.Println("goroutine 1: context canceled")
				return
			default:
				// 模拟一些耗时的操作，普通情况可能是rpc调用
				time.Sleep(1 * time.Second)
				fmt.Println("goroutine 1: doing some work")
			}
		}
	}(ctx)
	// 启动另一个 goroutine 来处理任务
	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				// 如果 context 被取消了，就退出循环
				fmt.Println("goroutine 2: context canceled")
				return
			default:
				// 模拟一些耗时的操作
				time.Sleep(1 * time.Second)
				fmt.Println("goroutine 2: doing some work")
			}
		}
	}(ctx)
	// 等待一段时间，然后取消 context
	time.Sleep(3 * time.Second)
	cancel()
	fmt.Println("main: context canceled")
	time.Sleep(1 * time.Second)
}
