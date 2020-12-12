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
	instructions := getInstructions(lines)
	fmt.Println(partOne(instructions))
	fmt.Println(partTwo(instructions))
}

func partTwo(instructions []instruction) int {
	wx, wy := 10, 1
	x, y := 0, 0
	for _, i := range instructions {
		switch i.direction {
		case north:
			wy += i.value
		case south:
			wy -= i.value
		case east:
			wx += i.value
		case west:
			wx -= i.value
		case forward:
			x += wx * i.value
			y += wy * i.value
		case left, right:
			anticlockwise := i.value / 90
			if i.direction == right {
				anticlockwise = 4 - anticlockwise
			}
			anticlockwise %= 4
			switch anticlockwise {
			case 1:
				temp := wy
				wy = wx
				wx = -temp
			case 2:
				wy = -wy
				wx = -wx
			case 3:
				temp := wy
				wy = -wx
				wx = temp
			}
		}
	}
	return abs(x) + abs(y)
}

func partOne(instructions []instruction) int {
	ship := &boat{0, 0, 0}
	for _, instruction := range instructions {
		ship.follow(instruction)
	}
	return abs(ship.x) + abs(ship.y)
}

func abs(x int) int {
	if x >= 0 {
		return x
	}
	return -x
}

func (ship *boat) follow(i instruction) {
	if i.direction < 4 {
		temp := ship.direction
		ship.direction = i.direction
		ship.move(i.value)
		ship.direction = temp
	} else if i.direction == left {
		ship.direction = (ship.direction + (i.value / 90)) % 4
	} else if i.direction == right {
		ship.direction = (ship.direction - (i.value / 90)) % 4
	} else {
		ship.move(i.value)
	}
	if ship.direction < 0 {
		ship.direction += 4
	}
}

func (ship *boat) move(distance int) {
	switch ship.direction {
	case east:
		ship.x += distance
	case west:
		ship.x -= distance
	case north:
		ship.y += distance
	case south:
		ship.y -= distance
	}
}

func getInstructions(lines []string) []instruction {
	instructions := make([]instruction, 0)

	for _, line := range lines {
		first := line[0:1]
		var direction int
		switch first {
		case "N":
			direction = north
		case "S":
			direction = south
		case "E":
			direction = east
		case "W":
			direction = west
		case "L":
			direction = left
		case "R":
			direction = right
		case "F":
			direction = forward
		}
		value, _ := strconv.Atoi(line[1:len(line)])
		x := instruction{
			direction,
			value,
		}
		instructions = append(instructions, x)
	}

	return instructions
}

type boat struct {
	direction int
	x         int
	y         int
}

type instruction struct {
	direction int
	value     int
}

const (
	east = iota
	north
	west
	south
	left
	right
	forward
)
