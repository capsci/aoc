package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/capsci/aoc/helper"
)

const ticketDataSampleFile = "16.data.sample.txt"
const ticketDataFile = "16.data.txt"

var sampleIntervals = []Interval{
	Interval{Start: 1, End: 3},
	Interval{Start: 5, End: 7},
	Interval{Start: 6, End: 11},
	Interval{Start: 33, End: 44},
	Interval{Start: 13, End: 40},
	Interval{Start: 45, End: 50},
}

var inputIntervals = []Interval{
	Interval{Start: 29, End: 917},
	Interval{Start: 943, End: 952}, // departure location
	Interval{Start: 50, End: 875},
	Interval{Start: 884, End: 954}, // departure station
	Interval{Start: 41, End: 493},
	Interval{Start: 503, End: 949}, // departure platform
	Interval{Start: 50, End: 867},
	Interval{Start: 875, End: 966}, // departure track
	Interval{Start: 30, End: 655},
	Interval{Start: 679, End: 956}, // departure date
	Interval{Start: 46, End: 147},
	Interval{Start: 153, End: 958}, // departure time
	Interval{Start: 50, End: 329},
	Interval{Start: 344, End: 968}, // arrival location
	Interval{Start: 42, End: 614},
	Interval{Start: 623, End: 949}, // arrival station
	Interval{Start: 35, End: 849},
	Interval{Start: 860, End: 973}, // arrival platform
	Interval{Start: 42, End: 202},
	Interval{Start: 214, End: 959}, // arrival track
	Interval{Start: 38, End: 317},
	Interval{Start: 329, End: 968}, // class
	Interval{Start: 44, End: 530},
	Interval{Start: 539, End: 953}, // duration
	Interval{Start: 28, End: 713},
	Interval{Start: 727, End: 957}, // price
	Interval{Start: 30, End: 157},
	Interval{Start: 179, End: 966}, // route
	Interval{Start: 38, End: 114},
	Interval{Start: 136, End: 969}, // row
	Interval{Start: 45, End: 441},
	Interval{Start: 465, End: 956}, // seat
	Interval{Start: 44, End: 799},
	Interval{Start: 824, End: 951}, // train
	Interval{Start: 41, End: 411},
	Interval{Start: 437, End: 953}, // type
	Interval{Start: 39, End: 79},
	Interval{Start: 86, End: 969}, // wagon
	Interval{Start: 48, End: 306},
	Interval{Start: 317, End: 974}, // zone
}

func main() {
	//	inputFile, inputIntervals := ticketDataSampleFile, sampleIntervals
	inputFile, inputIntervals := ticketDataFile, inputIntervals

	fmt.Println(ticketScanningError(readTicketData(inputFile), mergeOverlappingIntervals(inputIntervals)))
}

func ticketScanningError(tickets [][]float64, intervals []Interval) (errorRate float64) {
	for _, ticket := range tickets {
		for _, entry := range ticket {
			validEntry := false
			for _, interval := range intervals {
				if entry >= interval.Start && entry <= interval.End {
					validEntry = true
					break
				}
			}
			if !validEntry {
				errorRate += entry
			}
		}
	}
	return
}

func readTicketData(fileName string) (tickets [][]float64) {
	ch := make(chan string)
	go helper.StartReading(fileName, ch)
	for {
		val, ok := <-ch
		if !ok {
			break
		}
		list := strings.Split(val, ",")
		var ticket []float64
		for _, item := range list {
			i, err := strconv.ParseFloat(item, 64)
			helper.CheckErr(err)
			ticket = append(ticket, i)
		}
		tickets = append(tickets, ticket)
	}
	return
}

type Interval struct {
	Start, End float64
}

func mergeOverlappingIntervals(intervals []Interval) (merged []Interval) {
	merged = append(merged, intervals[0])
	for _, interval := range intervals[1:] {
		isMerged := false
		for i := 0; i < len(merged); i++ {
			saved := &merged[i]
			if interval.Start == saved.Start ||
				(interval.Start >= saved.Start && interval.Start <= saved.End) {
				// start range overlaps
				saved.End = math.Max(interval.End, saved.End)
				isMerged = true
			} else if interval.End == saved.End ||
				(interval.End >= saved.Start && interval.End <= saved.End) {
				// end range overlaps
				saved.Start = math.Min(interval.Start, saved.Start)
				isMerged = true
			} else if (interval.Start <= saved.Start && interval.End >= saved.End) ||
				(saved.Start <= interval.Start && saved.End >= interval.End) {
				// One range is included in other
				saved.Start = math.Min(interval.Start, saved.Start)
				saved.End = math.Max(interval.End, saved.End)
				isMerged = true
			}
		}
		if !isMerged {
			merged = append(merged, interval)
		}
	}
	return
}
