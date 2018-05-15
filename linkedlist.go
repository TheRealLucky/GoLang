package linkedlist

import "fmt"

type indexOutOfRange struct {
	message string
}

type linkedlist struct {
	root *node
}

type node struct {
	next *node
	data interface{}
}

func initNode(data interface{}) *node {
	n := new(node)
	n.data = data
	return n
}

func InitList() *linkedlist {
	l := new(linkedlist)
	l.root = nil
	return l
}

func initError(msg string) *indexOutOfRange {
	err := new(indexOutOfRange)
	err.message = msg
	return err
}

func (l *linkedlist) Add(data interface{}) {
	element := initNode(data)
	if l.root == nil {
		l.root = element
	} else {
		curr := l.root
		for {
			if(curr.next == nil) {
				break
			}
			curr = curr.next
		}
		curr.next = element
	}
}

func (l *linkedlist) Get(index int) (data interface{}, err *indexOutOfRange) {
	if l.root == nil {
		return nil, initError("List is empty!")
	}
	if index == 0 {
		return l.root.data, initError("")
	}else {
		curr := l.root
		for i := 0; i < index; i++ {
			if curr.next == nil {
				return nil, initError("Index out of Range!")
			}
			curr = curr.next
		}
		return curr.data, initError("")
	}
}


