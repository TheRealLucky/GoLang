package main

import "fmt"

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


func main(){
	fmt.Println("Testing...")
} 

