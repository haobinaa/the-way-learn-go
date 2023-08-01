package core_use

import (
	"fmt"
	"reflect"
	"testing"
)

// TestSimple 反射的基础使用
func TestSimple(t *testing.T) {
	var b = "hello"
	fmt.Printf("type: %v, value: %v", reflect.TypeOf(b), reflect.ValueOf(b))
}

func TestRelectionValueType(t *testing.T) {
	var x float64 = 3.4
	var y int = 100
	var s = S{
		a: 1,
		b: 90.9,
	}
	fmt.Println("type:", reflect.TypeOf(x))
	fmt.Println("value:", reflect.ValueOf(x))
	fmt.Println("type:", reflect.TypeOf(y))
	fmt.Println("value:", reflect.ValueOf(y))
	fmt.Println("type:", reflect.TypeOf(s))
	fmt.Println("value:", reflect.ValueOf(s))
}

func TestFieldCheck(t *testing.T) {
	example := Example{
		A: 0,
		B: "Hello World!",
		C: false,
	}

	fieldsToCheck := []string{"A", "B", "C", "D"}
	emptyFieldsResult := emptyFields(example, fieldsToCheck)

	fmt.Printf("空字段列表：%v\n", emptyFieldsResult)
}
