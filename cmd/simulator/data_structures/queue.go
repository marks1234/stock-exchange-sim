package datastructures

import (
	"log"
	"stock_exchange_sim/cmd/simulator/components"
)

type StockDetails = components.StockDetails

type Queue struct {
	store []StockDetails
}

func NewQueue() Queue {
	return Queue{}
}

func (queue *Queue) Push(item StockDetails) {
	queue.store = append(queue.store, item)
}

func (queue *Queue) Pull() (item StockDetails) {
	if len(queue.store) > 0 {
		item = queue.store[0]
		queue.store = queue.store[1:]
	} else {
		log.Fatal("Tried to pull from an empty queue")
	}
	return
}

func (queue *Queue) peek() (item StockDetails) {
	if len(queue.store) > 0 {
		item = queue.store[0]
	} else {
		log.Fatal("Tried to peek at empty queue")
	}
	return
}

func (queue *Queue) Len() int {
	return len(queue.store)
}