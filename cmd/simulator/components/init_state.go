package components

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// State structure defines the current state of a system.
type State struct {
	Time   int                     // Current Time of the system.
	Stocks map[string]StockDetails // Current Stocks/resources in the system.
	goal   map[string]bool
}

type StockDetails struct {
	isInitialized        bool
	Amount               int
	Producers, Consumers map[string]TaskDetails
}

// TaskDetails structure defines properties of a task.
type TaskDetails struct {
	StockNeeded, StockResults map[string]int // Prerequisites for the task.
	Delay                     int            // Time required to execute the task.
}

// Utility function to handle errors.
func check(err error) {
	if err != nil {
		panic(err)
	}
}

//^(\w+):\((\w+:\d+(?:;\w+:\d+)*)\):\((\w+:\d+(?:;\w+:\d+)*)\):(\d+)$

func (s State) IsTime() bool {
	return s.goal["time"]
}

func (s State) OptimizedList() string {
	var results string
	var count int
	for key := range s.goal {
		if key != "time" {
			count++
			results = key
		}
		if count > 1 {
			panic("Too many resources to optimize")
		}
	}
	if count == 0 && !s.IsTime() {
		panic("Optimization not set")
	}
	return results
}

// InitState initializes the state from a provided file.
func InitState(file *os.File) State {
	state := State{
		Time:   0,
		Stocks: make(map[string]StockDetails),
		goal:   make(map[string]bool),
	}
	scanner := bufio.NewScanner(file)
	optimize_found := false

	for scanner.Scan() {
		text := scanner.Text()
		if len(text) != 0 {
			if text[0] == '#' {
				continue
			}
		} else {
			continue
		}
		// Extract resource details from the current line.
		re := regexp.MustCompile(`\b\w+:[\d\.]+\b`)
		resources := re.FindAllString(text, -1)
		if len(resources) == 1 {

			name, details := initStock(resources, true)

			state.Stocks[name] = details
			continue
		} else if len(resources) > 1 {
			for _, resource := range resources {
				key_value := strings.Split(resource, ":")
				if !state.Stocks[key_value[0]].isInitialized {
					name, details := initStock([]string{resource}, false)

					state.Stocks[name] = details
				}
			}
		}

		// Extract task details from the current line.
		re = regexp.MustCompile(`^(\w+):\((\w+:\d+(?:;\w+:\d+)*)\):\((\w+:\d+(?:;\w+:\d+)*)\):(\d+)$`)
		resources = re.FindAllString(text, -1)
		if len(resources) == 1 {
			task, details := makeTask(resources[0])

			state.addConsumersAndProducers(task, details)
		}

		re = regexp.MustCompile(`optimize:\(.+\)`)
		resources = re.FindAllString(text, -1)
		if len(resources) == 1 {
			if optimize_found == true {
				panic("Too many Optimizations")
			}
			optimize_found = true
			state.goal = getGoals(resources[0])
		}

	}

	return state
}

func initStock(s_arr []string, write_amount bool) (stoc_name string, stock_details StockDetails) {
	if write_amount {

		key_value := strings.Split(s_arr[0], ":")
		result, err := strconv.Atoi(key_value[1])
		check(err)
		stoc_name = key_value[0]
		stock_details = StockDetails{
			isInitialized: true,
			Amount:        result,
			Producers:     make(map[string]TaskDetails),
			Consumers:     make(map[string]TaskDetails),
		}

	} else {
		key_value := strings.Split(s_arr[0], ":")
		stoc_name = key_value[0]
		stock_details = StockDetails{
			isInitialized: true,
			Amount:        0,
			Producers:     make(map[string]TaskDetails),
			Consumers:     make(map[string]TaskDetails),
		}
	}
	return
}

func (state *State) addConsumersAndProducers(task string, task_details TaskDetails) {
	for stock_name := range task_details.StockNeeded {
		stock := state.Stocks[stock_name]
		stock.Consumers[task] = task_details
	}
	for stock_name := range task_details.StockResults {
		stock := state.Stocks[stock_name]
		stock.Producers[task] = task_details
	}
}

func getGoals(s string) map[string]bool {
	results := regexp.MustCompile(`\w+`).FindAllString(s, -1)
	mapped := map[string]bool{}
	for _, result := range results[1:] {
		mapped[result] = true
	}

	return mapped
}

// makeTask parses a string to extract task name and its details.
func makeTask(s string) (string, TaskDetails) {
	name := ""
	task := TaskDetails{}
	open_bracket := 0
	colon_count := 0
	str_store := ""

	for _, r := range s {
		// Handle nested content within parentheses.
		if r == '(' {
			open_bracket++
		}
		if r == ')' && open_bracket > 0 {
			open_bracket--
		}

		if r == ':' && open_bracket == 0 {
			colon_count++
			switch colon_count {
			case 1:
				// Extract task name.
				name = str_store
				str_store = ""
			case 2:
				// Extract task StockNeeded.
				task.StockNeeded = makeStockneededResultMap(str_store)
				str_store = ""
			case 3:
				// Extract task Results.
				task.StockResults = makeStockneededResultMap(str_store)
				str_store = ""
			}
			continue
		}

		str_store += string(r)

	}
	// Convert Delay (stored as string) to integer.
	Delay, err := strconv.Atoi(str_store)
	check(err)
	task.Delay = Delay

	return name, task
}

// makeStockneededResultMap generates a map of dependencies/Results from a string.
func makeStockneededResultMap(s string) map[string]int {
	return_map := make(map[string]int)
	Results := regexp.MustCompile("\\w+").FindAllString(s, -1)

	// Iterate over the Results in pairs (name, value).
	for i, value := range Results {
		if i%2 == 0 {
			result, err := strconv.Atoi(Results[i+1])
			check(err)
			return_map[value] = result
		}
	}

	return return_map
}
