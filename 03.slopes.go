package main

import (
	"fmt"

	"github.com/capsci/aoc/helper"
)

const slopeDataSampleFile = "3.data.sample.txt"
const slopeDataFile = "3.data.txt"

const encountered = "Trees encountered in %dr%dd move : %d\n"

func main() {
	t2 := slopes1(slopeDataFile, 3, 1)
	fmt.Printf(encountered, 3, 1, t2)
	tt := slopes2(slopeDataFile)
	fmt.Printf(encountered, -1, -1, tt*t2)
}

func slopes2(slopeDataFile string) int {
	t1 := slopes1(slopeDataFile, 1, 1)
	//	fmt.Printf(encountered, 1, 1, t1)
	t3 := slopes1(slopeDataFile, 5, 1)
	//	fmt.Printf(encountered, 5, 1, t3)
	t4 := slopes1(slopeDataFile, 7, 1)
	//	fmt.Printf(encountered, 7, 1, t4)
	t5 := slopes1(slopeDataFile, 1, 2)
	//	fmt.Printf(encountered, 1, 2, t5)
	return t1 * t3 * t4 * t5

}

func slopes1(fileName string, right, down int) (trees int) {
	col := 0
	row := 0
	ch := make(chan string)
	go helper.StartReading(slopeDataFile, ch)
	for {
		level, ok := <-ch
		if !ok {
			break
		}
		// Skip first level
		if row == 0 {
			row++
			continue
		}
		// Check only for valid down level moves
		if row%down != 0 {
			row++
			continue
		}

		col = (col + right) % len(level)
		row++
		if []byte(level)[col] == []byte("#")[0] {
			trees++
		}

	}
	return
}
