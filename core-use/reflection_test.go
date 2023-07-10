package core_use

import (
	"fmt"
	"reflect"
	"testing"
)

func TestReflectionUse(t *testing.T) {
	ReflectionUse()
}

func TestRelectionValueType(t *testing.T) {
	var x float64 = 3.4
	var y int = 100
	var s = S{
		a: 1,
		b: 90.9,
	}
	fmt.Println("type:", reflect.TypeOf(x))
	fmt.Println("value:", reflect.ValueOf(x).String())
	fmt.Println("type:", reflect.TypeOf(y))
	fmt.Println("value:", reflect.ValueOf(y).String())
	fmt.Println("type:", reflect.TypeOf(s))
	fmt.Println("value:", reflect.ValueOf(s).String())
}
