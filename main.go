package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	s := "Apple:(Red:3;Green:5):(Taste:7;Texture:8):9" +
		"\nBanana:(Yellow:5):(Flavor:4;Shape:2):6" +
		"\nOrange:(Color:10):(Citrus:4;Size:8):12" +
		"\nCherry:(Red:3):(Sweet:4;Sour:2):5" +
		"\nMango:(Yellow:2;Green:5):(Taste:6;Feel:2):10" +
		"\nKiwi:(Green:6):(Tasty:3;Hard:1):4" +
		"\nBlueberry:(Blue:8):(Tart:3;Size:1):3" +
		"\nPear:(Green:2;Yellow:4):(Juicy:9;Size:7):8" +
		"\nRaspberry:(Red:5):(Taste:2;Texture:3):6" +
		"\nPapaya:(Orange:4):(Flavor:5;Texture:7):10"

	re := regexp.MustCompile("^(\\w+):\\((\\w+:\\d+(?:;\\w+:\\d+)*)\\):\\((\\w+:\\d+(?:;\\w+:\\d+)*)\\):(\\d+)$")
	for i, line := range strings.Split(s, "\n") {
		if re.MatchString(line) {
			fmt.Println(line)
			fmt.Println(i)
		}
	}
}
