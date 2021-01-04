package main

import (
	"fmt"
	"strconv"

	"github.com/capsci/aoc/helper"
)

const expensePair = "Found a pair (%d,%d) with sum %d and product %d\n"
const expenseTriplet = "Found a triplet (%d,%d,%d) with sum %d and product %d\n"

const input = "1.data.txt"

func main() {
	expenses := readElvesData(input)
	total := 2020
	a1, b1, found := elves1(expenses, total)
	if found {
		fmt.Printf(expensePair, a1, b1, a1+b1, a1*b1)
	}
	a2, b2, c2, found := elves2(expenses, total)
	if found {
		fmt.Printf(expenseTriplet, a2, b2, c2, a2+b2+c2, a2*b2*c2)
	}
}

func elves2(expenses []int, total int) (first, second, third int, present bool) {
	for _, third = range expenses {
		first, second, present = elves1(expenses, total-third)
		if present {
			return
		}
	}
	return
}

func elves1(expenses []int, total int) (first, second int, present bool) {
	ledger := make(map[int]bool)
	for _, first = range expenses {
		second = total - first
		_, present = ledger[second]
		if present {
			return
		}
		// if not present; add
		ledger[first] = true
	}
	return
}

func readElvesData(fileName string) (expenses []int) {
	ch := make(chan string)
	go helper.StartReading(fileName, ch)
	for {
		val, ok := <-ch
		if !ok {
			break
		}
		exp, err := strconv.Atoi(val)
		helper.CheckErr(err)
		expenses = append(expenses, exp)
	}

	return expenses
}
