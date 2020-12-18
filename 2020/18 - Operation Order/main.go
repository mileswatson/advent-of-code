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
	str = strings.ReplaceAll(str, "(", "( ")
	str = strings.ReplaceAll(str, ")", " )")

	lines := strings.Split(str, "\n")

	fmt.Println(partOne(lines))
	fmt.Println(partTwo(lines))
}

func partOne(lines []string) int {
	total := 0
	for _, line := range lines {
		tokens := tokenise(line)
		total += evaluate(tokens)
	}
	return total
}

func partTwo(lines []string) int {
	total := 0
	for _, line := range lines {
		tokens := insertBrackets(tokenise(line))
		total += evaluate(tokens)
	}
	return total
}

func insertBrackets(tokens []string) []string {
	i := 0
	for i < len(tokens) {
		if tokens[i] == "+" {
			openBracket := nextValue(tokens, i-1, -1)
			tokens = insert(tokens, openBracket, "(")
			i++
			closeBracket := nextValue(tokens, i+1, 1) + 1
			tokens = insert(tokens, closeBracket, ")")
		}
		i++
	}
	return tokens
}

func insert(tokens []string, index int, token string) []string {
	newArr := make([]string, len(tokens)+1)
	copy(newArr, tokens[:index])
	newArr[index] = token
	copy(newArr[index+1:], tokens[index:])
	return newArr
}

func nextValue(tokens []string, start, increment int) int {
	layers := 0
	for i := start; i >= 0 && i < len(tokens); i += increment {
		if tokens[i] == "(" {
			layers++
		} else if tokens[i] == ")" {
			layers--
		}

		if layers == 0 {
			return i
		}
	}
	return -1
}

func evaluate(tokens []string) int {
	valStack := make([]int, 1)
	valStack[0] = 0

	opStack := make([]string, 1)
	opStack[0] = "+"

	for _, token := range tokens {
		switch token {
		case "(":
			valStack = append(valStack, 0)
			opStack = append(opStack, "+")
		case ")":
			value := valStack[len(valStack)-1]
			valStack = valStack[:len(valStack)-1]
			opStack = opStack[:len(opStack)-1]

			switch opStack[len(opStack)-1] {
			case "+":
				valStack[len(valStack)-1] += value
			case "*":
				valStack[len(valStack)-1] *= value
			default:
				panic("INVALID OPERATOR!")
			}

			opStack[len(opStack)-1] = ""

		case "+":
			opStack[len(opStack)-1] = "+"
		case "*":
			opStack[len(opStack)-1] = "*"
		default:
			value, _ := strconv.Atoi(token)

			switch opStack[len(opStack)-1] {
			case "+":
				valStack[len(valStack)-1] += value
			case "*":
				valStack[len(valStack)-1] *= value
			default:
				panic("INVALID OPERATOR!")
			}

			opStack[len(opStack)-1] = ""
		}

	}

	return valStack[0]

}

func tokenise(expression string) []string {
	return strings.Split(expression, " ")
}
