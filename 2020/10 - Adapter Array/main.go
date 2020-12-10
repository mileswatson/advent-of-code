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

	nodes, max := createGraph(lines)

	fmt.Println(partOne(nodes, max))

	fmt.Println(partTwo(nodes, max))
}

func partOne(nodes map[int][]int, max int) int {
	count := make([]int, 3)
	current := 0
	for current != max {
		for i := 0; i < 3; i++ {
			_, exists := nodes[current+i+1]
			if exists {
				count[i]++
				current = current + i + 1
				break
			}
		}
	}
	return count[0] * (count[2] + 1)
}

func partTwo(nodes map[int][]int, target int) int {
	memo := make([]int, target+1)
	for index := range memo {
		memo[index] = -1
	}
	memo[target] = 1
	return countDFS(nodes, target, 0, memo)
}

func countDFS(nodes map[int][]int, target, current int, memo []int) int {
	if memo[current] != -1 {
		return memo[current]
	}
	total := 0
	if current == target {
		total = 1
	} else {
		for _, item := range nodes[current] {
			total += countDFS(nodes, target, item, memo)
		}
	}
	memo[current] = total
	return total
}

func createGraph(lines []string) (map[int][]int, int) {
	nodes := make(map[int][]int)
	nodes[0] = make([]int, 0)

	max := 0
	for _, line := range lines {
		num, _ := strconv.Atoi(line)
		nodes[num] = make([]int, 0)
		if num > max {
			max = num
		}
	}

	for node := range nodes {
		for i := 1; i <= 3; i++ {
			_, exists := nodes[node+i]
			if exists {
				nodes[node] = append(nodes[node], node+i)
			}
		}
	}

	return nodes, max
}
