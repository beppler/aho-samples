package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"
	"time"

	ahocorasick "github.com/BobuSumisu/aho-corasick"
)

func main() {
	timeBuildStart := time.Now()

	dict, err := readWords("words.txt")
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}

	matcher := ahocorasick.NewTrieBuilder().
		AddStrings(dict).
		Build()

	timeBuildEnd := time.Now()

	buffer := new(bytes.Buffer)
	if err := ahocorasick.Encode(buffer, matcher); err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}
	fmt.Printf("%d\n", buffer.Len())

	message := strings.ToLower("The return to capital of this business cycle yields a lot of money related to bank loan swap with very basic risk on this base year.")

	timeMatchStart := time.Now()

	terms := matcher.MatchString(message)

	timeMatchEnd := time.Now()

	fmt.Printf("dictionary: count = %d\n", len(dict))
	fmt.Printf("message: text = %s\n\n", string(message))

	for _, t := range terms {
		fmt.Printf("%d %s\n", t.Pos(), string(t.MatchString()))
	}

	fmt.Printf("\ntime: build = %v, execute = %v\n", timeBuildEnd.Sub(timeBuildStart), timeMatchEnd.Sub(timeMatchStart))
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
