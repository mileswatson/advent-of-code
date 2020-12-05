package main

import (
    "fmt"
    "io/ioutil"
    "strings"
)

func main() {
    file, _ := ioutil.ReadFile("input.txt")
    lines := strings.Split(string(file), "\n")
    fmt.Println(partOne(lines))
    fmt.Println(partTwo(lines))
}

func partOne(passes []string) int {
    largest := -1
    for _, pass := range passes {
        id := getID(getSeat(pass))
        if (id > largest) {
            largest = id
        }
    }
    return largest
}

func partTwo(passes []string) int {
    const MAX_ID = 127 * 8 + 7
    smallest := MAX_ID
    var seats [MAX_ID]bool;
    for _, pass := range passes {
        id := getID(getSeat(pass))
        seats[id] = true;
        if (id < smallest) {
            smallest = id
        }
    }
    for i := smallest; i < len(seats); i++ {
        if !seats[i] {
            return i;
        }
    }
    return -1;
}


func getSeat(pass string) (int, int) {
    return decodeBinary(pass[0:7]), decodeBinary(pass[7:10]);
}

func getID(row, column int) int {
    return 8 * row + column
}

func decodeBinary(pass string) int {
    total := 0
    current := 1 << (len(pass)-1)
    for _, char := range pass {
        if (char == 'B' || char == 'R') {
            total += current
        }
        current /= 2
    }
    return total
}
