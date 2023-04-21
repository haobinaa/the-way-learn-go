package goroutines

import (
	"fmt"
	"time"
)

// 向 channel 中发送数据： ch <- data
func sendData(ch chan string) {
	for {
		ch <- "Washington"
		ch <- "Tripoli"
		ch <- "London"
		ch <- "Beijing"
		ch <- "Tokyo"
		time.Sleep(1e9)
	}
}

// 接收 channel 中数据: data <- channel
// 如果变量没有申明: int2 := <- channel
// <- channel 用来单独获取通道的值
func getData(ch chan string) {
	//var input string
	// 可以通过 for range 循环读取通道数据
	for data := range ch {
		fmt.Printf("%s ", data)
	}
	//for {
	//	input = <-ch
	//	fmt.Printf("%s ", input)
	//}
}

// 对于同一个 chan:
// 发送操作: 在接收者接收数据之前发送是不会结束(阻塞)
// 接收操作: 如果 chan 中没有数据是阻塞的
// 无限给通道发送数据
func pump(ch chan int) {
	for i := 0; ; i++ {
		ch <- i
	}
}

// 无限从通道中读取数据
func suck(ch chan int) {
	for {
		fmt.Println(<-ch)
	}
}

func f1(in chan int) {
	fmt.Println(<-in)
}

func ChanDeadLock() {
	out := make(chan int)
	fmt.Println("before")
	// 由于 chan 是无缓存的， 在这里就阻塞了
	out <- 2
	fmt.Println("after")
	// 抛出 panic 错误， 所有线程都休眠了
	go f1(out)
}
