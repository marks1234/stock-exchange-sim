package datastructures

import "log"

type Stack struct {
	store []StockDetails
}

func NewStack() Stack {
	return Stack{}
}

func (stack *Stack) Push(item StockDetails) {
	stack.store = append(stack.store, item)
}

func (stack *Stack) Pop() (item StockDetails) {
	store := stack.store
	if len(store) > 0 {
		item = store[len(store)-1]
		stack.store = store[:len(store)-1]
	} else {
		log.Fatal("Tried to pop from an empty stack")
	}
	return
}

func (stack *Stack) Peek() (item StockDetails) {
	store := stack.store
	if len(store) > 0 {
		item = store[len(store)-1]
	} else {
		log.Fatal("Tried to peek at empty stack")
	}
	return
}

func (stack *Stack) Len() int {
	return len(stack.store)
}
