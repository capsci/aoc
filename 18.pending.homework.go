package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/capsci/aoc/helper"
)

const homeworkSampleDataFile = "18.data.sample.txt"
const homeworkDataFile = "18.data.txt"

const sumResultingVal = "Sum of resulting values is %d\n"

func main() {
	fmt.Printf(sumResultingVal, solveHomeWork(homeworkDataFile, rule1))
	fmt.Printf(sumResultingVal, solveHomeWork(homeworkSampleDataFile, rule2))
}

func solveHomeWork(fileName string, rule func(string) string) (ans int) {
	ch := make(chan string)
	go helper.StartReading(fileName, ch)

	for {
		line, ok := <-ch
		if !ok {
			break
		}
		ans += solveQuestion(line, rule)
	}
	return
}

func rule2(question string) string {
	// Perform all additions first
	rePlus := regexp.MustCompile(`(\d+ + \d+)`)
	matches := rePlus.FindStringSubmatch(question)
	for len(matches) > 1 {
		question = strings.Replace(question, matches[0], rule1(matches[1]), 1)
		matches = rePlus.FindStringSubmatch(question)
	}
	return rule1(question)
}

func rule1(question string) string {
	parts := strings.Split(question, " ")
	var operation string
	ans, err := strconv.Atoi(parts[0])
	helper.CheckErr(err)
	for _, part := range parts[1:] {
		if len(operation) == 0 {
			operation = part
		} else {
			num2, err := strconv.Atoi(part)
			helper.CheckErr(err)
			switch operation {
			case "+":
				ans += num2
			case "*":
				ans *= num2
			}
			operation = ""
		}
	}
	return strconv.Itoa(ans)
}

func solveQuestion(question string, rule func(string) string) (ans int) {
	reGroup := regexp.MustCompile(`\(([\s\d\+\*]+)\)`)
	matches := reGroup.FindStringSubmatch(question)
	//	fmt.Println(question, ":", strings.Join(matches, ","))
	for len(matches) > 1 {
		question = strings.Replace(question, matches[0], rule(matches[1]), 1)
		matches = reGroup.FindStringSubmatch(question)
	}
	ans, err := strconv.Atoi(rule(question))
	helper.CheckErr(err)
	fmt.Println(ans)
	return
}
