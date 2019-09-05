package leetcode

import (
	"testing"
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	var resultNode *ListNode
	var currentNode *ListNode

	t1, t2, leftVal := l1, l2, 0
	for t1 != nil || t2 != nil {
		v1, v2 := 0, 0
		if t1 != nil {
			v1 = t1.Val
			t1 = t1.Next
		}
		if t2 != nil {
			v2 = t2.Val
			t2 = t2.Next
		}

		currentVal := v1 + v2 + leftVal
		if currentVal >= 10 {
			currentVal -= 10
			leftVal = 1
		} else {
			leftVal = 0
		}

		tmpNode := new(ListNode)
		tmpNode.Val = currentVal

		if resultNode == nil {
			resultNode = tmpNode
			currentNode = resultNode
		} else {
			currentNode.Next = tmpNode
			currentNode = currentNode.Next
		}

	}
	if leftVal == 1 {
		lastNode := new(ListNode)
		lastNode.Val = 1
		currentNode.Next = lastNode
	}

	return resultNode
}

func TestAddTwoNumbers(t *testing.T) {
	l1, l2 := ToListNode(50), ToListNode(50)
	s1, s2 := ToSlice(l1), ToSlice(l2)
	fmt.Println(s1)
	fmt.Println(s2)

	r := addTwoNumbers(l1, l2)
	sr := ToSlice(r)
	fmt.Println(sr)

}

func TestToListNode(t *testing.T) {
	val := 239845
	listNode := ToListNode(val)
	arr := ToSlice(listNode)
	fmt.Println(arr)
}

func ToListNode(target int) *ListNode {

	var resultNode *ListNode
	var currentNode *ListNode

	for target > 0 {
		if resultNode == nil {
			resultNode = new(ListNode)
			currentNode = resultNode
		} else {
			currentNode.Next = new(ListNode)
			currentNode = currentNode.Next
		}
		currentNode.Val = target % 10
		target /= 10
	}
	return resultNode
}

func ToSlice(listNode *ListNode) []int {
	result := make([]int, 0)
	for listNode != nil {
		result = append(result, listNode.Val)
		listNode = listNode.Next
	}
	return result
}
