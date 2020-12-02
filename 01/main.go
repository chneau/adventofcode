package main

import (
	"log"
	"sort"

	"github.com/chneau/adventofcode2020/common"
	"github.com/chneau/tt"
)

func main() {
	input := common.DownloadInput("https://adventofcode.com/2020/day/1/input")
	defer tt.T()()
	sort.Ints(input)
	log.Println("PART ONE")
	for i := 0; i < len(input); i++ {
		for j := i + 1; j < len(input); j++ {
			add := input[i] + input[j]
			if add == 2020 {
				log.Println(input[i], "+", input[j], "= 2020")
				log.Println(input[i], "*", input[j], "=", input[i]*input[j])
			}
			if add > 2020 {
				break
			}
		}
	}
	log.Println("PART TWO")
	for i := 0; i < len(input); i++ {
		for j := i + 1; j < len(input); j++ {
			ijadd := input[i] + input[j]
			if ijadd > 2020 {
				break
			}
			for k := j + 1; k < len(input); k++ {
				add := ijadd + input[k]
				if add == 2020 {
					log.Println(input[i], "+", input[j], "+", input[k], "= 2020")
					log.Println(input[i], "*", input[j], "*", input[k], "=", input[i]*input[j]*input[k])
				}
				if add > 2020 {
					break
				}

			}
		}
	}
}
