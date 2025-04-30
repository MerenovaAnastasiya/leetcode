package tasks

func reverseList(head *ListNode) *ListNode {
	current := head
	var prev *ListNode
	var next *ListNode

	for current != nil {
		// записываем в переменную next ссылку на следующий элемент
		next = current.Next
		// у текущего элемента меняем ссылку на следующий(это будет null для первого элемента)
		current.Next = prev
		// теперь предыдущий элемент это текущий(для следующей ноды он будет NEXT)
		prev = current
		// а текущий элемент это предыдущая ссылка на следующий(чтобы пройтись по всем элементам)
		current = next
	}
	return prev

}
