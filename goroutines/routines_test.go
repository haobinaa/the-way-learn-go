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

	time.Sleep(1e9)
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

func TestTimeTickerUsage(t *testing.T) {
	TimeTickerUsage()
}

// TestBenchmark benchmark
func TestBenchmark(t *testing.T) {
	fmt.Println(" sync", testing.Benchmark(BenchmarkChannelSync).String())
	fmt.Println("buffered", testing.Benchmark(BenchmarkChannelBuffered).String())
}
