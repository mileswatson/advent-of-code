package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {

	data, _ := ioutil.ReadFile("input.txt")

	str := strings.ReplaceAll(string(data), "\r", "")

	lines := strings.Split(str, "\n")

	bags := getBags(lines)

	fmt.Println(partOne(bags, "shiny gold"))
	fmt.Println(partTwo(bags, "shiny gold"))
}

func partOne(bags map[string]map[string]int, target string) int {
	memo := make(map[string]int)
	total := 0
	memo[target] = yes
	for color := range bags {
		if color == target {
			continue
		}
		if checkBag(bags, memo, color, target) {
			total++
		}
	}
	return total
}

func partTwo(bags map[string]map[string]int, target string) int {
	memo := make(map[string]int)
	return containedBags(bags, memo, target)
}

func containedBags(bags map[string]map[string]int,
	memo map[string]int,
	color string) int {

	number, contains := memo[color]

	if contains {
		return number
	}
	total := 0
	for contains, number := range bags[color] {
		total += number * (1 + containedBags(bags, memo, contains))
	}
	return total
}

const (
	no = iota
	yes
)

func checkBag(bags map[string]map[string]int,
	memo map[string]int,
	current string,
	target string) bool {

	status, contains := memo[current]
	if contains {
		if status == no {
			return false
		} else if status == yes {
			return true
		}
	}

	for inner := range bags[current] {
		if inner == target || checkBag(bags, memo, inner, target) {
			memo[current] = yes
			return true
		}
	}
	memo[current] = no
	return false
}

func getBags(lines []string) map[string]map[string]int {
	bags := make(map[string]map[string]int)

	for _, line := range lines {
		parts := strings.Split(line, " bags contain ")
		name := parts[0]
		trimmed := parts[1][:len(parts[1])-1]
		contains := make(map[string]int)
		bags[name] = contains
		if trimmed == "no other bags" {
			continue
		}
		for _, bag := range strings.Split(trimmed, ", ") {
			number, _ := strconv.Atoi(bag[0:1])
			if number == 1 {
				bag = bag[2 : len(bag)-4]
			} else {
				bag = bag[2 : len(bag)-5]
			}
			contains[bag] = number
		}
	}
	return bags
}
