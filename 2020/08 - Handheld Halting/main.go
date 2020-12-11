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

	vm := &virtualMachine{instructions, 0, 0}

	fmt.Println(partOne(vm))
	fmt.Println(partTwo(vm))
}

func partOne(vm *virtualMachine) int {
	value, _ := vm.Run()
	return value
}

func partTwo(vm *virtualMachine) int {
	for i := 0; i < len(vm.instructions); i++ {
		if vm.instructions[i].opcode == nop {
			vm.instructions[i].opcode = jmp
			value, success := vm.Run()
			if success {
				return value
			}
			vm.instructions[i].opcode = nop
		} else if vm.instructions[i].opcode == jmp {
			vm.instructions[i].opcode = nop
			value, success := vm.Run()
			if success {
				return value
			}
			vm.instructions[i].opcode = jmp
		}
	}
	return -1
}

type virtualMachine struct {
	instructions []instruction
	pc           int
	acc          int
}

func (vm *virtualMachine) Next() {
	current := vm.instructions[vm.pc]
	switch operand := current.operand; current.opcode {
	case nop:
		vm.pc++
	case acc:
		vm.acc += operand
		vm.pc++
	case jmp:
		vm.pc += operand
	}
}

func (vm *virtualMachine) Run() (int, bool) {
	vm.pc = 0
	vm.acc = 0
	visited := make([]bool, len(vm.instructions))
	for vm.pc < len(visited) && !visited[vm.pc] {
		visited[vm.pc] = true
		vm.Next()
	}
	return vm.acc, vm.pc == len(visited)
}

func getInstructions(lines []string) []instruction {
	operations := make([]instruction, len(lines))

	for i, line := range lines {
		parts := strings.Split(line, " ")
		switch parts[0] {
		case "nop":
			operations[i].opcode = nop
		case "acc":
			operations[i].opcode = acc
		case "jmp":
			operations[i].opcode = jmp
		}
		operations[i].operand, _ = strconv.Atoi(parts[1])
	}

	return operations
}

const (
	nop = iota
	acc
	jmp
)

type instruction struct {
	opcode  int
	operand int
}
