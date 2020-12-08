package main

import (
	"log"
	"strconv"
	"strings"

	"github.com/chneau/adventofcode/common"
)

func main() {
	raw := common.DownloadInputFor(2015, 6)
	square := createSquare(1000)
	for _, instruction := range strings.Split(raw, "\n") {
		var action func(int, int)
		switch instruction[:7] {
		case "turn on":
			action = func(i, j int) { square[i][j] = 1 }
		case "turn of":
			action = func(i, j int) { square[i][j] = 0 }
		case "toggle ":
			action = func(i, j int) {
				if square[i][j] == 1 {
					square[i][j] = 0
				} else {
					square[i][j] = 1
				}
			}
		}
		words := strings.Fields(instruction)
		x, y := extractPos(words[len(words)-3])
		X, Y := extractPos(words[len(words)-1])
		for i := x; i <= X; i++ {
			for j := y; j <= Y; j++ {
				action(i, j)
			}
		}
	}
	log.Println("how many lights are lit?", countLightsOn(square))
	square = createSquare(1000)
	for _, instruction := range strings.Split(raw, "\n") {
		var action func(int, int)
		switch instruction[:7] {
		case "turn on":
			action = func(i, j int) { square[i][j]++ }
		case "turn of":
			action = func(i, j int) {
				square[i][j]--
				if square[i][j] < 0 {
					square[i][j] = 0
				}
			}
		case "toggle ":
			action = func(i, j int) { square[i][j] += 2 }
		}
		words := strings.Fields(instruction)
		x, y := extractPos(words[len(words)-3])
		X, Y := extractPos(words[len(words)-1])
		for i := x; i <= X; i++ {
			for j := y; j <= Y; j++ {
				action(i, j)
			}
		}
	}
	log.Println("What is the total brightness of all lights combined after following Santa's instructions?", countLightsOn(square))
}

func countLightsOn(lights [][]int) (count int) {
	for i := 0; i < len(lights); i++ {
		for j := 0; j < len(lights[i]); j++ {
			count += lights[i][j]
		}
	}
	return
}

func extractPos(str string) (x, y int) {
	xy := strings.Split(str, ",")
	x, _ = strconv.Atoi(xy[0])
	y, _ = strconv.Atoi(xy[1])
	return
}

func createSquare(size int) [][]int {
	square := make([][]int, size)
	for i := range square {
		square[i] = make([]int, size)
	}
	return square
}
