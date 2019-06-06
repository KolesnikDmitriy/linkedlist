package linkedlist

import "testing"

func TestLinkedList_FindAndRemoveLoop(t *testing.T) {
	t.Run("without loop", func(t *testing.T) {
		n1 := &Node{next: nil}
		n2 := &Node{next: n1}
		n3 := &Node{next: n2}
		ll := LinkedList{head: n3}

		ll.FindAndRemoveLoop()

		if n1.next != nil {
			t.Fatal("Linked list has loop")
		}
	})
	t.Run("with loop at start", func(t *testing.T) {
		n1 := &Node{next: nil}
		n2 := &Node{next: n1}
		n3 := &Node{next: n2}
		n1.next = n3
		ll := LinkedList{head: n3}

		ll.FindAndRemoveLoop()

		if n1.next != nil {
			t.Fatal("Linked list has loop")
		}
	})
	t.Run("with loop not at start", func(t *testing.T) {
		n1 := &Node{next: nil}
		n2 := &Node{next: n1}
		n3 := &Node{next: n2}
		n1.next = n2
		ll := LinkedList{head: n3}

		ll.FindAndRemoveLoop()

		if n1.next != nil {
			t.Fatal("Linked list has loop")
		}
	})
}
