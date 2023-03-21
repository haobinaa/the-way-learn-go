package core_use

import "fmt"

//按照约定，只包含一个方法的）接口的名字由方法名加 er 后缀组成，例如 Printer、Reader、Writer、Logger、Converter 等等。还有一些不常用的方式（当后缀 er 不合适时），比如 Recoverable，此时接口名以 able 结尾，或者以 I 开头（像 .NET 或 Java 中那样）。

// 定义接口
type Shaper interface {
	Area() float32
}

type Square struct {
	side float32
}

// 结构体 Square 实现接口 Shaper
func (sq *Square) Area() float32 {
	return sq.side * sq.side
}

// 使用接口变量
func UseInterface() {
	sq1 := new(Square)
	sq1.side = 5
	// 给接口变量赋值 1
	var areaIntf Shaper
	// 结构体赋值给接口变量
	// 还可以直接 areaIntf := Shaper(sq1)
	// 或者 areaIntf := sq
	areaIntf = sq1
	fmt.Printf("The square has area: %f\n", areaIntf.Area())
}

type stockPosition struct {
	ticker     string
	sharePrice float32
	count      float32
}

func (s stockPosition) getValue() float32 {
	return s.sharePrice * s.count
}

type car struct {
	make  string
	model string
	price float32
}

func (c car) getValue() float32 {
	return c.price
}

// 定义接口， 不同的 struct 可以有 getValue 方法
type valuable interface {
	getValue() float32
}

func showValue(asset valuable) {
	fmt.Printf("Value of the asset is %f\n", asset.getValue())
}

// 多态
func MultiValuable() {
	var o valuable = stockPosition{"GOOG", 577.20, 4}
	showValue(o)
	o = car{"BMW", "M3", 66500}
	showValue(o)
}

// TypeInterface 运行时判断接口变量的类型, 语法如下:
// if v, ok := varI.(T); ok {  // checked type assertion
// Process(v)
// return
// }
func TypeInterface() {
	var o valuable = &stockPosition{"GOOG", 577.20, 4}
	// 判断接口类型是 *stockPosition
	if t, ok := o.(*stockPosition); ok {
		fmt.Printf("The type of valuable is: %T\n", t)
	} else {
		fmt.Println("areaIntf does not contain a variable of type Circle")
	}
	o = car{"BMW", "M3", 66500}
	// 这里接口类型是 car 值类型， 而不是指针
	if u, ok := o.(*car); ok {
		fmt.Printf("The type of areaIntf is: %T\n", u)
	} else {
		fmt.Println("areaIntf does not contain a variable of type car")
	}
}

// 所有类型都实现了空接口, 类似 java 里面的 Object
type Any interface {
}

type People struct {
	name string
	age  int
}

func EmptyInterface() {
	var val Any
	val = 5
	fmt.Printf("val has the value: %v\n", val)
	val = str
	fmt.Printf("val has the value: %v\n", val)
	pers1 := new(People)
	pers1.name = "Rob Pike"
	pers1.age = 55
	val = pers1
	fmt.Printf("val has the value: %v\n", val)
	switch t := val.(type) {
	case int:
		fmt.Printf("Type int %T\n", t)
	case string:
		fmt.Printf("Type string %T\n", t)
	case bool:
		fmt.Printf("Type boolean %T\n", t)
	case *Person:
		fmt.Printf("Type pointer to Person %T\n", t)
	default:
		fmt.Printf("Unexpected type %T", t)
	}
}
