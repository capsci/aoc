package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/capsci/aoc/helper"
)

const monsterDataSampleFile = "19.data.sample.txt"
const monsterDataFile = "19.data.txt"
const monsterData2File = "19.data.2.txt"

var reRuleChar = regexp.MustCompile(`"(\w)"`)
var reRuleOr = regexp.MustCompile(`(.*) | (.*)`)
var reRuleAnd = regexp.MustCompile(`(\d+) `)

var rules map[string]Rule

func main() {
	inputs := readMonsterInput(monsterData2File)
	expand("0")
	fmt.Println("8", rules["8"].Expanded, rules["8"].Rule, rules["8"].Expansion)
	fmt.Println("11", rules["11"].Expanded, rules["11"].Rule, rules["11"].Expansion)

	fmt.Println(len(rules["0"].Expansion))
	matches := 0
	return
	for _, input := range inputs {
		for _, r := range rules["0"].Expansion {
			if input == r {
				matches++
			}
		}
	}
	fmt.Printf("Number of strings that matches the rules:%d\n", matches)
}

type Rule struct {
	Rule      string
	Expanded  bool
	Expansion []string
}

func expand(key string) []string {
	r, ok := rules[key]
	if !ok {
		panic("Rule not found: " + key)
	}
	if r.Expanded {
		return r.Expansion
	}
	ored := []string{}
	for _, ors := range strings.Split(rules[key].Rule, " | ") {
		anded := []string{}
		for _, ands := range strings.Split(ors, " ") {
			if ands == key {
				anded = and(anded, []string{"*"})
				continue
			}
			anded = and(anded, expand(ands))
		}
		ored = append(ored, anded...)
	}
	rules[key] = Rule{
		Rule:      r.Rule,
		Expanded:  true,
		Expansion: ored,
	}
	return ored
}

func and(part1, part2 []string) (result []string) {
	if len(part1) == 0 {
		return part2
	}
	if len(part2) == 0 {
		return part1
	}
	for _, p1 := range part1 {
		for _, p2 := range part2 {
			result = append(result, p1+p2)
		}
	}
	return
}

func readMonsterInput(fileName string) (inputs []string) {
	content := strings.Split(helper.ReadFile(fileName), "\n\n")

	rules = make(map[string]Rule)
	for _, rule := range strings.Split(content[0], "\n") {
		reRule := regexp.MustCompile(`(\d+): (.*)`)
		matches := reRule.FindStringSubmatch(rule)
		if reRuleChar.MatchString(matches[2]) {
			// matches[1] will be the character
			char := reRuleChar.FindStringSubmatch(matches[2])[1]
			rules[matches[1]] = Rule{Rule: matches[2],
				Expanded:  true,
				Expansion: []string{char}}
		} else {
			rules[matches[1]] = Rule{Rule: matches[2],
				Expansion: []string{}}
		}

	}
	inputs = strings.Split(content[1], "\n")
	return
}
