package main

import (
	"log"
	"strings"

	"github.com/chneau/adventofcode/common"
)

func main() {
	raw := common.DownloadInput("https://adventofcode.com/2015/day/5/input")
	log.Println("PART ONE")
	{
		vowels := []string{"a", "e", "i", "o", "u"}
		unwanteds := []string{"ab", "cd", "pq", "xy"}
		goodLines := 0
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
			goodLines++
		}
		log.Println("Good lines:", goodLines)
	}
	log.Println("PART TWO")
	{
		goodLines := 0
		for _, line := range strings.Split(raw, "\n") {
			hasDouble := false
			doubles := ""
			for i := 0; i < len(line)-2; i++ {
				if strings.Count(line, line[i:i+2]) >= 2 {
					hasDouble = true
				}
			}
			if hasDouble == false {
				continue
			}
			if strings.Count(line, doubles) < 2 {
				continue
			}
			hasWeirdDouble := false
			for i := 0; i < len(line)-2; i++ {
				if line[i] == line[i+2] {
					hasWeirdDouble = true
					break
				}
			}
			if hasWeirdDouble == false {
				continue
			}

			goodLines++
		}
		log.Println("Good lines:", goodLines)
	}
}
