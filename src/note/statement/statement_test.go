package statement

import (
	"testing"
	"fmt"
)

func TestFor1(t *testing.T) {
	for i := 1; i <= 20; i += 3 {
		fmt.Println(i)
	}
}
func TestFor2(t *testing.T) {
	i := 5
	for {
		fmt.Println(i)
		i += 5
	}
}
