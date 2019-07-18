package ch2

import "testing"

/**
连续赋值，iota 从 0 开始，每出现一次加 1
 */
const (
	State1 = iota + 1 //iota=0
	State2            //相当于 iota + 1,iota=1
	State3 = iota * 5 //iota=2
	State4            //相当于 iota * 5,iota=3
)

func TestConst(t *testing.T) {
	t.Log("State1 =", State1, ",State2 =", State2, ",State3 =", State3, ",State4 =", State4)
}
