package pack1

import "fmt"

type struct1 struct {
	i1  int
	f1  float32
	str string
}

func structUse() {
	// 等同于 var ms *struct1 = new struct1()
	ms := new(struct1)
	ms.i1 = 10
	ms.f1 = 15.5
	ms.str = "Chris"

	fmt.Printf("The int is: %d\n", ms.i1)
	fmt.Printf("The float is: %f\n", ms.f1)
	fmt.Printf("The string is: %s\n", ms.str)
	// 直接打印结构体内容， 类似于 %v 格式化
	fmt.Println(ms)
}

type Person struct {
	firstName string
	lastName  string
}

// 结构体初始化的方式
func structAlloc() {
	// 定义变量的方式
	var pers1 Person
	pers1.firstName = "Chris"
	pers1.lastName = "Woodward"
	fmt.Printf("The name of the person is %s %s\n", pers1.firstName, pers1.lastName)

	// 使用 new 关键字
	pers2 := new(Person)
	pers2.firstName = "Chris"
	pers2.lastName = "Woodward"
	(*pers2).lastName = "Woodward" // 这是合法的
	fmt.Printf("The name of the person is %s %s\n", pers2.firstName, pers2.lastName)

	// 结构体字面量 &
	pers3 := &Person{"Chris", "Woodward"}
	fmt.Printf("The name of the person is %s %s\n", pers3.firstName, pers3.lastName)
}
