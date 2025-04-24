package tasks

type ListNode struct {
	Val  int
	Next *ListNode
}

func IsPalindromeLinkedList(head *ListNode) bool {
	secondHead := reverse(head)
	for secondHead != nil && head != nil {
		if secondHead.Val != head.Val {
			return false
		}
	}
	return true
}

func reverse(head *ListNode) *ListNode {
	// a -> b -> -> null
	// a <- b <- c <- null
	current := head
	var prev *ListNode
	for current != nil {
		tmp := current.Next
		current.Next = prev
		prev = current
		current = tmp
	}
	return prev
}
