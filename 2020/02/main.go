package main

import (
	"log"
	"strconv"
	"strings"

	"github.com/chneau/adventofcode/common"
	"github.com/chneau/tt"
)

func main() {
	raw := common.DownloadInput("https://adventofcode.com/2020/day/2/input")
	_ = raw
	goodPasswords := 0
	goodPasswords2 := 0
	defer tt.T()()
	log.Println("PART ONE & TWO")
	for _, policy := range strings.Split(raw, "\n") {
		if policy == "" {
			continue
		}
		dashIndex := strings.Index(policy, "-")
		spaceIndex := strings.Index(policy, " ")
		from, err := strconv.Atoi(policy[:dashIndex])
		if err != nil {
			log.Panicln(err)
		}
		to, err := strconv.Atoi(policy[dashIndex+1 : spaceIndex])
		if err != nil {
			log.Panicln(err)
		}
		letter := policy[spaceIndex+1 : spaceIndex+2]
		password := policy[spaceIndex+4:]
		count := strings.Count(password, letter)
		if count >= from && count <= to {
			goodPasswords++
		}
		// part two
		goodIndex := 0
		if from-1 < len(password) {
			if password[from-1:from] == letter {
				goodIndex++
			}
		}
		if to-1 < len(password) {
			if password[to-1:to] == letter {
				goodIndex++
			}
		}
		if goodIndex == 1 {
			goodPasswords2++
		}
	}
	log.Println("There is", goodPasswords, "good passwords.")
	log.Println("There is", goodPasswords2, "good passwords with the new job rule.")
}
