package main

import (
	"fmt"
	"strconv"

	"github.com/capsci/aoc/helper"
)

const xmasDataFile = "9.data.txt"
const xmasDataPreamble = 25
const xmasSampleDataFile = "9.data.sample.txt"
const xmasSampleDataPreamble = 5

const invalidNumber = "Invalid number in the list is: %d\n"
const weaknessSum = "Weakness sum of invalidNumber is: %d\n"

func main() {
	numbers := readXmasData(xmasDataFile)
	invalid := findInvalid(numbers, xmasDataPreamble)
	fmt.Printf(invalidNumber, invalid)
	fmt.Printf(weaknessSum, findContiguousSum(numbers, invalid))
}

type Queue struct {
	Items   []int
	Present map[int]bool
	Size    int
}

func (q *Queue) Add(val int) {
	if len(q.Items) > q.Size {
		delete(q.Present, q.Items[0])
		q.Items = append(q.Items[1:], val)
	} else {
		q.Items = append(q.Items, val)
	}
	q.Present[val] = true
}

func (q *Queue) Valid(val int) bool {
	if len(q.Items) < q.Size {
		return true
	}
	for k, _ := range q.Present {
		_, exists := q.Present[val-k]
		if exists {
			return true
		}
	}
	return false
}

func readXmasData(fileName string) (numbers []int) {
	ch := make(chan string)
	go helper.StartReading(fileName, ch)
	for {
		line, ok := <-ch
		if !ok {
			break
		}
		val, err := strconv.Atoi(line)
		helper.CheckErr(err)
		numbers = append(numbers, val)
	}
	return
}

func findInvalid(numbers []int, preambleSize int) int {
	q := &Queue{
		Size:    preambleSize,
		Present: make(map[int]bool),
	}
	for _, num := range numbers {
		if !q.Valid(num) {
			return num
		}
		q.Add(num)
	}
	return -1
}

func findContiguousSum(numbers []int, target int) int {
	var window []int
	var windowSum int
	for _, number := range numbers {
		windowSum += number
		window = append(window, number)
		for {
			if windowSum <= target {
				break
			}
			windowSum -= window[0]
			window = window[1:]
		}
		//		fmt.Println(windowSum, ":", window)
		if windowSum == target {
			return sumLargSmall(window)
		}
	}
	return -1
}

func sumLargSmall(numbers []int) int {
	min, max := numbers[0], numbers[0]
	for _, num := range numbers {
		if num < min {
			min = num
		} else if num > max {
			max = num
		}
	}
	return min + max
}

/*
Incorrect:
11
*/
