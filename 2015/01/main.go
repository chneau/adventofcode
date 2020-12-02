package main

import (
	"log"
	"strings"

	"github.com/chneau/adventofcode2020/common"
)

func main() {
	raw := common.DownloadInput("https://adventofcode.com/2015/day/1/input")
	log.Println("PART ONE")
	floor := strings.Count(raw, "(") - strings.Count(raw, ")")
	log.Println("The instructions takes santa to the floor", floor, ".")
	log.Println("PART TWO")
	currentFloor := 0
	for i := 0; i < len(raw); i++ {
		if raw[i] == '(' {
			currentFloor++
		} else {
			currentFloor--
		}
		if currentFloor == -1 {
			log.Println("Character is one basement at index", i+1, ".")
			break
		}
	}
}
