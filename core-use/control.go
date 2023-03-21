package core_use

import (
	"fmt"
	"strconv"
	"strings"
)

// if 流程控制使用
func ifusage() {
	str1 := "hello"
	if strings.HasPrefix(str1, "he") {
		fmt.Println("get string he prefix")
	} else {
		fmt.Println("no hello")
	}
}

// 带 error 返回函数用法
func ErrFunc() {
	orig := "ABC"
	an, err := strconv.Atoi(orig)
	if err != nil {
		fmt.Printf("orig %s is not integer", orig)
		return
	}
	fmt.Printf("the integer is %d", an)
}

func forusage() {
	str := "hello world"
	// 类似 foreach 语法
	for pos, char := range str {
		fmt.Printf("index %d, char: %c", pos, char)
	}
}
