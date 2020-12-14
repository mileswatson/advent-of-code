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
	str = strings.ReplaceAll(str, "mem[", "")
	str = strings.ReplaceAll(str, "]", "")

	lines := strings.Split(str, "\n")

	instructions := getInstructions(lines)

	fmt.Println(partOne(instructions))

	fmt.Println(partTwo(instructions))
}

func partTwo(instructions [][]string) int {
	m := make(map[int]int)

	mask := "XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
	for _, instruction := range instructions {
		if instruction[0] == "mask" {
			mask = instruction[1]
		} else {
			address, _ := strconv.Atoi(instruction[0])
			value, _ := strconv.Atoi(instruction[1])
			for _, possible := range getPossible(maskTwo(getBits(address), mask)) {
				a := getInt(possible)
				m[a] = value
			}
		}
	}
	total := 0
	for _, value := range m {
		total += value
	}
	return total
}

func partOne(instructions [][]string) int {
	m := make(map[int]int)

	mask := "XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
	for _, instruction := range instructions {
		if instruction[0] == "mask" {
			mask = instruction[1]
		} else {
			address, _ := strconv.Atoi(instruction[0])
			value, _ := strconv.Atoi(instruction[1])
			m[address] = getInt(maskOne(getBits(value), mask))
		}
	}
	total := 0
	for _, value := range m {
		total += value
	}
	return total
}

func getBits(x int) string {
	s := ""
	for i := 0; i < 36; i++ {
		if x%2 == 0 {
			s = "0" + s
		} else {
			s = "1" + s
		}
		x /= 2
	}
	return s
}

func getInt(x string) int {
	total := 0
	increase := 1
	for i := 35; i >= 0; i-- {
		if x[i] == '1' {
			total += increase
		}
		increase *= 2
	}
	return total
}

func maskOne(x string, mask string) string {
	s := ""
	for i := range mask {
		if mask[i] == 'X' {
			s += string(x[i])
		} else {
			s += string(mask[i])
		}
	}
	return s
}

func getPossible(address string) []string {
	if address == "" {
		return []string{""}
	}
	possible := getPossible(address[1:])
	if address[0] == 'X' {
		all := make([]string, len(possible)*2)
		copy(all[:len(possible)], prependToAll(possible, "0"))
		copy(all[len(possible):], prependToAll(possible, "1"))
		return all
	}
	return prependToAll(possible, string(address[0]))
}

func prependToAll(addresses []string, bit string) []string {
	prepended := make([]string, len(addresses))
	for i := range addresses {
		prepended[i] = bit + addresses[i]
	}
	return prepended
}

func maskTwo(x string, mask string) string {
	s := ""
	for i := range mask {
		if mask[i] == 'X' {
			s += "X"
		} else if mask[i] == '1' {
			s += "1"
		} else {
			s += string(x[i])
		}
	}
	return s
}

func getInstructions(lines []string) [][]string {
	instructions := make([][]string, len(lines))
	for i := range lines {
		instructions[i] = strings.Split(lines[i], " = ")
	}
	return instructions
}
