package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

const steps int = 6

func main() {
	data, _ := ioutil.ReadFile("input.txt")
	str := strings.ReplaceAll(string(data), "\r", "")
	lines := strings.Split(str, "\n")

	fmt.Println(partOne(newGrid(lines)))
	fmt.Println(partTwo(newHypergrid(lines)))
}

func partTwo(h *hypergrid) int {
	for i := 0; i < steps; i++ {
		h.next()
	}
	total := 0
	for _, frame1 := range h.arr {
		for _, frame2 := range frame1 {
			for _, row := range frame2 {
				for _, item := range row {
					if item {
						total++
					}
				}
			}
		}
	}
	return total
}

type hypergrid struct {
	arr  [][][][]bool
	size int
}

func (h *hypergrid) next() {
	size := h.size
	newArr := createHyperarray(size)
	for x := 0; x < size; x++ {
		for y := 0; y < size; y++ {
			for z := 0; z < size; z++ {
				for w := 0; w < size; w++ {
					newArr[x][y][z][w] = getHyperSurrounding(x, y, z, w, h.arr)
				}

			}
		}
	}
	h.arr = newArr
}

func getHyperSurrounding(x, y, z, w int, arr [][][][]bool) bool {
	total := 0
	size := len(arr)
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			for k := -1; k <= 1; k++ {
				for l := -1; l <= 1; l++ {
					if i == 0 && j == 0 && k == 0 && l == 0 {
						continue
					}
					newX := x + i
					newY := y + j
					newZ := z + k
					newW := w + l

					if newX < 0 || newX >= size {
						continue
					}
					if newY < 0 || newY >= size {
						continue
					}
					if newZ < 0 || newZ >= size {
						continue
					}
					if newW < 0 || newW >= size {
						continue
					}

					if arr[newX][newY][newZ][newW] {
						total++
					}
				}
			}
		}
	}
	isActive := arr[x][y][z][w]
	if isActive && (total == 2 || total == 3) {
		return true
	} else if !isActive && total == 3 {
		return true
	}
	return false
}

func createHyperarray(size int) [][][][]bool {
	arr := make([][][][]bool, size)
	for i := range arr {
		arr[i] = make([][][]bool, size)
		for j := range arr[i] {
			arr[i][j] = make([][]bool, size)
			for k := range arr[i][j] {
				arr[i][j][k] = make([]bool, size)
			}
		}
	}
	return arr
}

func newHypergrid(lines []string) *hypergrid {
	size := len(lines) + steps*2
	g := &hypergrid{createHyperarray(size), size}
	for x, line := range lines {
		x = x - len(lines)/2
		for y, char := range line {
			y = y - len(lines)/2
			if char == '#' {
				g.arr[size/2][size/2][x+size/2][y+size/2] = true
			} else {
				g.arr[size/2][size/2][x+size/2][y+size/2] = false
			}
		}
	}
	return g
}

func partOne(g *grid) int {
	for i := 0; i < steps; i++ {
		g.next()
	}
	total := 0
	for _, frame := range g.arr {
		for _, row := range frame {
			for _, item := range row {
				if item {
					total++
				}
			}
		}
	}
	return total
}

func (g *grid) next() {
	size := g.size
	newArr := createArray(size)
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			for k := 0; k < size; k++ {
				newArr[i][j][k] = getSurrounding(i, j, k, g.arr)
			}
		}
	}
	g.arr = newArr
}

func getSurrounding(x, y, z int, arr [][][]bool) bool {
	total := 0
	size := len(arr)
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			for k := -1; k <= 1; k++ {
				if i == 0 && j == 0 && k == 0 {
					continue
				}
				newX := x + i
				newY := y + j
				newZ := z + k

				if newX < 0 || newX >= size {
					continue
				}
				if newY < 0 || newY >= size {
					continue
				}
				if newZ < 0 || newZ >= size {
					continue
				}

				if arr[newX][newY][newZ] {
					total++
				}
			}
		}
	}
	isActive := arr[x][y][z]
	if isActive && (total == 2 || total == 3) {
		return true
	} else if !isActive && total == 3 {
		return true
	}
	return false
}

type grid struct {
	arr  [][][]bool
	size int
}

func createArray(size int) [][][]bool {
	arr := make([][][]bool, size)
	for i := range arr {
		arr[i] = make([][]bool, size)
		for j := range arr[i] {
			arr[i][j] = make([]bool, size)
		}
	}
	return arr
}

func newGrid(lines []string) *grid {
	size := len(lines) + steps*2
	g := &grid{createArray(size), size}
	for x, line := range lines {
		x = x - len(lines)/2
		for y, char := range line {
			y = y - len(lines)/2
			if char == '#' {
				g.arr[size/2][x+size/2][y+size/2] = true
			} else {
				g.arr[size/2][x+size/2][y+size/2] = false
			}
		}
	}
	return g
}

func (g *grid) print() {
	zp := g.size / 2
	for z, frame := range g.arr {
		fmt.Println("z =", z-zp)
		for _, row := range frame {
			for _, item := range row {
				if item {
					fmt.Print("#")
				} else {
					fmt.Print(".")
				}
			}
			fmt.Println()
		}
		fmt.Println()
	}
}
