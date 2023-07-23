package core_use

import (
	"fmt"
	"testing"
)

func TestNil(t *testing.T) {
	err := Process()
	fmt.Println(err)

	fmt.Println(err == nil)
}

func TestPrintType(t *testing.T) {
	printTypeAndV()
}
