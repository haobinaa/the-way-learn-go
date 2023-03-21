package goroutines

import (
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
