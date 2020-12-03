package main

import (
	"log"
	"sort"
	"strconv"
	"strings"

	"github.com/chneau/adventofcode/common"
)

func main() {
	raw := common.DownloadInput("https://adventofcode.com/2015/day/2/input")
	totalWrap := 0
	totalRibbon := 0
	for _, dimmension := range strings.Split(raw, "\n") {
		firstX := strings.Index(dimmension, "x")
		lastX := strings.LastIndex(dimmension, "x")
		l, err := strconv.Atoi(dimmension[:firstX])
		if err != nil {
			log.Panicln(err)
		}
		w, err := strconv.Atoi(dimmension[firstX+1 : lastX])
		if err != nil {
			log.Panicln(err)
		}
		h, err := strconv.Atoi(dimmension[lastX+1:])
		if err != nil {
			log.Panicln(err)
		}

		presentDimmension := 2*l*w + 2*w*h + 2*h*l
		sortedDims := []int{l, w, h}
		sort.Ints(sortedDims)
		extra := sortedDims[0] * sortedDims[1]
		totalWrap += presentDimmension + extra

		totalRibbon += sortedDims[0]*2 + sortedDims[1]*2 + l*w*h
	}
	log.Println("Total square feet needed", totalWrap)
	log.Println("Total ribbon needed", totalRibbon)
}
