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

	time, _ := strconv.Atoi(lines[0])

	buses := getBuses(lines[1])

	fmt.Println(partOne(int64(time), buses))
	fmt.Println(partTwo(int64(time), buses))
}

func partTwo(time int64, buses []int64) int64 {
	var counter int64 = 0
	var increase int64 = 1
	for wait, divisor := range buses {
		w := int64(wait)
		d := int64(divisor)
		if d == 0 {
			continue
		}
		r := (d - w) % d
		if r < 0 {
			r += d
		}
		for i := int64(0); true; i++ {
			testNum := counter + i*increase
			if testNum%d == r {
				counter = testNum
				increase *= d
				break
			}
		}
	}
	return counter
}

func partOne(time int64, buses []int64) int64 {
	first := true
	minWait := int64(0)
	closestBus := int64(0)
	for _, bus := range buses {
		if bus == 0 {
			continue
		}
		wait := (bus - (time % bus)) % bus
		if first || wait < minWait {
			first = false
			minWait = wait
			closestBus = bus
		}
	}
	return closestBus * minWait
}

func getBuses(list string) []int64 {
	buses := make([]int64, 0)
	for _, b := range strings.Split(list, ",") {
		if b == "x" {
			buses = append(buses, 0)
		} else {
			busNum, _ := strconv.Atoi(b)
			buses = append(buses, int64(busNum))
		}
	}
	return buses
}
