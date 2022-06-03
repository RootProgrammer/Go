package main

import (
	"fmt"
	"strings"
)

type node struct {
	value int
	next  *node
}

type linkedList struct {
	head   *node
	length int
}

func (l *linkedList) add(value int) {
	newNode := new(node)
	newNode.value = value

	if l.head == nil {
		l.head = newNode
	} else {
		iterator := l.head

		//  getting to the end of a linked list, and then adding a node to it.
		/*
			NB: the way it's written is not idiomatic for Go,
			to put all of that specific logic in the loop header.
		*/
		for ; iterator.next != nil; iterator = iterator.next {
		}
		iterator.next = newNode
	}
}

func (l *linkedList) remove(value int) {

}

func (l linkedList) get(value int) *node {
	return nil
}

func (l linkedList) String() string {
	string_builder := strings.Builder{}

	for iterator := l.head; iterator != nil; iterator = iterator.next {
		string_builder.WriteString(fmt.Sprintf("%d", iterator.value))
	}

	return string_builder.String()
}

func main() {
	fmt.Println("Linked List.")

	linked_list := linkedList{}
	linked_list.add(1)

	fmt.Println(linked_list)
}
