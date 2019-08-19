package ch09

import (
	"math/rand"
	"testing"
	"time"
	"fmt"
)

func returnMultiValues() (int, int) {
	return rand.Intn(10), rand.Intn(20)
}

func TestReturnMultiValues(t *testing.T) {
	a, b := returnMultiValues()
	t.Log(a, b)
}

func timeSpent(f1 func(op int) int) func(op int) int { // 函数参数
	return func(op int) int {
		s := time.Now()
		ret := f1(op)
		cost := time.Since(s).Seconds()
		fmt.Println("cost:", cost)
		return ret
	}
}

func showFunc(op int) int {
	time.Sleep(time.Second * 2)
	return op + 1
}
func TestTimeSpent(t *testing.T) {
	tsFn := timeSpent(showFunc)
	a := tsFn(10)
	t.Log(a)
}

func sum(values ...int) int { //可变参数
	ret := 0
	for _, value := range values {
		ret += value
	}
	return ret
}
func TestSum(t *testing.T) {
	s1 := sum(1, 3, 5)
	s2 := sum(1, 3, 5, 7)
	t.Log(s1, s2)
}

func clear() {
	fmt.Println("clear resource")
}
func TestDefer(t *testing.T) {
	defer clear()
	fmt.Println("start..")
	panic("error")
}
