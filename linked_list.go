package linkedlist

type Node struct {
	next *Node
}

type LinkedList struct {
	head *Node
}

func (ll *LinkedList) HasLoop() bool {
	if ll.head == nil {
		return false
	}
	slow := ll.head
	fast := ll.head
	for {
		if fast.next == nil {
			return false
		}
		fast = fast.next
		slow = slow.next
		if fast.next == nil {
			return false
		}
		fast = fast.next

		if fast == slow {
			return true
		}
	}
}

func (ll *LinkedList) RemoveLoop() {
	if ll.head == nil {
		return
	}
	slow := ll.head
	fast := ll.head
	for {
		if fast.next == nil {
			return
		}
		fast = fast.next
		slow = slow.next
		if fast.next == nil {
			return
		}
		fast = fast.next

		if fast == slow {
			slow = ll.head
			for slow.next != fast.next {
				fast = fast.next
				slow = slow.next
			}
			fast.next = nil
		}
	}
}
