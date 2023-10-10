package main

import (
	"fmt"
	"os"

	"stock_exchange_sim/cmd/simulator/components"
)

type State struct {
	time  int
	stock map[string]int
	tasks map[string]TaskDetails
}

type TaskDetails struct {
	dependecies map[string]int
	results     map[string]int
	delay       int
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
	fmt.Println(state)
}

func tasksEligible(state State) {
}
