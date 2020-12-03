package main

import (
	"log"
	"strings"

	"github.com/chneau/adventofcode/common"
	"github.com/chneau/tt"
)

func main() {
	raw := common.DownloadInput("https://adventofcode.com/2020/day/3/input")
	defer tt.T()()
	log.Println("PART ONE")
	log.Println("Trees encountered", move(raw, 3, 1))
	log.Println("PART TWO")
	log.Println("Multiplication result", move(raw, 1, 1)*move(raw, 3, 1)*move(raw, 5, 1)*move(raw, 7, 1)*move(raw, 1, 2))
}

func move(raw string, right, down int) (trees int) {
	visitedLines := 0
	for i, line := range strings.Split(raw, "\n") {
		if i != 0 && i%down != 0 {
			continue
		}
		position := visitedLines * right
		if line[position%len(line)] == '#' {
			trees++
		}
		visitedLines++
	}
	return
}
