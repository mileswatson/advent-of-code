package main

import (
    "fmt"
    "io/ioutil"
    "strings"
    "strconv"
)

func main() {

    data, _ := ioutil.ReadFile("input.txt")

    lineArray := strings.Split(string(data), "\n")

    lines := lineArray[:]
    
    fmt.Println(partOne(lines))
    fmt.Println(partTwo(lines))

}

func partOne(lines []string) int {
    passports := combine(lines)
    total := 0
    for _, passport := range passports {
        if containsAllFields(passport) {
            total++
        }
    }
    return total;
}

func partTwo(lines []string) int {
    passports := combine(lines)
    total := 0
    for _, passport := range passports {
        if (!containsAllFields(passport)) {
            continue
        }
        var valid bool
        for _, field := range getFields(passport) {
            switch field.name {
            case "byr":
                valid = digits(field.value, 4, 1920, 2002)
            case "iyr":
                valid = digits(field.value, 4, 2010, 2020)
            case "eyr":
                valid = digits(field.value, 4, 2020, 2030)
            case "hgt":
                length := len(field.value)
                units := field.value[length-2:length]
                value := field.value[:length-2]
                if (units == "cm") {
                    valid = digits(value, 3, 150, 193)
                } else if (units == "in") {
                    valid = digits(value, 2, 59, 76)
                } else {
                    valid = false
                }
            case "hcl":
                if (field.value[0:1] != "#") {
                    valid = false
                }
                valid = allChars(field.value[1:], "0123456789abcdef", 6)
            case "ecl":
                combinations := []string {"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}
                valid = contains(combinations, field.value)
            case "pid":
                valid = allChars(field.value, "0123456789", 9)
            case "cid":
                valid = true
            default:
                valid = false
            }
            if !valid {
                break
            }
        }
        if valid {
            total++
        }
    }

    return total;
}

func contains(s []string, e string) bool {
    for _, a := range s {
        if a == e {
            return true
        }
    }
    return false
}

func allChars(str, allowed string, length int) bool {
    if len(str) != length {
        return false
    }
    valid := true
    for _, char := range str {
        if !strings.Contains(str, string(char)) {
            valid = false
            break
        }
    }
    return valid
}

func combine(lines []string) []string {
    var passports []string
    combined := ""
    for _, line := range lines {
        if (len(line) != 0) {
            combined += line + " "
            continue
        }
        passports = append(passports, combined[:len(combined)-1])
        combined = ""
    }
    return passports
}

func containsAllFields(passport string) bool {
    required := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}

    allFound := true
    for _, field := range required {
        if !strings.Contains(passport, field) {
            allFound = false
            break
        }
    }

    return allFound;
}

type Field struct {
    name    string
    value   string
}

func getFields(passport string) []Field {
    var fields []Field
    parts := strings.Split(passport, " ")

    for _, part := range parts {
        field := Field{
            part[:3], part[4:],
        }
        fields = append(fields, field)
    }

    return fields
}

func digits(str string, num, min, max int) bool {
    if (len(str) != num) {
        return false
    }

    i, err := strconv.Atoi(str)

    if err != nil {
        return false;
    }

    if i < min || i > max {
        return false;
    }

    return true;
}
