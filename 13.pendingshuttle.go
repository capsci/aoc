package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/capsci/aoc/helper"
)

const shuttleDataFile = "13.data.txt"
const shuttleSampleDataFile = "13.data.sample.txt"

const minWaitTimeBus = "Need to wait %d minutes to catch Bus%d (Answer: %d)\n"
const contestBus = "Timestamp after which every bus leaves in a minute : %d\n"

func main() {
	dep, buses := readShuttleData(shuttleDataFile)
	minWaitTime := buses[0] - int(math.Mod(float64(dep), float64(buses[0])))
	minWaitBus := buses[0]
	// part1
	for _, bus := range buses {
		if bus == 0 {
			// x
			continue
		}
		waitTime := bus - int(math.Mod(float64(dep), float64(bus)))
		if waitTime < minWaitTime {
			minWaitTime = waitTime
			minWaitBus = bus
		}
	}
	fmt.Printf(minWaitTimeBus, minWaitTime, minWaitBus, minWaitTime*minWaitBus)
	// part2
	fmt.Printf(contestBus, contest(buses))

}

func readShuttleData(fileName string) (departure int, buses []int) {
	var err error

	lines := helper.ReadLines(fileName)

	departure, err = strconv.Atoi(lines[0])
	helper.CheckErr(err)
	busTimes := strings.Split(lines[1], ",")
	for _, busTime := range busTimes {
		if busTime == "x" {
			busTime = "0"
		}
		busID, err := strconv.Atoi(busTime)
		helper.CheckErr(err)
		buses = append(buses, busID)
	}
	return

}

// Brute Force
func contest(buses []int) int64 {
	timestamp := int64(buses[0])
	for {
		valid := true
		fmt.Println("Checking ", timestamp)
		for i := 1; i < len(buses); i++ {
			if buses[i] == 0 { // 'x' buses
				continue
			}
			if (timestamp+int64(i))%int64(buses[i]) != 0 {
				valid = false
				break
			}
		}
		if valid {
			return timestamp
		}
		timestamp += int64(buses[0])
	}
	return 0
}
