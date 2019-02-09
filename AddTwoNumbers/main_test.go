package main

import (
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	l1 := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 3,
			},
		},
	}

	l2 := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val: 4,
			},
		},
	}

	resNode := addTwoNumbers(l1, l2)
	var res int
	expected := 807

	for i := 1; resNode != nil; i *= 10 {
		res += i * resNode.Val
		resNode = resNode.Next
	}

	if res != expected {
		t.Errorf("Expected %d got %d", expected, res)
	}
}
