package main

import (
	"fmt"

	"github.com/capsci/aoc/helper"
)

const customsDataFile = "6.data.txt"
const customsSampleDataFile = "6.data.sample.txt"

const yesAnyoneCount = "Number Questions answered Yes by anyone in the group is %d\n"
const yesEveryoneCount = "Number Questions answered Yes by everyOneone in the group is %d\n"

func main() {
	grps := readCustomsData(customsDataFile)

	var anyone, everyone int
	for _, grp := range grps {
		anyone += grp.AnyoneCount
		everyone += grp.EveryoneCount()
	}
	fmt.Printf(yesAnyoneCount, anyone)
	fmt.Printf(yesEveryoneCount, everyone)
}

type Group struct {
	Answers     [26]int
	AnyoneCount int
	MemberCount int
}

func (g *Group) SetAnswer(answer string) {
	for _, ch := range answer {
		if g.Answers[ch-97] == 0 {
			g.AnyoneCount++
		}
		g.Answers[ch-97] += 1
	}
}

func (g *Group) EveryoneCount() (count int) {
	for _, yays := range g.Answers {
		if yays == g.MemberCount {
			count++
		}
	}
	return
}

func readCustomsData(fileName string) (groups []Group) {
	var group Group

	ch := make(chan string)
	go helper.StartReading(fileName, ch)
	for {
		line, ok := <-ch
		if !ok {
			break
		}
		if len(line) == 0 {
			groups = append(groups, group)
			group = Group{}
			continue
		}
		group.SetAnswer(line)
		group.MemberCount++
	}
	groups = append(groups, group)
	return
}
