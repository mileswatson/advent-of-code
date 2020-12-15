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

	parts := strings.Split(str, ",")

	numbers := make([]int, len(parts))

	for index, part := range parts {
		numbers[index], _ = strconv.Atoi(part)
	}

	fmt.Println(partOne(numbers, 2020))
	fmt.Println(partTwo(numbers, 30000000))
}

func partOne(numbers []int, target int) int {
	memory := make(map[int]int)
	mostRecent := 0
	for index, number := range numbers {
		if index != 0 {
			memory[mostRecent] = index - 1
		}
		mostRecent = number
	}
	for i := len(numbers); i < target; i++ {
		last, contains := memory[mostRecent]
		memory[mostRecent] = i - 1
		if contains {
			mostRecent = i - last - 1
		} else {
			mostRecent = 0
		}
	}
	return mostRecent
}

func partTwo(numbers []int, target int) int {
	memory := make(map[int]int, 5000000)
	mostRecent := 0
	for index, number := range numbers {
		if index != 0 {
			memory[mostRecent] = index - 1
		}
		mostRecent = number
	}
	for i := len(numbers); i < target; i++ {
		last, contains := memory[mostRecent]
		memory[mostRecent] = i - 1
		if contains {
			//fmt.Println(mostRecent, contains, i-last)
			mostRecent = i - last - 1
		} else {
			//fmt.Println(mostRecent, contains, 0)
			mostRecent = 0
		}
	}
	return mostRecent
}
