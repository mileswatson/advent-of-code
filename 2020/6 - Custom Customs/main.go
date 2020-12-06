package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	file, _ := ioutil.ReadFile("input.txt")

	cleaned := strings.ReplaceAll(string(file), "\r", "")

	combined := combine(strings.Split(cleaned, "\n"))

	fmt.Println(partOne(combined))

	fmt.Println(partTwo(combined))
}

func combine(lines []string) [][]string {
	var output [][]string
	var current []string
	for _, line := range lines {
		if len(line) == 0 {
			output = append(output, current)
			current = make([]string, 0)
			continue
		}
		current = append(current, line)

	}
	return output
}

func partOne(groups [][]string) int {
	total := 0
	for _, group := range groups {
		set := make(map[rune]bool)
		for _, line := range group {
			for _, char := range line {
				set[char] = true
			}
		}
		total += len(set)
	}
	return total
}

func partTwo(groups [][]string) int {
	total := 0
	for _, group := range groups {
		set := make(map[rune]int)
		for _, line := range group {
			for _, char := range line {
				_, exists := set[char]
				if exists {
					set[char]++
				} else {
					set[char] = 1
				}
			}
		}
		for _, value := range set {
			if value == len(group) {
				total++;
			}
		}
	}
	return total
}

