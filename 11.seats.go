package main

import (
	"fmt"

	"github.com/capsci/aoc/helper"
)

const seatsDataFile = "11.data.txt"
const seatsSampleDataFile = "11.data.sample.txt"

const equiSeat = "No of occupied seats on equilibrium is %d\n"
const equiSeatNew = "No of occupied seats on new equilibrium is %d\n"

const empty = byte(76)    //L
const occupied = byte(35) //#
const noSeat = byte(46)   //.

func main() {
	seats := readSeatsData(seatsDataFile)
	fmt.Printf(equiSeat, equilibriumSeats(seats, occupiedNeighbors, 4))
	fmt.Printf(equiSeatNew, equilibriumSeats(seats, occupiedFirstNeighbor, 5))
}

func equilibriumSeats(seats []string, countNeighbors func(int, int, []string) int, allowedNeighbors int) (occu int) {
	var changed bool
	for {
		seats, changed = simulateSeating(seats, countNeighbors, allowedNeighbors)
		if !changed {
			break
		}
	}
	occu, _, _ = countSeats(seats)
	return occu
}

func simulateSeating(seats []string, countNeighbors func(int, int, []string) int, allowedNeighbors int) (newSeats []string, changed bool) {
	newSeats = make([]string, len(seats))
	copy(newSeats, seats)
	for r := 0; r < len(seats); r++ {
		seatBytes := []byte(newSeats[r])
		for c := 0; c < len(seats[r]); c++ {
			n := countNeighbors(r, c, seats)
			switch seats[r][c] {
			case occupied:
				if n >= allowedNeighbors {
					changed = true
					seatBytes[c] = empty
				}
			case empty:
				if n == 0 {
					changed = true
					seatBytes[c] = occupied
				}
			}
		}
		newSeats[r] = string(seatBytes)
	}
	return
}

func countSeats(seats []string) (cOccupied, cEmpty, cNoSeat int) {
	for r := 0; r < len(seats); r++ {
		for c := 0; c < len(seats[r]); c++ {
			switch seats[r][c] {
			case occupied:
				cOccupied++
			case empty:
				cEmpty++
			case noSeat:
				cNoSeat++
			}
		}
	}
	return
}

func occupiedNeighbors(row, col int, seats []string) (neighbors int) {
	for r := row - 1; r <= row+1; r++ {
		if r < 0 || r >= len(seats) {
			// OutOfBonds
			continue
		}
		for c := col - 1; c <= col+1; c++ {
			if c < 0 || c >= len(seats[r]) {
				// OutOfBonds
				continue
			}
			if r == row && c == col {
				// Self
				continue
			}
			if []byte(seats[r])[c] == occupied {
				neighbors++
			}
		}
	}
	return
}

// Looks at first seat (skips all non-seats)
func occupiedFirstNeighbor(row, col int, seats []string) (neighbors int) {
	for r := -1; r <= 1; r++ {
		for c := -1; c <= 1; c++ {
			if r == 0 && c == 0 {
				// Self
				continue
			}
			level := 1
			for {
				currR := row + level*r
				currC := col + level*c
				if currR < 0 || currR >= len(seats) ||
					currC < 0 || currC >= len(seats[currR]) {
					// Out of bounds
					break
				}
				if seats[currR][currC] == noSeat {
					level++
					continue
				}
				if seats[currR][currC] == occupied {
					neighbors++
				}
				break
			}
		}
	}
	return
}

func readSeatsData(fileName string) (seats []string) {
	ch := make(chan string)
	go helper.StartReading(fileName, ch)
	for {
		val, ok := <-ch
		if !ok {
			break
		}
		seats = append(seats, val)
	}
	return
}
