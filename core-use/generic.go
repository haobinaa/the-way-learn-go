package core_use

import "fmt"

// Ordered 类型集， 用来约束泛型
type Ordered interface {
	int | int32 | int64 | float32 | float64
}

// Min 用类型集Ordered 来约束泛型函数Min
func Min[T Ordered](x, y T) T {
	if x < y {
		return x
	}
	return y
}

// GenericInference 泛型类型推导
func GenericInference() {
	var a, b, m float64

	m = Min[float64](a, b) // 完整书写
	m = Min(a, b)          // 不传递类型参数
	fmt.Println(m)
}
