package main

import (
	"log"
	"strconv"
	"strings"

	"github.com/chneau/adventofcode/common"
)

// Contains ...
type Contains struct {
	Count int
	Name  string
}

func main() {
	raw := common.DownloadInputFor(2020, 7)
	allBags := map[string][]Contains{}
	bagsContained := map[string][]string{}
	for _, line := range strings.Split(raw, "\n") {
		words := strings.Fields(line)
		bagName := words[0] + " " + words[1]
		contains := []Contains{}
		for i := 7; i < len(words); i++ {
			word := words[i]
			if strings.HasPrefix(word, "bag") {
				bName := words[i-2] + " " + words[i-1]
				count, _ := strconv.Atoi(words[i-3])
				contains = append(contains, Contains{Count: count, Name: bName})
				bagsContained[bName] = append(bagsContained[bName], bagName)
			}
		}
		allBags[bagName] = contains
	}
	everContains := map[string]bool{}
	for queue := []string{"shiny gold"}; len(queue) != 0; {
		element := queue[0]
		queue = queue[1:]
		queue = append(queue, bagsContained[element]...)
		everContains[element] = true
	}
	log.Println("How many bag colors can eventually contain at least one shiny gold bag?", len(everContains)-1)
	log.Println("How many individual bags are required inside your single shiny gold bag?", CountBags(allBags, "shiny gold")-1)
}

// CountBags ...
func CountBags(allBags map[string][]Contains, actualBag string) int {
	result := 1
	for _, bag := range allBags[actualBag] {
		result += CountBags(allBags, bag.Name) * bag.Count
	}
	return result
}
