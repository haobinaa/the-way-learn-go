package core_use

import "fmt"

// 非命名返回值
func getX2AndX3(input int) (int, int) {
	return 2 * input, 3 * input
}

// 命名返回值(返回类型里面申明了返回的变量名， 此刻已经初始化好了零值)
func getX2AndX3_2(input int) (x2 int, x3 int) {
	x2 = 2 * input
	x3 = 3 * input
	// return x2, x3
	return
}

// defer 关键字将语句推迟到函数返回(这个返回并非是指的执行 return 后的语句)之前在执行(有点像 java 的 finally， 一般用于关键资源的释放)
func function1() {
	fmt.Printf("In function1 at the top\n")
	// 最后执行
	defer function2()
	fmt.Printf("In function1 at the bottom!\n")
}

func function2() {
	fmt.Printf("Function2: Deferred until the end of the calling function!")
}

// 多个 defer 注册， 逆序执行(类似栈， 先进后出)
func fDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d ", i)
	}
}

func Add(a, b int) {
	fmt.Printf("The sum of %d and %d is: %d\n", a, b, a+b)
}

// 函数作为参数， 在其他函数内执行(回调)
func callback(y int, f func(int, int)) {
	f(y, 2) // this becomes Add(1, 2)
}

func F() {
	for i := 0; i < 4; i++ {
		// 匿名函数， 需要将函数赋值给变量
		g := func(i int) { fmt.Printf("%d ", i) }
		g(i)
		//fmt.Printf(" - g is of type %T and has value %v\n", g, g)
	}
}

// 闭包与 defer
func fClosure() (ret int) {
	fmt.Println(ret)
	// defer 会改写命名参数的返回值
	defer func() {
		ret++
	}()
	return 1
}

// 将闭包当做函数的返回值(返回一个函数)
func Add2() func(b int) int {
	return func(b int) int {
		return b + 2
	}
}

func Adder() func(int) int {
	var x int
	return func(delta int) int {
		x += delta
		return x
	}
}

// 输出: 1 - 21 - 321
// f 指向闭包函数 Adder， f 里面定义的变量 x 在多次调用中值是累积的
func closureMultiCall() {
	var f = Adder()
	fmt.Print(f(1), " - ")
	fmt.Print(f(20), " - ")
	fmt.Print(f(300))
}

// 闭包实现斐波那契
func fib_closure() func() int {
	a, b := 1, 1
	return func() int {
		temp := a
		a, b = b, a+b
		return temp
	}
}

func PrintFib(n int) {
	f := fib_closure()
	for i := 0; i < n; i++ {
		fmt.Println(f())
	}
}
