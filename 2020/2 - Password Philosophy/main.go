package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// Line is a struct to represent each password and policy
type Line struct {
	password string
	target   string
	num1     int
	num2     int
}

func main() {

	data, _ := ioutil.ReadFile("input.txt")

	str := strings.ReplaceAll(string(data), "\r", "")

	lineArray := strings.Split(str, "\n")

	lineSlice := lineArray[:]

	var lines = make([]*Line, len(lineSlice))

	for i := 0; i < len(lineSlice); i++ {
		lines[i] = new(Line)
		fmt.Sscanf(lineSlice[i], "%d-%d %s %s",
			&lines[i].num1,
			&lines[i].num2,
			&lines[i].target,
			&lines[i].password,
		)
		lines[i].target = lines[i].target[:1]
	}

	fmt.Println(partOne(lines))
	fmt.Println(partTwo(lines))

}

func partOne(lines []*Line) int {
	total := 0

	for _, line := range lines {
		num := strings.Count(line.password, line.target)
		if line.num1 <= num && num <= line.num2 {
			total++
		}
	}

	return total
}

func partTwo(lines []*Line) int {
	total := 0

	for _, line := range lines {
		atOne := line.password[line.num1-1:line.num1] == line.target
		atTwo := line.password[line.num2-1:line.num2] == line.target

		if (atOne || atTwo) && !(atOne && atTwo) {
			total++
		}

	}

	return total
}
