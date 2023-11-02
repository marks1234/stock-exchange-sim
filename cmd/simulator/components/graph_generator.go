package components

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func GenerateDOT(s State) {
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
	os.WriteFile("output.dot", []byte(output.String()), 0o644)
	cmd := exec.Command("dot", "-Tpng", "output.dot", "-o", "simple.png")

	_, err := cmd.Output()
	check(err)
	// fmt.Println("command Out:", out)
}
