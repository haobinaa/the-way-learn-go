package code_block

import "fmt"

type Coder interface {
	code()
}
type Gopher struct {
	name string
}

func (g Gopher) code() {
	fmt.Printf("%s is coding\n", g.name)
}

// ItabAndData 测试接口的动态类型和动态值
// 接口底层实现 iface 由 *itab(接口表指针， 指向类型信息) + data(数据指针， 指向具体数据) 组成
// itab 称为动态类型  data 称为动态值，  接口为 nil 代表这两个部分都为 nil
func ItabAndData() {
	var c Coder
	// true
	fmt.Println(c == nil)
	// c: <nil>, <nil>
	fmt.Printf("c: %T, %v\n", c, c)
	var g *Gopher
	// true
	fmt.Println(g == nil)
	c = g
	// false
	fmt.Println(c == nil)
	// *code_block.Gopher, <nil>
	fmt.Printf("c: %T, %v\n", c, c)
}
