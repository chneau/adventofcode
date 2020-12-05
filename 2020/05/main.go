package main

import (
	"log"
	"strconv"
	"strings"

	"github.com/chneau/adventofcode/common"
)

func main() {
	raw := common.DownloadInput("https://adventofcode.com/2020/day/5/input")
	biggestSeatID := 0
	max := 127*8 + 7
	seats := make([]bool, max)
	for _, boardingpass := range strings.Split(raw, "\n") {
		row := toDecimal(boardingpass[:7], "B")
		column := toDecimal(boardingpass[7:], "R")
		seatID := row*8 + column
		if seatID > biggestSeatID {
			biggestSeatID = seatID
		}
		seats[seatID] = true
	}
	log.Println("What is the highest seat ID on a boarding pass?", biggestSeatID)
	for i := 1; i < len(seats); i++ {
		current := seats[i]
		previous := seats[i-1]
		if !current && previous {
			log.Println("What is the ID of your seat?", i)
			break
		}
	}
}

func toDecimal(str, one string) int {
	bin := ""
	for i := range str {
		if str[i] == one[0] {
			bin += "1"
			continue
		}
		bin += "0"
	}
	result, _ := strconv.ParseInt(bin, 2, 64)
	return int(result)
}
