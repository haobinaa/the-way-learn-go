package core_use

import "fmt"

// 指针的使用
func Pointer() {
	var intP *int
	var i = 15
	intP = &i
	s := "good bye"
	var strP *string = &s
	// 取指针的值
	fmt.Printf("int pointer value is %d \n", *intP)
	// 指针本身指向一个地址
	fmt.Printf("int pointer  is %p \n", intP)
	fmt.Printf("string pointer is: %s \n", *strP)
	fmt.Printf("string pointer is: %p \n", strP)
}
