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

	parts := strings.Split(str, "\n\n")

	ruleLines := strings.Split(parts[0], "\n")
	messages := strings.Split(parts[1], "\n")

	rules := getRules(ruleLines)

	fmt.Println(partOne(messages, rules))

	fmt.Println(partTwo(messages, rules))
}

func partTwo(messages []string, rules map[int]rule) int {
	maxLength := 0
	for _, message := range messages {
		if len(message) > maxLength {
			maxLength = len(message)
		}
	}

	memo := make(map[int][]string)

	fourtytwo := make(map[string]bool)
	for _, option := range rules[42].generate(rules, memo) {
		fourtytwo[option] = true
	}

	thirtyone := make(map[string]bool)
	for _, option := range rules[31].generate(rules, memo) {
		thirtyone[option] = true
	}

	total := 0
	for _, message := range messages {
		if check42(message, fourtytwo, thirtyone, 0) {
			total++
		}
	}
	return total
}

func check42(message string, fourtytwo, thirtyone map[string]bool, found int) bool {
	if message == "" {
		return false
	}
	for i := len(message); i > 0; i-- {
		if fourtytwo[message[:i]] {
			if found > 0 {
				if check31(message[i:], thirtyone, 0, found) {
					return true
				}
			}
			return check42(message[i:], fourtytwo, thirtyone, found+1)
		}
	}
	return false
}

func check31(message string, thirtyone map[string]bool, found, max int) bool {
	if len(message) == 0 {
		return found > 0 && max >= 0
	}
	if max == 0 {
		return false
	}
	for i := len(message); i > 0; i-- {
		if thirtyone[message[:i]] && check31(message[i:], thirtyone, found+1, max-1) {
			return true
		}
	}
	return false
}

func partOne(messages []string, rules map[int]rule) int {
	options := make(map[string]bool)
	memo := make(map[int][]string)
	for _, option := range rules[0].generate(rules, memo) {
		options[option] = true
	}
	total := 0
	for _, message := range messages {
		if options[message] {
			total++
		}
	}
	return total
}

func getRules(lines []string) map[int]rule {
	rules := make(map[int]rule)

	for _, line := range lines {
		parts := strings.Split(line, ": ")
		num, _ := strconv.Atoi(parts[0])
		if parts[1][0] == '"' {
			rules[num] = &charRule{num, string(parts[1][1])}
			continue
		}
		totalRules := make([][]int, 0)
		for _, ruleSet := range strings.Split(parts[1], " | ") {
			rules := make([]int, 0)
			for _, rule := range strings.Split(ruleSet, " ") {
				str, _ := strconv.Atoi(rule)
				rules = append(rules, str)
			}
			totalRules = append(totalRules, rules)
		}
		rules[num] = &optionRule{num, totalRules}
	}

	return rules
}

type rule interface {
	number() int
	generate(map[int]rule, map[int][]string) []string
}

type charRule struct {
	num   int
	match string
}

func (r *charRule) number() int {
	return r.num
}

func (r *charRule) generate(dictionary map[int]rule, memo map[int][]string) []string {
	val, contains := memo[r.num]
	if contains {
		return val
	}
	val = []string{r.match}
	memo[r.num] = val
	return val
}

type optionRule struct {
	num      int
	rulesets [][]int
}

func (r *optionRule) number() int {
	return r.num
}

func (r *optionRule) generate(dictionary map[int]rule, memo map[int][]string) []string {
	val, contains := memo[r.num]
	if contains {
		return val
	}

	totalRules := make([]string, 0)
	for _, ruleset := range r.rulesets {
		setRules := []string{""}
		for _, ruleNum := range ruleset {
			rule := dictionary[ruleNum]
			generated := rule.generate(dictionary, memo)
			newArr := make([]string, 0)
			for _, option := range generated {
				for _, setRule := range setRules {
					newArr = append(newArr, setRule+option)
				}
			}
			setRules = newArr
		}
		totalRules = append(totalRules, setRules...)
	}

	memo[r.num] = totalRules
	return totalRules
}
