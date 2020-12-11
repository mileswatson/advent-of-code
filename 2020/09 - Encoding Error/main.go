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

	numbers := make([]int, len(lines))

	for index, line := range lines {
		numbers[index], _ = strconv.Atoi(line)
	}

	weakness := partOne(numbers)

	fmt.Println(weakness)

	fmt.Println(partTwo(numbers, weakness))
}

func partTwo(numbers []int, weakness int) int {
	numRange := getRange(numbers, weakness)
	var min, max int
	firstMin := true
	for _, num := range numRange {
		if firstMin {
			min = num
			max = num
			firstMin = false
		} else if num < min {
			min = num
		} else if num > max {
			max = num
		}
	}
	return min + max
}

func getRange(numbers []int, weakness int) []int {
	sum := make([]int, len(numbers)+1)
	for i := 0; i < len(numbers); i++ {
		sum[i+1] = sum[i] + numbers[i]
		for j := 0; j < i; j++ {
			if sum[i+1]-sum[j] == weakness {
				return numbers[j : i+1]
			}
		}
	}
	return []int{}
}

func partOne(numbers []int) int {
	const n int = 25
	preamble := make(map[int]bool)

	for i := 0; i < n; i++ {
		preamble[numbers[i]] = true
	}

	for i := n; i < len(numbers); i++ {
		current := numbers[i]
		found := false
		for key := range preamble {
			if preamble[current-key] {
				if current == 2*key {
					continue
				}
				found = true
				break
			}
		}
		if !found {
			return current
		}
		delete(preamble, numbers[i-n])
		preamble[numbers[i]] = true
	}

	return -1
}
