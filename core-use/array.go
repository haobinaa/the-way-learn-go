package core_use

import "fmt"

// 类型是 *[]int, 可以直接修改原数组
var arr2 = new([5]int)

// 类型是 []int,
var arr3 [5]int

func ArrSimple() {
	var arr1 [5]int
	for i := 0; i < len(arr1); i++ {
		arr1[i] = i * 2
	}
	// for 循环处理
	for i := 0; i < len(arr1); i++ {
		fmt.Println(arr1[i])
	}
	// for range 处理, 语法: for ix, value := range slice1
	for i, _ := range arr1 {
		fmt.Printf("%d \t", arr1[i])
	}
	fmt.Printf("\n")
	// 地址拷贝
	arr3 = *arr2
	arr3[0] = 444
	fmt.Println(arr2[0])
}

func arrConstInit() {
	// 常量初始化
	var arrAge = [5]int{1, 2, 3, 4, 5}
	// 切片
	var arrLazy = [...]int{1, 2, 3}
	// 指定索引赋值
	var arrKeyValue = []string{3: "hello", 4: "world"}
	fmt.Println(arrAge, arrLazy, arrKeyValue)
}

func sliceGrow() {
	slFrom := []int{1, 2, 3}
	slTo := make([]int, 10)
	// 将原来的 slice 复制到一个更大的里面
	n := copy(slTo, slFrom)
	fmt.Println(slTo)
	fmt.Printf("Copied %d elements\n", n) // n == 3

	sl3 := []int{1, 2, 3}
	// 追加元素并返回新的 slice
	sl3 = append(sl3, 4, 5, 6)
}
