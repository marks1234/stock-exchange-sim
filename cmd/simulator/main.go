package main

import (
	"fmt"
	"os"
	

	"stock_exchange_sim/cmd/simulator/components"
)

type (
	State        = components.State
	TaskDetails  = components.TaskDetails
	StockDetails = components.StockDetails
)

type Node struct {
	name   string
	weight int
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Not the right amount of arguments")
		fmt.Println(`Expected: go run cmd/simulator/main.go "Txt file here"`)
		return
	}

	file, err := os.Open(os.Args[1])
	check(err)
	defer file.Close()

	state := components.InitState(file)

	pathCreator(state)
	// eligibile_tasks := tasksEligible(state)
	// fmt.Println(eligibile_tasks)
	// fmt.Println(state.IsTime())
	// components.GenerateDOT(state)
}

func hasEmptyProducer(stock StockDetails) bool {
	if len(stock.Producers) <= 0 {
		return true
	}
	return false
}

func pathCreator(state State) {
	optimization := state.Optimized()
	stock := state.Stocks[optimization]
	// finished := false
	// var return_stack [][]StockDetails

	for finished {
		var stack []StockDetails
		var queue []StockDetails
		stack = append(stack, stock)
		
		if hasEmptyProducer(stock) {
			stack
		} 

	}
}


