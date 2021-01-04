package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/capsci/aoc/helper"
)

const luggageDataFile = "7.data.txt"
const luggageSampleDataFile = "7.data.sample.txt"
const luggageSample2DataFile = "7.data.sample.2.txt"

const outerBagCount = "Number of bags that can contain %s bag : %d\n"
const innerBagCount = "Number of bags that can be contained in %s bag : %d\n"

func main() {
	rules := readLuggageData(luggageDataFile)
	outerBags := make(map[string]int)
	findOuterBags("shiny gold", rules, outerBags)
	fmt.Printf(outerBagCount, "shiny gold", len(outerBags))

	newrules := readLuggageData2(luggageDataFile)
	// Subtract 1 since we dont want "shiny gold" bag itself
	fmt.Printf(innerBagCount, "shiny gold", findInnerBags("shiny gold", newrules)-1)
}

func findOuterBags(color string, rulebook map[string][]string, bags map[string]int) {
	for _, outer := range rulebook[color] {
		bags[outer] = 1
		findOuterBags(outer, rulebook, bags)
	}
	return
}

func findInnerBags(color string, rulebook map[string]map[string]int) (sum int) {
	for k, v := range rulebook[color] {
		sum += v * findInnerBags(k, rulebook)
	}
	// Add self to can contain bags
	sum++
	return
}

func readLuggageData(fileName string) (rulebook map[string][]string) {
	re := regexp.MustCompile(`([a-z ]+) bag`)
	rulebook = make(map[string][]string)
	ch := make(chan string)
	go helper.StartReading(fileName, ch)
	for {
		line, ok := <-ch
		if !ok {
			break
		}
		matches := re.FindAllStringSubmatch(line, -1)
		outer := strings.Trim(matches[0][1], " ")
		for _, match := range matches[1:] {
			inner := strings.Trim(match[1], " ")
			addToOuterRules(outer, inner, rulebook)
		}
	}
	return
}

func readLuggageData2(fileName string) (rulebook map[string]map[string]int) {
	re := regexp.MustCompile(`(\d*)?\s*([a-z ]+) bag`)
	rulebook = make(map[string]map[string]int)
	ch := make(chan string)
	go helper.StartReading(fileName, ch)
	for {
		line, ok := <-ch
		if !ok {
			break
		}
		matches := re.FindAllStringSubmatch(line, -1)
		outer := strings.Trim(matches[0][2], " ")
		rulebook[outer] = make(map[string]int)
		for _, inner := range matches[1:] {
			count, err := strconv.Atoi(inner[1])
			helper.CheckErr(err)
			rulebook[outer][inner[2]] = count
		}
	}
	return
}

// Returns which bags can be outer bags
func addToOuterRules(outer, inner string, rulebook map[string][]string) {
	_, exists := rulebook[inner]
	if !exists || !inList(rulebook[inner], outer) {
		rulebook[inner] = append(rulebook[inner], outer)
	}
}

func inList(list []string, find string) bool {
	for _, item := range list {
		if item == find {
			return true
		}
	}
	return false
}
