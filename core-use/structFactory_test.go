package core_use

import (
	"fmt"
	"testing"
)

func TestNewTwoInts(t *testing.T) {
	fmt.Println("hello test")
	F()
}

func TestTwoInts_AddThem(t *testing.T) {
	two1 := &twoInts{1, 2}
	fmt.Printf("add result: %d", two1.AddThem())
}

// 测试指针和值作为 receiver
func TestPointerAndValue(t *testing.T) {
	var b1 B // b1 是值
	b1.change()
	fmt.Println(b1.write())

	b2 := new(B) // b2 是指针
	b2.change()
	fmt.Println(b2.write())
}

func TestMultiValuable(t *testing.T) {
	MultiValuable()
}

func TestTypeInterface(t *testing.T) {
	TypeInterface()
}

func TestStructSet(t *testing.T) {
	StructSet()
}

func TestChoose(t *testing.T) {
	for {
		fmt.Println(1)
	}
}
