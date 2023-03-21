package exp_use

import (
	"errors"
	"fmt"
)

// err := errors.New("msg") 来定义错误类型
var errorNotFound error = errors.New("Not found error")

func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, errors.New("math - square root of negative number")
	}
	// implementation of Sqrt
	return f, nil
}

func UseError() {
	if f, err := Sqrt(-1); err != nil {
		fmt.Printf("Error: %s\n", err)
	} else {
		fmt.Printf("sqrt result is %f", f)
	}
}

// 当错误不可恢复， 程序不能运行的时候， 可以使用 panic 函数产生中止程序错误的运行时错误
func UsePanic() {
	fmt.Println("Starting the program")
	panic("A severe error occurred: stopping the program!")
	fmt.Println("Ending the program")
}
