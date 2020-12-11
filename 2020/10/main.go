package main

import (
	"fmt"
	"log"
	"os"
	"sort"

	"github.com/chneau/adventofcode/common"
)

func main() {
	testPart2()
	raw := common.DownloadInputFor(2020, 10)
	ratings := common.StrToInts(raw)
	sort.Ints(ratings)
	jolts := map[int]int{3: 1}
	previous := 0
	for _, rating := range ratings {
		diff := rating - previous
		jolts[diff]++
		previous = rating
	}
	log.Println("What is the number of 1-jolt differences multiplied by the number of 3-jolt differences?", jolts[1]*jolts[3])

	log.Println("What is the total number of distinct ways you can arrange the adapters to connect the charging outlet to your device?")
}

func testPart2() {
	ratings := common.StrToInts(`28
33
18
42
31
14
46
20
48
47
24
23
49
45
19
38
39
11
1
32
25
35
8
17
7
9
4
2
34
10
3`)
	sort.Ints(ratings)
	ratings = append(append(ratings[:0:0], 0), ratings...)
	fmt.Println(ratings)
	expected := 19208
	has := 1
	for i := range ratings {
		j := i - 3
		if j < 0 {
			j = 0
		}
		possibilities := -1
		for ; j < len(ratings); j++ {
			if i == j {
				continue
			}
			if i < j {
				if ratings[j]-ratings[i] > 3 {
					continue
				}
			} else {
				if ratings[i]-ratings[j] > 3 {
					continue
				}
			}
			possibilities++
		}
		fmt.Print(possibilities)
		fmt.Print(" ")
		has *= possibilities
	}
	fmt.Println("expected", expected, "has", has)
	os.Exit(0)
}
