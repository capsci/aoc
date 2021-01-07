package main

import "fmt"

const turnDe = "Input : %v Turn %d : Spoken number is %d\n"

func main() {
	gameDataSeed := []int{5, 2, 8, 16, 18, 0, 1}
	// gameDataSeed := []int{0, 3, 6} // Sample
	rambunc2020(gameDataSeed, 2020)
	rambunc2020(gameDataSeed, 30000000)
}

func rambunc2020(numbers []int, numTurns int) {
	memory := make(map[int]int)
	var lastNum int
	turn := 1
	for _, lastNum = range numbers {
		memory[lastNum] = turn
		turn++
	}
	// First num will always be spoken for the first time; hence
	turn++
	lastNum = 0

	for turn <= numTurns {
		speak := 0
		lastSpoken, spoken := memory[lastNum]
		if spoken {
			speak = turn - 1 - lastSpoken
		}
		//		fmt.Printf("Turn %4d: Last-%4d, Speak-%4d\n", turn, lastNum, speak)
		memory[lastNum] = turn - 1

		lastNum = speak
		turn++
	}
	fmt.Printf(turnDe, numbers, turn-1, lastNum)
}
