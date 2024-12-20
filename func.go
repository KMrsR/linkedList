package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
}

// После последнего узла
func (list *LinkedList) insertAtBack(data int) {
	newNode := &Node{data: data, next: nil}

	if list.head == nil {
		list.head = newNode
		return
	}

	current := list.head
	for current.next != nil {
		current = current.next
	}

	current.next = newNode
}

// Перед первым узлом
func (list *LinkedList) insertAtFront(data int) {
	if list.head == nil {
		newNode := &Node{data: data, next: nil}
		list.head = newNode
		return
	}

	newNode := &Node{data: data, next: list.head}
	list.head = newNode
}

// После значения узла/=
func (list *LinkedList) insertAfterValue(afterValue, data int) {
	newNode := &Node{data: data, next: nil}

	current := list.head

	for current != nil {
		if current.data == afterValue {
			newNode.next = current.next
			current.next = newNode
			return
		}
		current = current.next
	}

	fmt.Printf("Cannot insert node, after value %d doesn't exist", afterValue)
	fmt.Println()
}

// Перед значением узла
func (list *LinkedList) insertBeforeValue(beforeValue, data int) {
	if list.head == nil {
		return
	}

	if list.head.data == beforeValue {
		newNode := &Node{data: data, next: list.head}
		list.head = newNode
		return
	}

	current := list.head
	for current.next != nil {
		if current.next.data == beforeValue {
			newNode := &Node{data: data, next: current.next}
			current.next = newNode
			return
		}
		current = current.next
	}
	fmt.Printf("Cannot insert node, before value %d doesn't exist", beforeValue)
	fmt.Println()
}

// Удаление первого узла
func (list *LinkedList) deleteFront() {
	if list.head != nil {
		list.head = list.head.next
		fmt.Printf("Head node of linked list has been deleted\n")
		return
	}
}

// Удаление последнего узла
func (list *LinkedList) deleteLast() {
	if list.head == nil {
		fmt.Printf("Linked list is empty\n")
	}

	if list.head.next == nil {
		list.head = nil
		fmt.Printf("Last node of linked list has been deleted\n")
		return
	}

	var current *Node = list.head
	for current.next.next != nil {
		current = current.next
	}
	current.next = nil

	fmt.Printf("Last node of linked list has been deleted")
}

// Удаление дубликатов
func (list *LinkedList) deleteDuplicate() {
	var current *Node = list.head
	if current == nil {
		return
	}
	for current.next != nil {
		if current.data == current.next.data {
			fmt.Printf("found duplicate at %d\n", current.data)
			current.next = current.next.next
			continue
		}
		current = current.next
	}
}

// Удаление после значения узла
func (list *LinkedList) deleteAfterValue(afterValue int) {
	var current *Node = list.head

	for current != nil && current.data != afterValue {
		current = current.next
	}

	if current == nil {
		fmt.Printf("Node with after value %d doesn't exist\n", afterValue)
		return
	}

	if current.next == nil {
		fmt.Printf("Node with after value %d is the last node\n", afterValue)
		return
	}

	var temp *Node = current.next
	fmt.Printf("Node after value %d has data %d and will be deleted", afterValue, temp.data)

	current.next = current.next.next
}

// Удаление перед значением узла
func (list *LinkedList) deleteBeforeValue(beforeValue int) {
	var previous *Node
	current := list.head

	if current == nil || current.next == nil {
		fmt.Printf("Node with before value %d doesn't exist\n", beforeValue)
		return
	}

	for current.next != nil {
		if current.next.data == beforeValue {
			if previous == nil {
				fmt.Printf("Node before value %d has data %d and will be deleted\n", beforeValue, current.data)
				list.head = current.next
			} else {
				fmt.Printf("Node before value %d has data %d and will be deleted\n", beforeValue, current.data)
				previous.next = current.next
			}
			return
		}
		previous = current
		current = current.next
	}
	fmt.Printf("Node before value %d not found\n", beforeValue)
}

// Подсчет общего числа узлов
func (list *LinkedList) countNodes() (count int) {
	current := list.head
	for current != nil {
		current = current.next
		count++
	}
	return
}

// Нахождение узла по заданному индексу
func (list *LinkedList) findNodeAt(index int) *Node {
	var count int = 0
	var current *Node = list.head

	for current != nil {
		count++
		current = current.next
	}

	if index <= 0 || index > count {
		return nil
	}

	current = list.head
	for count = 1; count < index; count++ {
		current = current.next
	}
	return current
}

// Вывод связного списка
func (list *LinkedList) print() {
	var current *Node = list.head
	for current != nil {
		fmt.Printf("%d -> ", current.data)
		current = current.next
	}
	fmt.Println()
}
