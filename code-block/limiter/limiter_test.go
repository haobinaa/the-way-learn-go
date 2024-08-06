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
