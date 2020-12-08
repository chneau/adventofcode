package main

import (
	"log"
	"strconv"
	"strings"

	"github.com/chneau/adventofcode/common"
)

func main() {
	raw := common.DownloadInputFor(2020, 8)
	instructions := strings.Split(raw, "\n")
	accumulator, _ := do(instructions, -1)
	log.Println("what value is in the accumulator?", accumulator)
	for i := 0; i < len(instructions); i++ {
		accumulator, cleanExit := do(instructions, i)
		if cleanExit {
			log.Println("What is the value of the accumulator after the program terminates?", accumulator)
			break
		}
	}
}

func do(instructions []string, swap int) (accumulator int, cleanExit bool) {
	visited := make([]bool, len(instructions))
	i := 0
	for i < len(visited) && !visited[i] {
		visited[i] = true
		instruction := instructions[i][:3]
		if i == swap {
			switch instruction {
			case "nop":
				instruction = "jmp"
			case "jmp":
				instruction = "nop"
			}
		}
		switch instruction {
		case "nop":
			i++
		case "acc":
			value, _ := strconv.Atoi(instructions[i][4:])
			accumulator += value
			i++
		case "jmp":
			value, _ := strconv.Atoi(instructions[i][4:])
			i += value
		}
	}
	cleanExit = i == len(instructions)
	return
}
