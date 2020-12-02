package main

import (
	"log"

	"github.com/chneau/adventofcode/common"
)

// Pos ...
type Pos struct{ x, y int }

func main() {
	raw := common.DownloadInput("https://adventofcode.com/2015/day/3/input")
	log.Println("PART ONE")
	{
		visits := map[Pos]int{{0, 0}: 1}
		current := Pos{0, 0}
		for i := range raw {
			char := raw[i : i+1]
			switch char {
			case ">":
				current.x++
			case "<":
				current.x--
			case "^":
				current.y++
			case "v":
				current.y--
			}
			visits[current]++
		}
		log.Println("Visits", len(visits))
	}
	log.Println("PART TWO")
	{
		visits := map[Pos]int{{0, 0}: 2}
		santa := Pos{0, 0}
		robo := Pos{0, 0}
		for i := range raw {
			char := raw[i : i+1]
			user := &santa
			if i%2 == 1 {
				user = &robo
			}
			switch char {
			case ">":
				user.x++
			case "<":
				user.x--
			case "^":
				user.y++
			case "v":
				user.y--
			}
			visits[*user]++
		}
		log.Println("Visits with robo", len(visits))
	}
}
