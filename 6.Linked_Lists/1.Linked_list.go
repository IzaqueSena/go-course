package main

import "fmt"

type ListNode struct {
	val  int
	next *ListNode
}

// Read
// return ptr to node at position i (0, 1, 2, ..). if length(list) < i =>  return nullptr
// O(n) time. (n=i), O(1) space
func getNode(head *ListNode, i int) *ListNode {
	ptr := head
	for j := 0; j < i && ptr != nil; j++ {
		ptr = ptr.next
	}
	return ptr
}

// Insert
// insert a node at position i, 0<=i<size
// O(n) time (n=i), O(1) space
func insertNodeAt(head *ListNode, i, value int) {
	ptr := head
	for j := 0; j < i-1 && ptr != nil; j++ {
		ptr = ptr.next
	}

	if ptr != nil {
		dummyNode := ptr.next
		newNode := &ListNode{value, dummyNode}
		ptr.next = newNode
	}
}

// Insert given a node which you want to insert a new node after.
// node != nil
// O(1) time O(1) space
func insertNodeGivenNode(previousNode *ListNode, value int) {
	dummyNode := previousNode.next
	newNode := &ListNode{value, dummyNode}
	previousNode.next = newNode
}

// Search
// O(n) time O(1) space
// if key is not in the list return nil
func searchList(head *ListNode, key int) *ListNode {
	var ptr *ListNode
	for ptr = head; ptr != nil && ptr.val != key; ptr = ptr.next {
	}
	return ptr
}

// Update
// O(n) time O(1) space
func updateList(head *ListNode, oldVal, newVal int) {
	var ptr *ListNode
	for ptr = head; ptr != nil && ptr.val != oldVal; ptr = ptr.next {

	}
	if ptr != nil {
		ptr.val = newVal
	}
}

// Delete
// O(n) time O(1) space
func deleteValueFromList(head *ListNode, val int) {
	headRunner := head
	var previousNode *ListNode

	for headRunner != nil && headRunner.val != val {
		previousNode = headRunner
		headRunner = headRunner.next
	}

	if headRunner != nil {
		previousNode.next = headRunner.next
	}
}

// Insert at the end given end node
// O(1) time O(1) space
func insertAtEndGivenEnd(endNode *ListNode, value int) *ListNode {
	newNode := &ListNode{value, nil}
	if endNode != nil {
		endNode.next = newNode
	}
	return newNode
}

// Printing
func printList(head *ListNode) {
	for ptr := head; ptr != nil; ptr = ptr.next {
		fmt.Print(ptr.val, " ")
	}
	fmt.Println()
}

func main() {
	lista := []int{10, 9, 3, 4, 1}
	head := insertAtEndGivenEnd(nil, 10)
	for i, ptr := 1, head; i < len(lista); i++ {
		ptr = insertAtEndGivenEnd(ptr, lista[i])
	}
	printList(head)

	for i := 0; i < len(lista); i++ {
		fmt.Print(getNode(head, i), " ")
	}
	fmt.Println()

	insertNodeAt(head, 3, 55)
	printList(head)

	insertNodeGivenNode(head, 99)
	printList(head)

	fmt.Println(searchList(head, 4).val)

	updateList(head, 4, 8)
	printList(head)

	deleteValueFromList(head, 8)
	printList(head)
}
