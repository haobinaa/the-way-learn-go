package core_use

import (
	"fmt"
	"strconv"
)

type twoInts struct {
	a int
	b int
}

// String
// 定制 String 方法来格式化 %v 产生的输出
func (tn *twoInts) String() string {
	return "(" + strconv.Itoa(tn.a) + "/" + strconv.Itoa(tn.b) + ")"
}

// NewTwoInts 只导出了工厂方法， matrix 结构体是包私有变量
func NewTwoInts() *twoInts {
	return new(twoInts)
}

func Mul(a int, b int) int {
	return a * b
}

// AddThem
// 方法定义: func (_ receiver_type) methodName(parameter_list) (return_value_list) { ... }
// 结构体 twoInts 作为 receiver 来定义方法
func (ti *twoInts) AddThem() int {
	return ti.a + ti.b
}

// AddThenValue
// 值传递给方法， 传递到方法的是拷贝
func (ti twoInts) AddThenValue() int {
	return ti.a + ti.b
}

type B struct {
	thing int
}

func (b *B) change() { b.thing = 1 }

func (b B) write() string { return fmt.Sprint(b) }
