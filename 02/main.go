package main

import (
	"log"
	"strconv"
	"strings"

	"github.com/chneau/adventofcode2020/common"
	"github.com/chneau/tt"
)

func main() {
	raw := common.DownloadInput("https://adventofcode.com/2020/day/2/input")
	_ = raw
	goodPasswords := 0
	defer tt.T()()
	log.Println("PART ONE")
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
	}
	log.Println("There is", goodPasswords, "good passords.")
}
