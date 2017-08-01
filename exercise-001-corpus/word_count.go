package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"
)

type KeyVal struct {
	Word  string
	Count int
}

type KeyVals []KeyVal

func (slice KeyVals) Len() int {
	return len(slice)
}

func (slice KeyVals) Swap(i, j int) {
	tmpK, tmpV := slice[i].Word, slice[i].Count
	slice[i].Word, slice[i].Count = slice[j].Word, slice[j].Count
	slice[j].Word, slice[j].Count = tmpK, tmpV
}

func (slice KeyVals) Less(i, j int) bool {
	// backwards to sort by most frequently seen
	return slice[i].Count > slice[j].Count
}

func main() {
	var arg string
	if len(os.Args) != 2 {
		fmt.Println("Improper arguments given")
		os.Exit(0)
	} else {
		arg = os.Args[1]
	}

	file, err := os.Open(arg)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	word_count := make(map[string]int)
	// regex for puncuation
	reg := regexp.MustCompile("[^ '\\-a-zA-Z]+")

	scanner := bufio.NewScanner(file)
	var line_text string
	for scanner.Scan() {
		line_text = reg.ReplaceAllString(scanner.Text(), " ")
		for _, word := range strings.Split(line_text, " ") {
			if word != "" {
				word_count[strings.ToLower(word)]++
			}
		}
	}

	var word_list KeyVals
	for k, v := range word_count {
		word_list = append(word_list, KeyVal{Word: k, Count: v})
	}
	sort.Sort(word_list)
	for _, kv := range word_list {
		fmt.Printf("%s %d\n", kv.Word, kv.Count)
	}
}
