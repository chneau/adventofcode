package main

import (
	"log"
	"strings"

	"github.com/chneau/adventofcode/common"
)

func main() {
	raw := common.DownloadInput("https://adventofcode.com/2015/day/5/input")
	vowels := []string{"a", "e", "i", "o", "u"}
	unwanteds := []string{"ab", "cd", "pq", "xy"}
	goodLines := 0
	log.Println("PART ONE")
	for _, line := range strings.Split(raw, "\n") {
		if line == "" {
			continue
		}
		vowelsCount := 0
		for _, vowel := range vowels {
			vowelsCount += strings.Count(line, vowel)
			if vowelsCount == 3 {
				break
			}
		}
		if vowelsCount < 3 {
			continue
		}
		hasDouble := false
		for i := 0; i < len(line)-1; i++ {
			if line[i] == line[i+1] {
				hasDouble = true
				break
			}
		}
		if hasDouble == false {
			continue
		}
		hasForbiden := false
		for _, unwanted := range unwanteds {
			if strings.Contains(line, unwanted) {
				hasForbiden = true
				break
			}
		}
		if hasForbiden {
			continue
		}
		log.Println(vowelsCount, hasDouble, hasForbiden)
		goodLines++
		log.Println(line)
	}
	log.Println("Good lines:", goodLines)
}
