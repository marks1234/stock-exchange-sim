package components

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
	"strings"
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

// regex to find all the matches for the processes
//^(\w+):\((\w+:\d+(?:;\w+:\d+)*)\):\((\w+:\d+(?:;\w+:\d+)*)\):(\d+)$

func InitState(file *os.File) State {
	state := State{
		time:  0,
		stock: map[string]int{},
		tasks: map[string]TaskDetails{},
	}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		re := regexp.MustCompile(`\b\w+:[\d\.]+\b`)
		resources := re.FindAllString(scanner.Text(), -1)
		if len(resources) == 1 {
			key_value := strings.Split(resources[0], ":")
			result, err := strconv.Atoi(key_value[1])
			check(err)
			state.stock[key_value[0]] = result
			continue
		}
		re = regexp.MustCompile(`^(\w+):\((\w+:\d+(?:;\w+:\d+)*)\):\((\w+:\d+(?:;\w+:\d+)*)\):(\d+)$`)
		resources = re.FindAllString(scanner.Text(), -1)
		if len(resources) == 1 {
			name, task := makeTask(resources[0])
			state.tasks[name] = task
		}

	}

	return state
}

func makeTask(s string) (string, TaskDetails) {
	name := ""
	task := TaskDetails{}
	open_bracket := 0
	colon_count := 0
	str_store := ""
	for _, r := range s {
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
				name = str_store
				str_store = ""
				continue
			case 2:
				task.dependecies = makeDependencyResultMap(str_store)
				str_store = ""
				continue
			case 3:
				task.results = makeDependencyResultMap(str_store)
				str_store = ""
				continue
			}
		}

		str_store += string(r)

	}
	delay, err := strconv.Atoi(str_store)
	check(err)
	task.delay = delay

	return name, task
}

func makeDependencyResultMap(s string) map[string]int {
	return_map := map[string]int{}
	results := regexp.MustCompile("\\w+").FindAllString(s, -1)

	for i, value := range results {
		if i%2 == 0 {
			result, err := strconv.Atoi(results[i+1])
			check(err)
			return_map[value] = result
		}
	}

	return return_map
}
