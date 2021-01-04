package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/capsci/aoc/helper"
)

const maskingDataFile = "14.data.txt"
const maskingSampleDataFile = "14.data.sample.txt"
const maskingSample2DataFile = "14.data.sample.2.txt"

const sumPostInit = "Sum of values in memory after initialization is %d\n"
const sumPostInit2 = "Sum of values in memory after initialization using decoder v2 is %d\n"

func main() {
	instructions := readMaskingData(maskingDataFile)
	fmt.Printf(sumPostInit, bitmask(instructions))
	fmt.Printf(sumPostInit2, memMask(instructions))
}

// part 1
func bitmask(instructions []string) (sum int) {
	memory := make(map[int]int)
	reMask := regexp.MustCompile(`mask = ([0|1|X]+)`)
	reMem := regexp.MustCompile(`mem\[(\d+)\] = (\d+)`)
	var mask string
	for _, instruction := range instructions {
		if strings.HasPrefix(instruction, "mask") {
			matches := reMask.FindStringSubmatch(instruction)
			mask = matches[1]
		} else {
			matches := reMem.FindStringSubmatch(instruction)
			addr, err := strconv.Atoi(matches[1])
			helper.CheckErr(err)
			val, err := strconv.Atoi(matches[2])
			helper.CheckErr(err)
			memory[addr] = maskedVal(mask, val)
		}
	}
	for _, v := range memory {
		sum += v
	}
	return
}

func maskedVal(mask string, value int) int {
	l := len(mask) - 1
	for i, c := range mask {
		pos := uint(l - i)
		if c == 49 { //1 - set bit
			value |= (1 << pos)
		} else if c == 48 { //0 - clear bit
			value &= ^(1 << pos)
		}
	}
	return value
}

func readMaskingData(fileName string) (input []string) {
	ch := make(chan string)
	go helper.StartReading(fileName, ch)

	for {
		line, ok := <-ch
		if !ok {
			break
		}
		input = append(input, line)
	}
	return
}

// part 2
func memMask(instructions []string) (sum int64) {
	memory := make(map[int64]int64)
	reMask := regexp.MustCompile(`mask = ([0|1|X]+)`)
	reMem := regexp.MustCompile(`mem\[(\d+)\] = (\d+)`)
	var mask string
	for _, instruction := range instructions {
		if strings.HasPrefix(instruction, "mask") {
			matches := reMask.FindStringSubmatch(instruction)
			mask = matches[1]
		} else {
			matches := reMem.FindStringSubmatch(instruction)
			addr, err := strconv.Atoi(matches[1])
			helper.CheckErr(err)
			val, err := strconv.Atoi(matches[2])
			helper.CheckErr(err)

			flAddresses := memories(mask, addr)
			for _, flAddr := range flAddresses {
				memory[flAddr] = int64(val)
			}
		}
	}

	for _, v := range memory {
		sum += v
	}
	return
}

func memories(mask string, addr int) (addresses []int64) {
	// since we are using 36 bit numbers/masks
	bina := fmt.Sprintf("%036b", addr)

	result := make([]byte, len(mask))
	for i, _ := range mask {
		switch mask[i] {
		case []byte("0")[0]:
			result[i] = bina[i]
		default: // 1 and X
			result[i] = mask[i]
		}
	}
	addrs := getFloatingAddrs(string(result))
	for _, addr := range addrs {
		in, err := strconv.ParseInt(addr, 2, 64)
		helper.CheckErr(err)
		addresses = append(addresses, in)
	}
	return
}

func getFloatingAddrs(input string) (addresses []string) {
	if len(input) == 0 {
		return
	}
	addrs := getFloatingAddrs(input[1:])
	if input[0] == []byte("X")[0] {
		for _, addr := range addrs {
			addresses = append(addresses, "1"+addr)
			addresses = append(addresses, "0"+addr)
		}
		if len(addresses) == 0 {
			addresses = append(addresses, "1")
			addresses = append(addresses, "0")
		}
	} else {
		for _, addr := range addrs {
			addresses = append(addresses, string(input[0])+addr)
		}
		if len(addresses) == 0 {
			addresses = append(addresses, string(input[0]))
		}
	}
	return
}
