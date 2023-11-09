package code_block

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"runtime/debug"
	"sync"
	"time"
)

// PanicHandler 定义一个处理panic的函数类型
type PanicHandler func(p Panic)

// Panic 记录协程panic信息
type Panic struct {
	R     interface{} // recover 返回信息
	Stack []byte      // 堆栈信息
}

// ProtectGoroutine 协程守护
type ProtectGoroutine struct {
	sync.Mutex
	wg            *sync.WaitGroup
	panics        chan Panic
	panicHandler  PanicHandler
	Ctx           context.Context
	Cancel        context.CancelFunc
	retryInterval time.Duration
}

// NewProtectGoroutine 创建协程执行对象
func NewProtectGoroutine(c context.Context, panicHandler PanicHandler) *ProtectGoroutine {
	ctx, cancel := context.WithCancel(c)
	pg := &ProtectGoroutine{
		wg:            new(sync.WaitGroup),
		panics:        make(chan Panic, 1),
		panicHandler:  panicHandler,
		Ctx:           ctx,
		Cancel:        cancel,
		retryInterval: 3 * time.Second,
	}
	go pg.handlePanic()
	return pg
}

// Protect 协程守护
func (p *ProtectGoroutine) Protect(fn interface{}, params []interface{}) (ret []reflect.Value, err error) {
	c := make(chan struct{}, 1)
	localPanics := make(chan Panic, 1)
	go func() {
	protect:
		for {
			c <- struct{}{}
			go func() {
				defer func() {
					if e := recover(); e != nil {
						<-c
						stack := debug.Stack()
						localPanics <- Panic{
							R:     e,
							Stack: stack,
						}
					}
				}()
				ret, err = p.run(fn, params)
			}()
			select {
			case <-p.Ctx.Done():
				close(c)
				break protect
			case panicInfo := <-localPanics:
				// 当发生panic时，不阻塞select，允许重试逻辑继续执行
				p.panics <- panicInfo
			}
			time.Sleep(p.retryInterval)
		}
	}()
	return
}

// Go 协程启动
func (p *ProtectGoroutine) Go(fn interface{}, params []interface{}) (ret []reflect.Value, err error) {
	p.wg.Add(1)
	done := make(chan struct{})
	go func() {
		defer func() {
			p.wg.Done()
			if e := recover(); e != nil {
				p.panics <- Panic{
					R:     e,
					Stack: debug.Stack(),
				}
			}
			close(done)
		}()
		ret, err = p.run(fn, params)
	}()
	select {
	case <-p.Ctx.Done():
		return nil, p.Ctx.Err()
	case <-done:
		return ret, err
	}
}

// handlePanic 处理panic信息
func (p *ProtectGoroutine) handlePanic() {
	for panicInfo := range p.panics {
		if p.panicHandler != nil {
			p.panicHandler(panicInfo)
		} else {
			fmt.Println(panicInfo.R, string(panicInfo.Stack), time.Now())
		}
	}
}

// run 执行用户任务
func (p *ProtectGoroutine) run(fn interface{}, params []interface{}) (ret []reflect.Value, err error) {
	v := reflect.ValueOf(fn)
	if v.Type().Kind() != reflect.Func {
		return nil, errors.New("params[1] is not callable")
	}
	v.Type().NumIn()
	in := make([]reflect.Value, len(params))
	for k, param := range params {
		in[k] = reflect.ValueOf(param)
	}
	return v.Call(in), nil
}
