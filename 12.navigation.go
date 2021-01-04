package main

import (
	"fmt"
	"math"
	"regexp"
	"strconv"

	"github.com/capsci/aoc/helper"
)

const naviagtionSampleDataFile = "12.data.sample.txt"
const naviagtionDataFile = "12.data.txt"

const manhattanDistance = "Manhattan Distance of ship from starting point is %d\n"
const manhattanWaypointDistance = "Manhattan Distance of ship from starting point is when using Waypoint navigation %d\n"

func main() {
	readNavigationData(naviagtionDataFile)
}

func readNavigationData(fileName string) {
	ch := make(chan string)
	go helper.StartReading(fileName, ch)

	re := regexp.MustCompile(`^(\w)(\d+)`)
	s := &Ship{}
	s.PointEast()
	w := &WayPoint{}
	w.SetWayPoint(10, 1)
	for {
		line, ok := <-ch
		if ok != true {
			break
		}
		matches := re.FindStringSubmatch(line)
		command := matches[1]
		amount, err := strconv.ParseFloat(matches[2], 64)
		helper.CheckErr(err)
		switch command {
		case "N":
			s.MoveNorth(amount)
			w.MoveNorth(amount)
		case "S":
			s.MoveSouth(amount)
			w.MoveSouth(amount)
		case "E":
			s.MoveEast(amount)
			w.MoveEast(amount)
		case "W":
			s.MoveWest(amount)
			w.MoveWest(amount)
		case "L":
			s.SteerLeft(amount)
			w.MoveLeft(amount)
		case "R":
			s.SteerRight(amount)
			w.MoveRight(amount)
		case "F":
			s.MoveForward(amount)
			w.MoveForward(amount)
		default:
			panic("Found " + command)
		}
		//		fmt.Println(line, w.GetLocation())
	}
	e, n := s.Location()
	fmt.Printf(manhattanDistance, int(math.Abs(e)+math.Abs(n)))
	e, n = w.Location()
	fmt.Printf(manhattanDistance, int(math.Abs(e)+math.Abs(n)))
}

type Ship struct {
	direction float64 // 0 points to North
	east      float64 // distance travelled in East
	north     float64 // distance travelled in North
}

func (s *Ship) PointEast() {
	s.direction = math.Pi / 2
}

func (s *Ship) SteerLeft(degrees float64) {
	radians := degrees / 90
	s.direction -= (math.Pi * radians / 2)
}

func (s *Ship) SteerRight(degrees float64) {
	radians := degrees / 90
	s.direction += (math.Pi * radians / 2)
}

func (s *Ship) MoveNorth(units float64) {
	s.north += units
}

func (s *Ship) MoveSouth(units float64) {
	s.north -= units
}

func (s *Ship) MoveEast(units float64) {
	s.east += units
}

func (s *Ship) MoveWest(units float64) {
	s.east -= units
}

func (s *Ship) MoveForward(units float64) {
	e, n := math.Sincos(s.direction)
	s.MoveEast(math.Round(e) * units)
	s.MoveNorth(math.Round(n) * units)
}

func (s *Ship) Location() (float64, float64) {
	return s.east, s.north
}

type WayPoint struct {
	east  float64 // waypoint location
	north float64 // waypoint location
	shipE float64 // ship location
	shipN float64 // ship location
}

func (w *WayPoint) SetWayPoint(east, north float64) {
	w.east = east
	w.north = north
}

func (w *WayPoint) MoveLeft(degrees float64) {
	times := degrees / 90
	for {
		if times <= 0 {
			break
		}
		//		w.east, w.north = w.north, -w.east
		w.north, w.east = w.east, -w.north
		times--
	}
}

func (w *WayPoint) MoveRight(degrees float64) {
	times := degrees / 90
	for {
		if times <= 0 {
			break
		}
		//		w.east, w.north = -w.north, w.east
		w.north, w.east = -w.east, w.north
		times--
	}
}

func (w *WayPoint) MoveNorth(units float64) {
	w.north += units
}

func (w *WayPoint) MoveSouth(units float64) {
	w.north -= units
}

func (w *WayPoint) MoveEast(units float64) {
	w.east += units
}

func (w *WayPoint) MoveWest(units float64) {
	w.east -= units
}

func (w *WayPoint) MoveForward(units float64) {
	w.shipE += (w.east * units)
	w.shipN += (w.north * units)
}

func (w *WayPoint) GetLocation() string {
	return fmt.Sprintf("Ship(%fE,%fN) WayPoint(%fE,%fN)", w.shipE,
		w.shipN, w.east, w.north)
}

func (w *WayPoint) Location() (float64, float64) {
	return w.shipE, w.shipN
}
