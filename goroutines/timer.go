// Package goroutines timer 各种使用
package goroutines

import (
	"fmt"
	"log"
	"time"
)

// waitChannel chan 设置超时， 而不是永久阻塞
func waitChannel(conn <-chan string) bool {
	timer := time.NewTimer(1 * time.Second) //设置超时时间 1s
	select {
	case <-conn:
		timer.Stop() // 接收到数据后，要停止计时器
		return true
	case <-timer.C: //超时判断
		fmt.Println("WaitChannel timeout!")
		return false
	}
}

func WaitChan() {
	timeoutChan := make(chan string)
	waitChannel(timeoutChan)
}

// DelayFunction 延迟执行
func DelayFunction() {
	fmt.Println(time.Now())
	timer := time.NewTimer(5 * time.Second)
	select {
	case <-timer.C:
		log.Println("Delayed 5s, start to do something.")
		// do something
	}
}

// TimerUsage timer 使用
func TimerUsage() {
	// NewTimer 创建一个 Timer，会在指定时间后向 chan 发送当前时间
	timer1 := time.NewTimer(5 * time.Second)
	fmt.Println("开始时间：", time.Now().Format("2006-01-02 15:04:05"))
	go func(t *time.Timer) {
		times := 0
		for {
			<-t.C
			fmt.Println("timer", time.Now().Format("2006-01-02 15:04:05"))
			// 从t.C中获取数据，此时time.Timer定时器结束。如果想再次调用定时器，只能通过调用 Reset() 函数来执行
			// Reset 使 t 重新开始计时，（本方法返回后再）等待时间段 d 过去后到期。
			// 如果调用时 t 还在等待中会返回真；如果 t已经到期或者被停止了会返回假。
			times++
			// 调用 reset 重发数据到chan C
			fmt.Println("调用 reset 重新设置一次timer定时器，并将时间修改为2秒")
			t.Reset(2 * time.Second)
			if times > 3 {
				fmt.Println("调用 stop 停止定时器")
				t.Stop()
			}
		}
	}(timer1)
	time.Sleep(30 * time.Second)
	fmt.Println("结束时间：", time.Now().Format("2006-01-02 15:04:05"))
	fmt.Println("ok")
}

// TimerTickerUsage 计时器 ticker
func TimerTickerUsage() {
	tick := time.Tick(1e8)
	boom := time.After(5e8)
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			return
		default:
			fmt.Println("    .")
			time.Sleep(5e7)
		}
	}
}
