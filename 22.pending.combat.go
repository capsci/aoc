package main

import (
	"fmt"

	"github.com/capsci/aoc/helper"
)

const roundWin = "Player %s wins (%d against %d)\n"
const gameWin = "Winner score is %d\n"

type CombatData struct {
	Player1 []int
	Player2 []int
}

var combatSampleData = CombatData{
	Player1: []int{9, 2, 6, 3, 1},
	Player2: []int{5, 8, 4, 7, 10},
}

var combatData = CombatData{
	Player1: []int{24, 22, 26, 6, 14, 19, 27, 17, 39, 34, 40, 41, 23, 30, 36, 11, 28, 3, 10, 21, 9, 50, 32, 25, 8},
	Player2: []int{48, 49, 47, 15, 42, 44, 5, 4, 13, 7, 20, 43, 12, 37, 29, 18, 45, 16, 1, 46, 38, 35, 2, 33, 31},
}

func main() {
	part1(combatData)
}

func part1(data CombatData) {
	var p1, p2 helper.Queue
	for _, card := range data.Player1 {
		p1.Enqueue(card)
	}
	for _, card := range data.Player2 {
		p2.Enqueue(card)
	}

	for !p1.Empty() && !p2.Empty() {
		// P1 wins
		p1t := p1.Dequeue().(int)
		p2t := p2.Dequeue().(int)
		if p1t > p2t {
			p1.Enqueue(p1t)
			p1.Enqueue(p2t)
			fmt.Printf(roundWin, "P1", p1t, p2t)
		} else {
			p2.Enqueue(p2t)
			p2.Enqueue(p1t)
			fmt.Printf(roundWin, "P2", p2t, p1t)
		}
	}

	var winner helper.Queue
	winner = p1
	if p2.Size() > p1.Size() {
		winner = p2
	}
	score := 0
	for m := winner.Size(); m > 0; m-- {
		score += m * winner.Dequeue().(int)
	}
	fmt.Printf(gameWin, score)

}
