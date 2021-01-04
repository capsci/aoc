package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/capsci/aoc/helper"
)

const passwordDataFile = "2.data.txt"
const validPasswords = "Count of valid passwords : %d\n"

func main() {

	entries := readPasswordData(passwordDataFile)
	password1(entries)
	password2(entries)

}

func password1(entries []passwordData) {
	valid := 0
	for _, entry := range entries {
		cnt := strings.Count(entry.Password, entry.Letter)
		if cnt >= entry.Min && cnt <= entry.Max {
			valid += 1
		}
	}
	fmt.Printf(validPasswords, valid)
}

func password2(entries []passwordData) {
	valid := 0
	for _, entry := range entries {
		bytes := []byte(entry.Password)
		min := string(bytes[entry.Min-1]) == entry.Letter
		max := string(bytes[entry.Max-1]) == entry.Letter
		if (min || max) && !(min && max) {
			valid += 1
		}
	}
	fmt.Printf(validPasswords, valid)
}

type passwordData struct {
	Min, Max int
	Letter   string
	Password string
}

func readPasswordData(fileName string) (data []passwordData) {
	ch := make(chan string)
	go helper.StartReading(fileName, ch)
	for {
		val, ok := <-ch
		if !ok {
			break
		}
		r := regexp.MustCompile(`(?P<Min>\d+)-(?P<Max>\d+) (?P<Letter>\w): (?P<Passowrd>\w+)`)
		matches := r.FindStringSubmatch(val)
		/*
			matches[0] is complete match
			matches[1] is Min
			matches[2] is Max
			matches[3] is Letter
			matches[4] is Password
		*/
		min, err := strconv.Atoi(matches[1])
		helper.CheckErr(err)
		max, err := strconv.Atoi(matches[2])
		helper.CheckErr(err)
		data = append(data, passwordData{Min: min,
			Max: max, Letter: matches[3], Password: matches[4]})
	}
	return data
}
