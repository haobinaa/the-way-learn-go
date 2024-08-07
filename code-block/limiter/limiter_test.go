package limiter

import (
	"fmt"
	"testing"
	"time"
)

// TestFixedLimiter 固定窗口
func TestFixedLimiter(t *testing.T) {
	limiter := NewFixedWindowLimiter(1*time.Second, 3) // 每秒最多允许3个请求
	for i := 0; i < 15; i++ {
		now := time.Now().Format("15:04:05")
		if limiter.AllowRequest() {
			fmt.Println(now + " 请求通过")
		} else {
			fmt.Println(now + " 请求被限流")
		}
		time.Sleep(100 * time.Millisecond)
	}
}

// TestSlidingWindow 滑动窗口
func TestSlidingWindow(t *testing.T) {
	limiter := NewSlidingWindowLimiter(500*time.Millisecond, 2) // 每秒最多允许4个请求
	for i := 0; i < 15; i++ {
		now := time.Now().Format("15:04:05")
		if limiter.AllowRequest() {
			fmt.Println(now + " 请求通过")
		} else {
			fmt.Println(now + " 请求被限流")
		}
		time.Sleep(100 * time.Millisecond)
	}
}

func TestLeakyBucket(t *testing.T) {
	// 创建一个漏桶，速率为每秒处理3个请求，容量为4个请求
	leakyBucket := NewLeakyBucket(3, 4)

	// 模拟请求
	for i := 1; i <= 15; i++ {
		now := time.Now().Format("15:04:05")
		if leakyBucket.Allow() {
			fmt.Printf(now+"  第 %d 个请求通过\n", i)
		} else {
			fmt.Printf(now+"  第 %d 个请求被限流\n", i)
		}
		time.Sleep(200 * time.Millisecond) // 模拟请求间隔
	}
}

func TestTokenBucket(t *testing.T) {
	tokenBucket := NewTokenBucket(2.0, 3.0)

	for i := 1; i <= 10; i++ {
		now := time.Now().Format("15:04:05")
		if tokenBucket.Allow() {
			fmt.Printf(now+"  第 %d 个请求通过\n", i)
		} else { // 如果不能移除一个令牌，请求被拒绝。
			fmt.Printf(now+"  第 %d 个请求被限流\n", i)
		}
		time.Sleep(200 * time.Millisecond)
	}
}
