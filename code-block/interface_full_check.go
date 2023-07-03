package code_block

import "fmt"

type Shape interface {
	Sides() int
	Area() int
}
type Square struct {
	len int
}

func (s *Square) Sides() int {
	return 4
}

// InterfaceFullCheck 接口完整性检测
func InterfaceFullCheck() {
	s := Square{len: 5}
	fmt.Printf("%d\n", s.Sides())
	// 检测 Square 是否实现了 Shape 接口所有的方法(编译阶段强检测)
	// 另动态类型断言语法: m, ok := val.(T);ok {}
	//var _ Shape = (*Square)(nil)

}
