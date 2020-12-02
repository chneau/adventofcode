package main

import (
	"crypto/md5"
	"log"
	"strconv"
)

func main() {
	input := "yzbqklnj"
	i := 0
	log.Println("PART ONE")
	for ; ; i++ {
		res := md5.Sum([]byte(input + strconv.Itoa(i)))
		if res[0] == 0 && res[1] == 0 && res[2]>>4 == 0 {
			log.Println(i)
			break
		}
	}
	log.Println("PART TWO")
	for ; ; i++ {
		res := md5.Sum([]byte(input + strconv.Itoa(i)))
		if res[0] == 0 && res[1] == 0 && res[2] == 0 {
			log.Println(i)
			break
		}
	}
}
