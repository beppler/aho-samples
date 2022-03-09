package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	ahocorasick "github.com/anknown/ahocorasick"
)

func main() {
	timeBuildStart := time.Now()

	dict, err := readRunes("words.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	m := new(ahocorasick.Machine)
	if err := m.Build(dict); err != nil {
		fmt.Println(err)
		return
	}

	timeBuildEnd := time.Now()

	message := []rune(strings.ToLower("I aguaranteeee the return to capital of this business cycle yields a lot of money related to bank loan swap with very basic risk on this base year."))

	timeMatchStart := time.Now()

	terms := m.MultiPatternSearch(message, false)

	timeMatchEnd := time.Now()

	fmt.Printf("dictionary: count = %d\n", len(dict))
	fmt.Printf("message: text = %s\n\n", string(message))

	for _, t := range terms {
		fmt.Printf("%d %s\n", t.Pos, string(t.Word))
	}

	fmt.Printf("\ntime: build = %.5fs, execute = %.5fs\n", timeBuildEnd.Sub(timeBuildStart).Seconds(), timeMatchEnd.Sub(timeMatchStart).Seconds())
}

func readRunes(filename string) ([][]rune, error) {
	dict := [][]rune{}

	f, err := os.OpenFile(filename, os.O_RDONLY, 0660)
	if err != nil {
		return nil, err
	}

	s := bufio.NewScanner(f)
	for s.Scan() {
		l := strings.ToLower(strings.TrimSpace(s.Text()))
		dict = append(dict, []rune(l))
	}

	return dict, s.Err()
}
