package main

import "fmt"

const turnDe = "T %d : Last Num was %d, memory: %v\n"

func main() {
	//	gameDataSeed := []int{5, 2, 8, 16, 18, 0, 1}
	gameDataSeed := []int{0, 3, 6} // Sample
	rambunc2020(gameDataSeed, 10)
}

func rambunc2020(numbers []int, numTurns int) {
	memory := make(map[int]int)
	var i, lastNum int
	for i, lastNum = range numbers {
		memory[lastNum] = i + 1
	}
	// First num will always be spoken for the forst time
	i = len(memory) + 1
	memory[0] = i
	lastNum = 0
	fmt.Println(memory)
	for {
		fmt.Printf(turnDe, i, lastNum, memory)
		i++
		lastSeen, present := memory[lastNum]
		if !present {
			memory[lastNum] = i
			lastNum = 0
		} else {
			fmt.Println(lastNum, ":", lastSeen, "t", i)
			diff := i - lastSeen
			memory[diff] = i
			lastNum = diff
		}
		if i >= numTurns {
			break
		}

	}
}
