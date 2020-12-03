package main

import (
    "fmt"
    "io/ioutil"
    "strings"
)

func main() {

    data, _ := ioutil.ReadFile("input.txt")

    lineArray := strings.Split(string(data), "\n")

    lines := lineArray[:]
    
    fmt.Println(partOne(lines, 3, 1))
    fmt.Println(partTwo(lines))

}

func partOne(lines []string, right int, down int) int {
    total := 0
    across := 0
    for i := 0; i < len(lines); i += down {
        row := lines[i]
        if row[across] == '#' {
            total++
        }
        across = (across + right) % len(row)
    }
    return total
}

func partTwo(lines [] string) int {
    product := 1
    product *= partOne(lines, 1, 1)
    product *= partOne(lines, 3, 1)
    product *= partOne(lines, 5, 1)
    product *= partOne(lines, 7, 1)
    product *= partOne(lines, 1, 2)
    return product
}
