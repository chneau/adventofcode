package main

import (
	"log"
	"sort"
	"strconv"
	"strings"

	"github.com/chneau/adventofcode/common"
)

func main() {
	raw := common.DownloadInputFor(2020, 10)
	ratings := []int{}
	for _, ratingStr := range strings.Split(raw, "\n") {
		rating, _ := strconv.Atoi(ratingStr)
		ratings = append(ratings, rating)
	}
	previous := 0
	sort.Ints(ratings)
	jolts := map[int]int{3: 1}
	for _, rating := range ratings {
		diff := rating - previous
		jolts[diff]++
		previous = rating
	}
	log.Println("What is the number of 1-jolt differences multiplied by the number of 3-jolt differences?", jolts[1]*jolts[3])
}
