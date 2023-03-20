package pack1

import "fmt"

// 切片是对数组一个连续片段的引用
// 长度可在运行时修改， 0 <= len(s) <= cap(s), cap(s) 从 s[0] 到数组尾部
// arr1[2:] 和 arr1[2:len(arr1)] 相同，都包含了数组从第三个到最后的所有元素。
// arr1[:3] 和 arr1[0:3] 相同，包含了从第一个到第三个元素（不包括第四个）。
// 数组未定义的时候， 可以通过 make 创建切片，var slice []type = make([]type, len)， 简写 slice := make([]type, len)

func SliceSimple() {
	var arr1 [6]int
	// 切片定义方式 var slice1 []type = arr1[start:end]
	var slice []int = arr1[2:5]
	for i := 0; i < len(arr1); i++ {
		arr1[i] = i
	}
	// 打印切片
	for i := 0; i < len(slice); i++ {
		fmt.Printf("Slice at %d is %d\n", i, slice[i])
	}
	// 扩容切片， 切片最高可以扩容到上限 s = s[:cap(s)]
	slice = slice[0:4]
	for i := 0; i < len(slice); i++ {
		fmt.Printf("Slice at %d is %d\n", i, slice[i])
	}
}

// 切片传递给函数
func sliceFunc() {
	var arr = [5]int{0, 1, 2, 3, 4}
	// 数组分片为切片引用传递给函数
	sum(arr[:])
}

// 入参是一个切片
func sum(a []int) int {
	s := 0
	for i := 0; i < len(a); i++ {
		s += a[i]
	}
	return s
}

func makeSlice() {
	slice := make([]int, 10)
	// 初始化切片和数组
	for i := 0; i < len(slice); i++ {
		slice[i] = 5 * i
	}

	for i := 0; i < len(slice); i++ {
		fmt.Printf("Slice at %d is %d\n", i, slice[i])
	}
}
