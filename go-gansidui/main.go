package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	ahocorasick "github.com/gansidui/ahocorasick"
)

func main() {
	timeBuildStart := time.Now()

	dict, err := readWords("words.txt")
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}

	ac := ahocorasick.BuildNewMatcher(dict)

	timeBuildEnd := time.Now()

	message := strings.ToLower("I aguaranteeee the return to capital of this business cycle yields a lot of money related to bank loan swap with very basic risk on this base year.")

	timeMatchStart := time.Now()

	indexes := ac.Match(message)

	timeMatchEnd := time.Now()

	fmt.Printf("dictionary: count = %d\n", len(dict))
	fmt.Printf("message: text = %s\n\n", message)

	for _, term := range indexes {
		fmt.Printf("%d %s\n", term.EndPosition-len(dict[term.Index])+1, dict[term.Index])
	}

	fmt.Printf("\ntime: build = %.5fs, execute = %.5fs\n", timeBuildEnd.Sub(timeBuildStart).Seconds(), timeMatchEnd.Sub(timeMatchStart).Seconds())
}

func readWords(filename string) ([]string, error) {
	var dict []string

	f, err := os.OpenFile(filename, os.O_RDONLY, 0660)
	if err != nil {
		return nil, err
	}

	s := bufio.NewScanner(f)
	for s.Scan() {
		l := strings.ToLower(strings.TrimSpace(s.Text()))
		dict = append(dict, l)
	}

	return dict, s.Err()
}
