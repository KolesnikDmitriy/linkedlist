package linked_list

type Node struct {
	next    *Node
}

type LinkedList struct {
	head *Node
}

func (ll *LinkedList) HasLoop() bool {
	if ll.head == nil {
		return false
	}
	turtle := ll.head
	hare := ll.head
	for {
		if hare.next == nil {
			return false
		}
		hare = hare.next
		turtle = turtle.next
		if hare.next == nil {
			return false
		}
		hare = hare.next

		if hare == turtle {
			return true
		}
	}
}