package core_use

import (
	"fmt"
	"reflect"
)

type Foo struct{}

func (f Foo) Hello() string {
	return fmt.Sprintf("hello world")
}

type S struct {
	a int
	b float64
}

func ReflectionUse() {
	f := &Foo{}
	fmt.Println(f.Hello())
	// 反射
	foo := &Foo{}
	ft := reflect.TypeOf(foo)
	fv := reflect.New(ft)
	m := fv.MethodByName("Hello")
	if m.IsValid() {
		fmt.Println(m.Call(nil))
	}
}
