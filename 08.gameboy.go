package main

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/capsci/aoc/helper"
)

const gameboyDataFile = "08.data.txt"

const beforeLoop = "Accumulator value before looping starts is %d\n"
const afterFix = "Accumulator value after fixing gameboy is %d\n"

func main() {
	instructions := readGameBoyData(gameboyDataFile)
	val, _ := perform(0, 0, instructions)
	fmt.Printf(beforeLoop, val)
	instructions = readGameBoyData(gameboyDataFile)
	val = part2(instructions)
	fmt.Printf(afterFix, val)
}

func perform(idx, acc int, instructions []string) (int, bool) {
	if idx >= len(instructions) {
		return -1, false
	}
	if idx == len(instructions)-1 {
		return acc, true
	}
	if len(instructions[idx]) == 0 {
		return acc, false
	}

	re := regexp.MustCompile(`(\w{3}) ([+|-]\d+)`)
	matches := re.FindStringSubmatch(instructions[idx])
	command := matches[1]
	amount, err := strconv.Atoi(matches[2])
	helper.CheckErr(err)

	// Unset so looping operation can be detected
	instructions[idx] = ""

	switch command {
	case "nop":
		return perform(idx+1, acc, instructions)
	case "jmp":
		return perform(idx+amount, acc, instructions)
	case "acc":
		return perform(idx+1, acc+amount, instructions)
	}

	return acc, true
}

func readGameBoyData(fileName string) (instructions []string) {
	ch := make(chan string)
	go helper.StartReading(fileName, ch)
	for {
		line, ok := <-ch
		if !ok {
			break
		}
		instructions = append(instructions, line)
	}
	return
}

func part2(instructions []string) int {
	var acc, idx int
	re := regexp.MustCompile(`(\w{3}) ([+|-]\d+)`)

	for {
		matches := re.FindStringSubmatch(instructions[idx])
		command := matches[1]
		amount, err := strconv.Atoi(matches[2])
		helper.CheckErr(err)

		instructions[idx] = ""
		if command == "nop" {
			newIns := make([]string, len(instructions))
			copy(newIns, instructions)
			// perform jump and check
			newAcc, compl := perform(idx+amount, acc, newIns)
			if compl {
				return newAcc
			}
			// continue with nop
			idx++
		} else if command == "jmp" {
			newIns := make([]string, len(instructions))
			copy(newIns, instructions)
			// perform nop and check
			newAcc, compl := perform(idx+1, acc, newIns)
			if compl {
				return newAcc
			}
			// continue with jmp
			idx += amount
		} else if command == "acc" {
			acc += amount
			idx++
		}
		if idx >= len(instructions)-1 {
			fmt.Println("Hit ", idx, " with accumulator value ", acc)
			return -1
		}
	}
}
