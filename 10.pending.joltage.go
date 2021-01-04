package main

import (
	"fmt"
	"sort"
	"strconv"

	"github.com/capsci/aoc/helper"
)

const joltageDataFile = "10.data.txt"
const joltageSampleDataFile = "10.data.sample.txt"
const joltageSample2DataFile = "10.data.sample.2.txt"

const nextStepLimit = 3

const jolt1jolt3Product = "Product of 1jolt different jolts and 3jolt different jolts is : %d\n"
const connectionWays = "Different ways we can connect device to outlet is %d\n"

var memo map[int]int

func main() {
	joltage := readJoltageData(joltageDataFile)
	sort.Ints(joltage)
	fmt.Printf(jolt1jolt3Product, joltageChain(joltage))
	memo = make(map[int]int)
	fmt.Printf(connectionWays, findCombination(0, joltage))
	memo = make(map[int]int)
	fmt.Printf(connectionWays, findComb(0, joltage))
}

func joltageChain(joltage []int) int {
	var jolt1 int
	var jolt3 int

	prevJolt := joltage[0]
	for _, jolt := range joltage {
		if jolt-prevJolt == 1 {
			jolt1++
		} else if jolt-prevJolt == 3 {
			jolt3++
		}
		prevJolt = jolt
	}
	// 1 added to jolt1 to account aircraft outlet adapter
	// 1 added to jolt3 to account built-in adapter
	return (jolt1 + 1) * (jolt3 + 1)
}

func readJoltageData(fileName string) (joltage []int) {

	ch := make(chan string)
	go helper.StartReading(fileName, ch)
	for {
		line, ok := <-ch
		if !ok {
			break
		}
		jolt, err := strconv.Atoi(line)
		helper.CheckErr(err)
		joltage = append(joltage, jolt)
	}
	return
}

func findComb(idx int, adapters []int) (ways int) {
	maxIdx := len(adapters) - 1
	if idx == maxIdx {
		return 1
	}
	if v, ok := memo[adapters[idx]]; ok {
		return v
	}

	if idx+1 <= maxIdx && adapters[idx+1]-adapters[idx] <= nextStepLimit {
		ways += findComb(idx+1, adapters)
	}
	if idx+2 <= maxIdx && adapters[idx+2]-adapters[idx] <= nextStepLimit {
		ways += findComb(idx+2, adapters)
	}
	if idx+3 <= maxIdx && adapters[idx+3]-adapters[idx] <= nextStepLimit {
		ways += findComb(idx+3, adapters)
	}
	memo[adapters[idx]] = ways
	return
}

func findCombination(idx int, adapters []int) (ways int) {
	if idx == len(adapters)-1 {
		return 1
	}

	for i := idx + 1; i < len(adapters) && i < idx+4; i++ {
		nextStep := adapters[i] - adapters[idx]
		if nextStep <= nextStepLimit {
			//			ways += findCombination(i, adapters)
			v, present := memo[adapters[i]]
			if !present {
				memo[adapters[i]] = findCombination(i, adapters)
				ways += memo[adapters[i]]
				//				ways += findCombination(i, adapters)
			} else {
				ways += v
			}
		}
	}
	memo[adapters[idx]] = ways
	//	fmt.Println(idx, adapters[idx], memo)
	return ways
}

/* Incorrect:
49607173328384
*/
