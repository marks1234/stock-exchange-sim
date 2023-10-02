package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type State struct {
	time  int
	stock map[string]int
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

	initState(file)
}

func initState(file *os.File) State {
	state := State{
		time: 0,
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
		}
	}
	return state
}
