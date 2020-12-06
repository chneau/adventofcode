package main

import (
	"log"
	"strings"

	"github.com/chneau/adventofcode/common"
)

func main() {
	raw := common.DownloadInputFor(2020, 6)
	totalAnyoneVotes := 0
	totalEveryoneVotes := 0
	for _, group := range strings.Split(raw, "\n\n") {
		votes := map[rune]int{}
		nbPerson := strings.Count(group, "\n") + 1
		for _, c := range strings.ReplaceAll(group, "\n", "") {
			votes[c]++
		}
		totalAnyoneVotes += len(votes)
		for _, v := range votes {
			if v == nbPerson {
				totalEveryoneVotes++
			}
		}
	}
	log.Println("P1 - What is the sum of those counts?", totalAnyoneVotes)
	log.Println("P2 - What is the sum of those counts?", totalEveryoneVotes)
}
