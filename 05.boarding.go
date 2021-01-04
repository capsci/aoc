package main

import (
	"fmt"
	"sort"

	"github.com/capsci/aoc/helper"
)

const boardingDataFile = "5.data.txt"

const highestSeatIDMsg = "Highest Seat ID on borading passes is : %d\n"
const vacantSeatIDMsg = "Vacant Seat ID is : %d\n"
const notFoundMsg = "%s not foind for ticket %s [%d,%d]"

func main() {
	tickets := readBoardingData(boardingDataFile)
	var seatIDs []int
	var highestSeatID int
	for _, ticket := range tickets {
		sid := seatID(ticket)
		if sid > highestSeatID {
			highestSeatID = sid
		}
		seatIDs = append(seatIDs, sid)
	}

	fmt.Printf(highestSeatIDMsg, highestSeatID)
	fmt.Printf(vacantSeatIDMsg, findVacantSeat(seatIDs))
}

func findVacantSeat(list []int) int {
	sort.Ints(list)
	// sum of all tickets should be
	sum := sumToN(list[len(list)-1]) - sumToN(list[0]-1)
	// actual sum of tickets
	var actual int
	for _, elem := range list {
		actual += elem
	}
	// Missing seat
	return sum - actual
}

func sumToN(n int) int {
	fmt.Println(n, ":", n*(n+1)/2)
	return n * (n + 1) / 2
}

func seatID(ticket string) int {
	var row, col int
	min, max := 0, 127
	for _, r := range ticket[:7] {
		//		fmt.Println(ticket[:7], ":", string(r), " ", min, "-", max)
		if string(r) == "F" {
			min, max = partition(min, max, false)
		} else {
			min, max = partition(min, max, true)
		}
	}
	if min != max {
		panic(fmt.Sprintf(notFoundMsg, "row", ticket, min, max))

	}
	row = min

	min, max = 0, 7
	for _, c := range ticket[7:] {
		if string(c) == "L" {
			min, max = partition(min, max, false)
		} else {
			min, max = partition(min, max, true)
		}
	}
	if min != max {
		panic(fmt.Sprintf(notFoundMsg, "col", ticket, min, max))
	}
	col = min

	return row*8 + col
}

func partition(min, max int, upper bool) (int, int) {
	mid := (max + 1 - min) / 2
	if upper {
		return min + mid, max
	}
	return min, max - mid
}

func readBoardingData(fileName string) (tickets []string) {
	ch := make(chan string)
	go helper.StartReading(fileName, ch)
	for {
		ticket, ok := <-ch
		if !ok {
			break
		}
		tickets = append(tickets, ticket)
	}
	return
}
