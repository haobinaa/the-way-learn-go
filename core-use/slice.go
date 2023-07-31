package core_use

import (
	"bytes"
	"fmt"
)

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

func ShareMemory() {
	path := []byte("AAAA/BBBBBBBBB")
	sepIndex := bytes.IndexByte(path, '/')
	//dir1 := path[:sepIndex]
	// FullSliceExpression   y := x[start, end, capacity]
	// 限制数组的右界， 防止 append 内存共享
	dir1 := path[:sepIndex:sepIndex]
	dir2 := path[sepIndex+1:]
	fmt.Println("dir1 =>", string(dir1)) //prints: dir1 => AAAA
	fmt.Println("dir2 =>", string(dir2)) //prints: dir2 => BBBBBBBBB
	dir1 = append(dir1, "suffix"...)
	fmt.Println("dir1 =>", string(dir1)) //prints: dir1 => AAAAsuffix
	fmt.Println("dir2 =>", string(dir2)) //prints: dir2 => uffixBBBB
}

func copyInt(a, b []int) {
	// 方式一: 直接使用 copy
	//copy(b, a)
	// 方式二: append
	b = append(a[:0:0], a...)
}

func deleteInt(a []int, index int) {
	// 方式一: append， 也可以扩展至删除一段区间的元素
	a = append(a[:index], a[index:]...)
	// 方式二: 截掉元素后不能保留原顺序
	a[index] = a[len(a)-1] // 将最后一个元素移到索引i处
	a = a[:len(a)-1]       // 截掉最后一个元素
}

// 过滤
func filter(a []int, keep func(x int) bool) {
	n := 0
	for _, x := range a {
		if keep(x) {
			a[n] = x // 保留该元素
			n++
		}
	}
	a = a[:n]
}

// 滑动窗口
func slidingWindow(size int, input []int) [][]int {
	// 返回入参的切片作为第一个元素
	if len(input) <= size {
		return [][]int{input}
	}

	// 以所需的精确大小分配切片
	r := make([][]int, 0, len(input)-size+1)

	for i, j := 0, size; j <= len(input); i, j = i+1, j+1 {
		r = append(r, input[i:j])
	}

	return r
}
