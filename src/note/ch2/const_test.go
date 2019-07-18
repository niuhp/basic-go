package ch2

import "testing"

const (
	State1 = iota + 1
	State2
	State3 = iota * 5
	State4
)

func TestConst(t *testing.T) {
	t.Log("State1 =", State1, ",State2 =", State2, ",State3 =", State3, ",State4 =", State4)
}
