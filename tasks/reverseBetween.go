package tasks

func ReverseBetween(head *ListNode, left int, right int) *ListNode {

	cur := head
	var prev *ListNode
	var next *ListNode

	// пропускаем элементы до left
	for i := 1; i < left; i++ {
		prev = cur
		cur = cur.Next
	}

	lastNodeBeforeReverse := prev
	lastNodeAfterReverse := cur

	for i := 0; i < right-left+1; i++ {
		next = cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}

	if lastNodeBeforeReverse != nil {
		lastNodeBeforeReverse.Next = prev
	} else {
		head = prev
	}
	lastNodeAfterReverse.Next = cur
	return head

}
