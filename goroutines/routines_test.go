package goroutines

import (
	"fmt"
	"testing"
	"time"
)

func TestSendChannel(t *testing.T) {
	ch := make(chan string)

	go sendData(ch)
	go getData(ch)

	time.Sleep(1e10)
}

func TestChanDeadLock(t *testing.T) {
	ChanDeadLock()
}

func TestSemaphore(t *testing.T) {
	Semaphore()
}

func TestChanFilter(t *testing.T) {
	ChanFilter()
}

// TestBenchmark benchmark
func TestBenchmark(t *testing.T) {
	fmt.Println(" sync", testing.Benchmark(BenchmarkChannelSync).String())
	fmt.Println("buffered", testing.Benchmark(BenchmarkChannelBuffered).String())
}

func TestSelectUse(t *testing.T) {
	SelectUse()
}

func TestChanSelect(t *testing.T) {
	ChanSelect()
}

func TestTimerUsage(t *testing.T) {
	TimerUsage()
}

func TestTimerChannelWait(t *testing.T) {
	WaitChan()
}

func TestDelayFunction(t *testing.T) {
	DelayFunction()
}
