package core_use

import (
	"fmt"
	"testing"
	"unsafe"
)

type Num struct {
	i string
	j int64
}

func TestPtrAddr(t *testing.T) {
	n := Num{i: "EDDYCJY", j: 1}

	fmt.Printf("%p\n", &n)
	fmt.Println(&n.i)
	// 将结构体指针转换为 unsafe.Pointer，实际上这里也是在转换 n.i 的指针
	// &n 指针中保留着结构体 Num 的信息，因此它能够直接访问 &n.i 和 &n.j
	// 而转换为 unsafe.Pointer 后，类似于转换成了 interface{}，没有保留 Num 的结构体信息，因此 nPointer 无法直接访问 nPointer.i 和 nPointer.j
	nPointer := unsafe.Pointer(&n)

	// 将 unsafe.Pointer 转换为 i 对应的 string 类型指针，重新赋值
	niPointer := (*string)(nPointer)
	*niPointer = "煎鱼"

	// 加上 n.j 的偏移量，获取 j 的指针，重新赋值
	njPointer := (*int64)(unsafe.Pointer(uintptr(nPointer) + unsafe.Offsetof(n.j)))
	*njPointer = 2

	fmt.Println(n) // {煎鱼 2}
}
