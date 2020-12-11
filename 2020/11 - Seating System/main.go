package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	data, _ := ioutil.ReadFile("input.txt")

	str := strings.ReplaceAll(string(data), "\r", "")

	lines := strings.Split(str, "\n")

	seats := getSeats(lines)

	fmt.Println(partOne(seats))

	fmt.Println(partTwo(seats))

}

func partOne(seats [][]int) int {
	changed := true
	for changed {
		seats, changed = nextOne(seats)
	}
	total := 0
	for _, row := range seats {
		for _, seat := range row {
			if seat == occupied {
				total++
			}
		}
	}
	return total
}

func partTwo(seats [][]int) int {
	changed := true
	for changed {
		seats, changed = nextTwo(seats)
	}
	total := 0
	for _, row := range seats {
		for _, seat := range row {
			if seat == occupied {
				total++
			}
		}
	}
	return total
}

func nextTwo(seats [][]int) ([][]int, bool) {
	newSeats := make([][]int, len(seats))
	for i := 0; i < len(seats); i++ {
		newSeats[i] = make([]int, len(seats[0]))
	}
	changed := false
	for x := range seats {
		for y := range seats[x] {
			numOccupied := countVisibleSeats(seats, x, y)
			if seats[x][y] == empty && numOccupied == 0 {
				newSeats[x][y] = occupied
				changed = true
			} else if seats[x][y] == occupied && numOccupied >= 5 {
				newSeats[x][y] = empty
				changed = true
			} else {
				newSeats[x][y] = seats[x][y]
			}
		}
	}
	return newSeats, changed
}

func countVisibleSeats(seats [][]int, x, y int) int {
	checks := [][2]int{
		[2]int{-1, -1},
		[2]int{-1, 0},
		[2]int{-1, 1},
		[2]int{0, -1},
		[2]int{0, 1},
		[2]int{1, -1},
		[2]int{1, 0},
		[2]int{1, 1},
	}

	total := 0
	for _, check := range checks {
		i := 1
		for {
			tempx := x + i*check[0]
			tempy := y + i*check[1]
			if tempx < 0 || tempx >= len(seats) {
				break
			} else if tempy < 0 || tempy >= len(seats[0]) {
				break
			} else if seats[tempx][tempy] == empty {
				break
			} else if seats[tempx][tempy] == occupied {
				total++
				break
			}
			i++
		}

	}
	return total
}

func nextOne(seats [][]int) ([][]int, bool) {
	newSeats := make([][]int, len(seats))
	for i := 0; i < len(seats); i++ {
		newSeats[i] = make([]int, len(seats[0]))
	}
	changed := false
	for x := range seats {
		for y := range seats[x] {
			numOccupied := countAdjacentSeats(seats, x, y)
			if seats[x][y] == empty && numOccupied == 0 {
				newSeats[x][y] = occupied
				changed = true
			} else if seats[x][y] == occupied && numOccupied >= 4 {
				newSeats[x][y] = empty
				changed = true
			} else {
				newSeats[x][y] = seats[x][y]
			}
		}
	}
	return newSeats, changed
}

func countAdjacentSeats(seats [][]int, x, y int) int {
	checks := [][2]int{
		[2]int{-1, -1},
		[2]int{-1, 0},
		[2]int{-1, 1},
		[2]int{0, -1},
		[2]int{0, 1},
		[2]int{1, -1},
		[2]int{1, 0},
		[2]int{1, 1},
	}

	total := 0
	for _, check := range checks {
		tempx := x + check[0]
		tempy := y + check[1]
		if tempx < 0 || tempx >= len(seats) {
			continue
		}
		if tempy < 0 || tempy >= len(seats[0]) {
			continue
		}
		if seats[tempx][tempy] == occupied {
			total++
		}
	}
	return total
}

const (
	floor = iota
	empty
	occupied
)

func getSeats(lines []string) [][]int {
	seats := make([][]int, len(lines))
	for i, line := range lines {
		seats[i] = make([]int, len(line))
		for j, chr := range line {
			switch chr {
			case '.':
				seats[i][j] = floor
			case 'L':
				seats[i][j] = empty
			}
		}
	}
	return seats
}
