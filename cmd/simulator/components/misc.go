package components

import "fmt"

func DisplayStockDetails(state State) {
	for stock, details := range state.Stocks {
		fmt.Printf("Stock: %s (Amount: %d)\n", stock, details.Amount)

		if len(details.Consumers) > 0 {
			fmt.Println("  Consumers:")
			for task, details := range details.Consumers {
				fmt.Printf("    Task: %s (Delay: %d)\n", task, details.Delay)
				for neededStock, amount := range details.StockNeeded {
					fmt.Printf("      Requires: %s (Amount: %d)\n", neededStock, amount)
				}
				for resultStock, amount := range details.StockResults {
					fmt.Printf("      Produces: %s (Amount: %d)\n", resultStock, amount)
				}
			}
		}

		if len(details.Producers) > 0 {
			fmt.Println("  Producers:")
			for task, details := range details.Producers {
				fmt.Printf("    Task: %s (Delay: %d)\n", task, details.Delay)
				for neededStock, amount := range details.StockNeeded {
					fmt.Printf("      Requires: %s (Amount: %d)\n", neededStock, amount)
				}
				for resultStock, amount := range details.StockResults {
					fmt.Printf("      Produces: %s (Amount: %d)\n", resultStock, amount)
				}
			}
		}
		fmt.Println("----------------------------")
	}
}
