package main

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	return calc(l1, l2)
}

func calc(l1 *ListNode, l2 *ListNode) *ListNode {
	var res *ListNode

	extra := 0
	res = &ListNode{}
	node := res
	for {
		n := l1.Val + l2.Val + extra

		if n < 10 {
			node.Val = n
			extra = 0
		} else {
			node.Val = n % 10
			extra = n / 10
		}

		if l1.Next == nil && l2.Next == nil {
			if extra > 0 {
				node.Next = &ListNode{Val: extra}
			}
			return res
		}

		node.Next = &ListNode{}
		node = node.Next

		if l1.Next == nil {
			l1.Val = 0
		} else {
			l1 = l1.Next
		}

		if l2.Next == nil {
			l2.Val = 0
		} else {
			l2 = l2.Next
		}
	}
}

func main() {
	//[1]
	//[9,9]
	//
	//fmt.Println(addTwoNumbers());
}
