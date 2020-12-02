package main

import (
    "bufio"
    "fmt"
    "io"
    "io/ioutil"
    "strconv"
    "strings"
)

func ReadInts(r io.Reader) []int {
    scanner := bufio.NewScanner(r)
    scanner.Split(bufio.ScanWords)
    var result []int
    for scanner.Scan() {
        x, _ := strconv.Atoi(scanner.Text())
        result = append(result, x)
    }
    return result
}

func main() {
    bytes, _ := ioutil.ReadFile("input.txt")
    numbers := ReadInts(strings.NewReader(string(bytes)))

    fmt.Println(partOne(numbers))
    fmt.Println(partTwo(numbers))
}

func partOne(numbers []int) int {
    for i:=0; i<len(numbers)-1; i++ {
        for j:=i+1;j<len(numbers); j++ {
            if numbers[i] + numbers[j] == 2020 {
                return numbers[i]*numbers[j]
            }
        }
    }
    return -1;
}

func partTwo(numbers []int) int {
    for i:=0; i<len(numbers)-2; i++ {
        for j:=i+1;j<len(numbers)-1; j++ {
            for k:=j+1;k<len(numbers); k++ {
                if numbers[i] + numbers[j] + numbers[k] == 2020 {
                    return numbers[i]*numbers[j]*numbers[k]
                }
            }
        }
    }
    return -1;
}
