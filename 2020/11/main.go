package main

import (
	"log"
	"strings"

	"github.com/chneau/adventofcode/common"
)

func main() {
	raw := common.DownloadInputFor(2020, 11)
	lines := common.StrsSplitBy(raw, "\n")
	i := 0
	for {
		new := mutate(lines)
		if linesEqual(lines, new) {
			break
		}
		lines = new
		i++
	}
	hashCount := 0
	for _, line := range lines {
		hashCount += strings.Count(line, "#")
	}
	log.Println("How many seats end up occupied?", hashCount)
}

func linesEqual(a, b []string) bool {
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func mutate(lines []string) []string {
	copy := cloneLines(lines)
	for y := 0; y < len(lines); y++ {
		for x := 0; x < len(lines[y]); x++ {
			counts := countAround(lines, y, x)
			switch lines[y][x] {
			case 'L':
				if counts['#'] == 0 {
					replaceAtIndex(&copy[y], x, "#")
				}
			case '#':
				if counts['#'] >= 4 {
					replaceAtIndex(&copy[y], x, "L")
				}
			}
		}
	}
	return copy
}

func countAround(lines []string, y, x int) (result map[byte]int) {
	result = map[byte]int{lines[y][x]: -1}
	yMax := min(y+2, len(lines))
	xMax := min(x+2, len(lines[0]))
	for y := max(y-1, 0); y < yMax; y++ {
		for x := max(x-1, 0); x < xMax; x++ {
			result[lines[y][x]]++
		}
	}
	return
}

func printLines(lines []string) {
	for _, line := range lines {
		log.Println(line)
	}
}

func cloneLines(lines []string) []string {
	result := []string{}
	for _, line := range lines {
		result = append(result, line[:])
	}
	return result
}

func replaceAtIndex(str *string, index int, replacement string) {
	*str = (*str)[:index] + replacement + (*str)[index+1:]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
