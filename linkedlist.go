package linkedlist

type Node struct {
	next *Node
}

type LinkedList struct {
	head *Node
}

func (ll *LinkedList) FindAndRemoveLoop() {
	if ll.head == nil {
		return
	}
	slow := &Node{next: ll.head}
	fast := &Node{next: ll.head}
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
			slow = &Node{next: ll.head}
			for slow.next != fast.next {
				fast = fast.next
				slow = slow.next
			}
			fast.next = nil
		}
	}
}
