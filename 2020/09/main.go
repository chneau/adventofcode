package main

import (
	"log"
	"sort"
	"strconv"
	"strings"

	"github.com/chneau/adventofcode/common"
)

func main() {
	raw := common.DownloadInputFor(2020, 9)
	numbers := toSliceOfInts(raw)
	target := 0
	for i := 25; i < len(numbers); i++ {
		if !isCorrect(numbers, 25, i) {
			target = numbers[i]
			log.Println("What is the first number that does not have this property?", target)
			break
		}
	}
	answer := findWeakness(numbers, target)
	log.Println("What is the encryption weakness in your XMAS-encrypted list of numbers?", answer)
}

func findWeakness(numbers []int, target int) int {
	weakness := []int{}
	for i := 0; i < len(numbers); i++ {
		sum := 0
		j := i
		for ; j < len(numbers); j++ {
			sum += numbers[j]
			if sum >= target {
				break
			}
		}
		if sum == target {
			weakness = numbers[i:j]
			break
		}
	}
	sort.Ints(weakness)
	min := weakness[0]
	max := weakness[len(weakness)-1]
	answer := min + max
	return answer
}

func isCorrect(numbers []int, lastN, index int) bool {
	lasts := []int{}
	for i := index - lastN; i < index; i++ {
		lasts = append(lasts, numbers[i])
	}
	sort.Ints(lasts)
	target := numbers[index]
	for i, value := range lasts { // remove numbers too high
		if value >= target {
			lasts = lasts[:i]
			break
		}
	}
	return hasTwoSum(lasts, target)
}

func hasTwoSum(nums []int, target int) bool { // sourced and modified from https://stackoverflow.com/a/34672624
	for _, v := range nums {
		v2 := target - v
		i := sort.SearchInts(nums, v2)
		if i == len(nums) {
			continue
		}
		if v2 == nums[i] {
			return true
		}
	}
	return false
}

func toSliceOfInts(raw string) []int {
	lines := strings.Split(raw, "\n")
	numbers := make([]int, len(lines))
	for i, numberStr := range lines {
		numbers[i], _ = strconv.Atoi(numberStr)
	}
	return numbers
}
