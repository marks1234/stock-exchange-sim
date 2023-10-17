package main

import (
	"fmt"
	"os"
	"strings"

	"stock_exchange_sim/cmd/simulator/components"
)

type Node struct {
	name         string
	dependencies map[string]int
	time         int
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
	// eligibile_tasks := tasksEligible(state)
	optimizations := state.OptimizedList()

	for stock, details := range state.Stocks {
		fmt.Println(stock, details.Amount)
		for task, details := range details.Consumers {
			fmt.Print("Consumer: ")
			fmt.Println(task, details)
		}
		for task, details := range details.Producers {
			fmt.Print("Producer: ")
			fmt.Println(task, details)
		}
	}
	// fmt.Println(eligibile_tasks)
	fmt.Println(optimizations)
	// fmt.Println(state.IsTime())

	dot := GenerateDOT(state)
	os.WriteFile("output.dot", []byte(dot), 0o644)
}

func GenerateDOT(s components.State) string {
	var output strings.Builder
	output.WriteString("digraph G {\n")
	output.WriteString("\tnode [shape=circle];\n")
	for stock, details := range s.Stocks {
		output.WriteString(fmt.Sprintf("\t\"%s %s\";\n", stock, fmt.Sprint(details.Amount)))
	}
	for stock, stockDetail := range s.Stocks {
		for taskName, taskDetail := range stockDetail.Producers {
			for consumedStock := range taskDetail.StockNeeded {
				output.WriteString(fmt.Sprintf("\t\"%s %s\" -> \"%s %s\" [label=\"%s\"];\n", consumedStock, fmt.Sprint(s.Stocks[consumedStock].Amount), stock, fmt.Sprint(s.Stocks[stock].Amount), taskName))
			}
		}
	}
	output.WriteString("}")
	return output.String()
}

func createDependencyGraph(state components.State) {
}
