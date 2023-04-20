package exp_use

import "testing"

func TestUseError(t *testing.T) {
	UseError()
}

func TestUsePanic(t *testing.T) {
	UsePanic()
}

func TestRecoverPanic(t *testing.T) {
	RecoverPanic()
}
