package core_use

import "fmt"

type person interface {
	get() string
	set(value string)
}

type Man struct {
	name string
}

func (m Man) get() string {
	return m.name
}

func (m *Man) set(value string) {
	m.name = value
}

func ValueAndPointers() {
	// 值类型
	vMan := Man{name: "zhangsan"}
	// 值类型 调用接收者也是值类型的方法
	fmt.Println(vMan.get())
	// 值类型 调用接收者是指针类型的方法
	vMan.set("docker")

	// 指针类型
	pMan := &Man{name: "lisi"}
	// 指针类型 调用接收者是值类型的方法
	fmt.Println(pMan.get())
	// 指针类型 调用接收者也是指针类型的方法
	pMan.set("worker")
}

func ValueAndPointersDiff() {
	var docker person = &Man{name: "zhangsan"}
	//var docker person = Man{name: "zhangsan"}
	docker.get()
	docker.set("docker")
	fmt.Println(docker.get())
}
