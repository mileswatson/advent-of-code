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

	fields := getFields(parts[0])

	myTicket := getMyTicket(parts[1])

	nearbyTickets := getNearbyTickets(parts[2])

	num1, invalid := partOne(fields, nearbyTickets)

	fmt.Println(num1)

	validTickets := make([][]int, 0)
	for index, ticket := range nearbyTickets {
		if !invalid[index] {
			validTickets = append(validTickets, ticket)
		}
	}

	fmt.Println(partTwo(fields, validTickets, myTicket))

}

func partTwo(fields []*field, validTickets [][]int, myTicket []int) int {
	assignements := make([]*field, len(fields))
	for i := range assignements {
		assignements[i] = nil
	}
	memo := make(map[int]bool)
	getValid(assignements, fields, validTickets, memo)
	total := 1
	for i, f := range assignements {
		if len(f.name) < 9 || f.name[:9] != "departure" {
			continue
		}
		total *= myTicket[i]
	}
	return total
}

func getValid(assignments []*field, available []*field, tickets [][]int, memo map[int]bool) bool {
	if len(available) == 0 {
		return true
	}
	currentField := available[0]
	for i := range assignments {
		if assignments[i] != nil {
			continue
		}
		if !checkField(currentField, tickets, len(assignments)-len(available), i, memo) {
			continue
		}
		assignments[i] = currentField
		if getValid(assignments, available[1:], tickets, memo) {
			return true
		}
		assignments[i] = nil
	}
	return false
}

func checkField(f *field, validTickets [][]int, fPosition, aPosition int, memo map[int]bool) bool {
	memoLocation := fPosition*len(validTickets[0]) + aPosition
	val, present := memo[memoLocation]
	if present {
		return val
	}
	for _, ticket := range validTickets {
		if !f.check(ticket[aPosition]) {
			memo[memoLocation] = false
			return false
		}
	}
	memo[memoLocation] = true
	return true
}

func partOne(fields []*field, tickets [][]int) (int, map[int]bool) {
	total := 0
	invalid := make(map[int]bool, 0)
	for i, ticket := range tickets {
		for _, num := range ticket {
			found := false
			for _, field := range fields {
				if field.check(num) {
					found = true
					break
				}
			}
			if !found {
				total += num
				invalid[i] = true
			}
		}

	}
	return total, invalid
}

func getMyTicket(part string) []int {
	line := strings.Split(part, "\n")[1]
	numbers := strings.Split(line, ",")
	ticket := make([]int, len(numbers))
	for j, number := range numbers {
		ticket[j], _ = strconv.Atoi(number)
	}
	return ticket
}

func getNearbyTickets(part string) [][]int {
	lines := strings.Split(part, "\n")[1:]
	tickets := make([][]int, len(lines))

	for i, line := range lines {
		numbers := strings.Split(line, ",")
		ints := make([]int, len(numbers))
		for j, number := range numbers {
			ints[j], _ = strconv.Atoi(number)
		}
		tickets[i] = ints
	}

	return tickets
}

type field struct {
	name  string
	check func(int) bool
}

func getFields(part string) []*field {
	lines := strings.Split(part, "\n")
	fields := make([]*field, len(lines))

	for index, line := range lines {
		parts := strings.Split(line, ": ")
		name := parts[0]
		ranges := strings.Split(parts[1], " or ")
		lim0 := strings.Split(ranges[0], "-")
		lim1 := strings.Split(ranges[1], "-")

		lower0, _ := strconv.Atoi(lim0[0])
		upper0, _ := strconv.Atoi(lim0[1])

		lower1, _ := strconv.Atoi(lim1[0])
		upper1, _ := strconv.Atoi(lim1[1])

		f := &field{
			name,
			func(x int) bool {
				range0 := lower0 <= x && x <= upper0
				range1 := lower1 <= x && x <= upper1
				return range0 || range1
			},
		}
		fields[index] = f
	}
	return fields
}
