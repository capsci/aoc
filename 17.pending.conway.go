package main

import (
	"fmt"
	"strings"

	"github.com/capsci/aoc/helper"
)

const conwayDataSampleFile = "17.data.sample.txt"

func main() {
	initial := readConwayData(conwayDataSampleFile)
	world := initWorld(initial, 6)

	printWorld(world)
	world = turn(world)
	printWorld(world)
}

func turn(start [][][]string) (final [][][]string) {
	final = helper.GiveEmpty3DArray(len(start[0]), len(start[0][0]), len(start))
	copy(final, start)

	for z := 0; z < len(start); z++ {
		for x := 0; x < len(start[z]); x++ {
			for y := 0; y < len(start[z][x]); y++ {
				count := activeNeighbors(x, y, z, start)
				active := start[z][x][y] == "#"

				if active && (count > 3 || count < 2) {
					fmt.Println("Dectivating", z, x, y, count)
					final[z][x][y] = "."
				} else if !active && count == 3 {
					fmt.Println("Activating", z, x, y, count)
					final[z][x][y] = "#"
				}
			}
		}
	}

	//	printWorld(final)
	return
}

func initWorld(initial [][]string, turns int) (world [][][]string) {
	// z*x*y array
	// world(in size) after 6 turns
	world = helper.Give3DArray(2*turns+len(initial), 2*turns+len(initial[0]), turns*2+1, ".")
	// initial will be at z:turns, x:turns-len(initial), y:turns-len(initial[0])
	for x, _ := range initial {
		for y, _ := range initial[x] {
			world[turns][x+turns+1-len(initial)/2][y+turns+1-len(initial[x])/2] = initial[x][y]
		}
	}
	return
}

func printWorld(world [][][]string) {
	for z := 0; z < len(world); z++ {
		fmt.Println("z = ", z-len(world)/2)
		for x := 0; x < len(world[z]); x++ {
			fmt.Println(world[z][x])
		}
	}
}

func activeNeighbors(x, y, z int, world [][][]string) (active int) {
	for zd := z - 1; zd >= 0 && zd < len(world); zd++ {
		for xd := x - 1; xd >= 0 && xd < len(world[zd]); xd++ {
			for yd := y - 1; yd >= 0 && yd < len(world[zd][xd]); yd++ {
				if x == xd && y == yd && z == zd {
					continue
				}
				if world[zd][xd][yd] == "#" {
					active++
				}
			}
		}
	}
	return
}

func readConwayData(fileName string) (initial [][]string) {
	ch := make(chan string)
	go helper.StartReading(fileName, ch)

	for {
		val, ok := <-ch
		if !ok {
			break
		}
		initial = append(initial, strings.Split(val, ""))
	}
	return
}
