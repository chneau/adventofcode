package common

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

// DownloadInput ...
func DownloadInput(inputURL string) []int {
	session := os.Getenv("AOC")
	if session == "" {
		log.Panicln("Env var AOC is not defined with the session value")
	}
	req, err := http.NewRequest("GET", inputURL, nil)
	if err != nil {
		log.Panicln(err)
	}
	req.Header.Set("Cookie", "session="+session)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Panicln(err)
	}
	defer resp.Body.Close()
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Panicln(err)
	}
	strs := strings.Split(string(b), "\n")
	result := make([]int, len(strs))
	for i, str := range strs {
		if str == "" {
			continue
		}
		result[i], err = strconv.Atoi(str)
		if err != nil {
			log.Panicln(err)
		}
	}
	return result
}
