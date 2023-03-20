package pack1

type matrix struct {
	square int
	width  string
}

// 只导出了工厂方法， matrix 结构体是包私有变量
func NewMatrix() *matrix {
	return new(matrix)
}

func Mul(a int, b int) int {
	return a * b
}
