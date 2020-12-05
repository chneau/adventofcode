package main

import (
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/chneau/adventofcode/common"
)

var requiredFields = []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
var hclRegex = regexp.MustCompile("#[0-9a-f]{6}")
var eclRegex = regexp.MustCompile("amb|blu|brn|gry|grn|hzl|oth") // lazy solution
var pidRegex = regexp.MustCompile("[0-9]{9}")

func main() {
	raw := common.DownloadInput("https://adventofcode.com/2020/day/4/input")
	goodPassports := 0
	for _, passport := range strings.Split(raw, "\n\n") {
		dic := map[string]string{}
		fields := strings.Fields(passport)
		for _, field := range fields {
			index := strings.Index(field, ":")
			dic[field[:index]] = field[index+1:]
		}
		if !correctFields(dic) { // for part 1
			continue
		}
		if !correctNumber(dic["byr"], 1920, 2002) {
			continue
		}
		if !correctNumber(dic["iyr"], 2010, 2020) {
			continue
		}
		if !correctNumber(dic["eyr"], 2020, 2030) {
			continue
		}
		if !correctHgt(dic["hgt"]) {
			continue
		}
		if len(dic["hcl"]) != 7 || !hclRegex.MatchString(dic["hcl"]) {
			continue
		}
		if len(dic["ecl"]) != 3 || !eclRegex.MatchString(dic["ecl"]) {
			continue
		}
		if len(dic["pid"]) != 9 || !pidRegex.MatchString(dic["pid"]) {
			continue
		}
		goodPassports++
	}
	log.Println("Good passports", goodPassports)
}

func correctHgt(hgt string) bool {
	good := true
	switch {
	case strings.HasSuffix(hgt, "cm"):
		good = correctNumber(strings.TrimSuffix(hgt, "cm"), 150, 193)
	case strings.HasSuffix(hgt, "in"):
		good = correctNumber(strings.TrimSuffix(hgt, "in"), 59, 76)
	default:
		good = false
	}
	return good
}

func correctNumber(byr string, from, to int) bool {
	good := true
	value, _ := strconv.Atoi(byr)
	if value < from || value > to {
		good = false
	}
	return good
}

func correctFields(dic map[string]string) bool {
	good := true
	for _, requirequiredField := range requiredFields {
		if _, exist := dic[requirequiredField]; !exist {
			good = false
			break
		}
	}
	return good
}
