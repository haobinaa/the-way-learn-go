package main

import (
	"fmt"
	"os"
)

// 常量 显示类型
const b string = "abc"

// 隐式常量类型
const c = "C"

// 常量并行赋值， iota 是一个关键字， 默认为 0
// 该方式还可以作为枚举赋值使用， iota 在新的一行被使用时， 它的值会自动 +1
// iota 的使用方式比较复杂， 需要单独在看看
const (
	e = iota
	f = iota
	g = iota
)

// 变量申明, 类型在后
var v int = 5

// 变量申明会自动赋零值(初始值)
var fl float32

// 编译器自动类型推断
var str = "Hello World Go"

// HOME 运行时自动类型推断
var HOME = os.Getenv("HOME")

// 并行赋值
var x, y, z = 2, 3, "Z"

type T struct {
}

// init 函数用于程序开始前执行， 先于 main
func init() {
	fmt.Println("init check")
}
func (t T) method1() {

}

func printConst() {
	fmt.Println(e, f, g)
	fmt.Print("hello:", "world")
	fmt.Print("\n")
}

func localVariable() {
	// 简短赋值， 只能在函数体内， 用于能自动推断类型的赋值
	// 局部变量申明了必须使用
	localA := 1
	fmt.Println(localA)
	fmt.Println(&localA)

}

func main() {
	//fmt.Println("hello world")
	//fmt.Println(v)
	//localVariable()
	//printConst()
	//pack1.Strconvrt()
	//pack1.ErrFunc()
	//pack1.PrintFib(10)
	//pack1.ArrSimple()
	//
	//pack2.ReturnPack2()

	var a *string
	fmt.Println(a)
	fmt.Println(&a)
}
