package goroutines

import (
	"fmt"
	"time"
)

var (
	// 申明 channel 只接受数据, 这种 chan 无法关闭
	send_only chan<- int
	// 申明 channel 只发送数据
	recv_only <-chan int
)

func sum(x, y int, c chan int) {
	c <- x + y
}

// Semaphore 信号量模式
func Semaphore() {
	c := make(chan int)
	go sum(12, 13, c)
	// 等待上面协成结果
	fmt.Println(<-c)
}

// send 数据到 chan
func generate() chan int {
	ch := make(chan int)
	go func() {
		for i := 2; ; i++ {
			ch <- i
		}
	}()
	return ch
}

// 过滤能被 prime 除尽的数据到 out
func filter(in chan int, prime int) chan int {
	out := make(chan int)
	go func() {
		for {
			if i := <-in; i%prime != 0 {
				out <- i
			}
		}
	}()
	return out
}

// 过滤
func sieve() chan int {
	out := make(chan int)
	go func() {
		ch := generate()
		for {
			prime := <-ch
			// 为每个数创建一个 chan(过滤能被自身整除的数， 即素数)
			ch = filter(ch, prime)
			out <- prime
		}
	}()
	return out
}

// ChanFilter The prime sieve: Daisy-chain filter processes together.
func ChanFilter() {
	primes := sieve()
	for {
		fmt.Println(<-primes)
	}
}

// CheckChanClosed close(ch) 用来关闭 cha
func CheckChanClosed(ch chan int) {
	data, ok := <-ch
	if ok {
		fmt.Println(data)
	} else {
		fmt.Println("chan is closed")
	}
}

func pump1(ch chan int) {
	for i := 0; ; i++ {
		ch <- i * 2
	}
}

func pump2(ch chan int) {
	for i := 0; ; i++ {
		ch <- i + 5
	}
}

func suck1(ch1, ch2 chan int) {
	for {
		// select 会随机选择一个就绪的 chan， 如果都阻塞了， 则等待其中一个可以处理
		// 在任何一个 case 中 break 或 return 就结束了
		select {
		case v := <-ch1:
			fmt.Printf("Received on channel 1: %d\n", v)
		case v := <-ch2:
			fmt.Printf("Received on channel 2: %d\n", v)
		}
	}
}

// ChanSelect chan select 模式
func ChanSelect() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go pump1(ch1)
	go pump2(ch2)
	go suck1(ch1, ch2)

	time.Sleep(1e9)
}

// TimeTickerUsage 计时器 ticker
func TimeTickerUsage() {
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
