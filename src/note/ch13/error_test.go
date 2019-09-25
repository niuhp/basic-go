package ch13

import (
	"errors"
	"testing"
	"fmt"
)

func Div(a int, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("b can't be 0")
	}
	return a / b, nil
}

func TestDiv(t *testing.T) {

	a, b := 9, 2
	v, err := Div(a, b)
	fmt.Println("v =", v, ",error =", err)

	a, b = 9, 0
	v, err = Div(a, b)
	fmt.Println("v =", v, ",error =", err)

}
