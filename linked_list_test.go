package linkedlist

import "testing"

func TestLinkedList_HasLoop(t *testing.T) {
	t.Run("without loop", func(t *testing.T) {
		n1 := Node{next: nil}
		n2 := Node{next: &n1}
		n3 := Node{next: &n2}
		ll := LinkedList{head: &n3}

		res := ll.HasLoop()

		if res != false {
			t.Fatalf("ll.HasLoop() equal %t, but want %t", res, false)
		}
	})
	t.Run("with loop at start", func(t *testing.T) {
		n1 := Node{next: nil}
		n2 := Node{next: &n1}
		n3 := Node{next: &n2}
		n1.next = &n3
		ll := LinkedList{head: &n3}

		res := ll.HasLoop()

		if res != true {
			t.Fatalf("ll.HasLoop() equal %t, but want %t", res, true)
		}
	})
	t.Run("with loop not at start", func(t *testing.T) {
		n1 := Node{next: nil}
		n2 := Node{next: &n1}
		n3 := Node{next: &n2}
		n1.next = &n2
		ll := LinkedList{head: &n3}

		res := ll.HasLoop()

		if res != true {
			t.Fatalf("ll.HasLoop() equal %t, but want %t", res, true)
		}
	})
}

func TestLinkedList_RemoveLoop(t *testing.T) {
	n1 := Node{next: nil}
	n2 := Node{next: &n1}
	n3 := Node{next: &n2}
	n1.next = &n2
	ll := LinkedList{head: &n3}

	ll.RemoveLoop()

	if ll.HasLoop() != false {
		t.Fatal("Linked list has loop")
	}
}