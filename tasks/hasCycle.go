package tasks

func hasCycle(head *ListNode) bool {
	fast, slow := head, head
	for head != nil && head.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			return true
		}
	}
	return false
}
