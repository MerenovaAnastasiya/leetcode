package tasks

func ReverseKGroup(head *ListNode, k int) *ListNode {

	var next *ListNode
	var prev *ListNode
	var currentLast *ListNode
	var newHead *ListNode
	var lastTail *ListNode
	current := head

	for hasKNodes(current, k) {

		// первый элемент группы до разворота(потом он окажется последним в группе)
		currentLast = current
		prev = nil

		for i := 0; i < k && current != nil; i++ {
			next = current.Next
			current.Next = prev
			prev = current
			current = next
		}

		// голова результирующего списка, устанавливается один раз
		if newHead == nil {
			newHead = prev
		}

		if lastTail != nil {
			lastTail.Next = prev
		}

		// хвост предыдущей развернутой группы
		lastTail = currentLast

	}
	if lastTail != nil {
		lastTail.Next = current
	}
	return newHead

}

func hasKNodes(node *ListNode, k int) bool {
	for i := 0; i < k; i++ {
		if node == nil {
			return false
		}
		node = node.Next
	}
	return true
}
