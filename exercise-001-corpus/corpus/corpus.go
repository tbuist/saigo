package corpus

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"
)

// KeyVal : pair for use with sorting
type KeyVal struct {
	Word  string
	Count int
}

// KeyVals : slice for use with sorting
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
	if slice[i].Count != slice[j].Count {
		return slice[i].Count > slice[j].Count
	}

	for k := 0; k < len(slice[i].Word) && k < len(slice[j].Word); k++ {
		if []rune(slice[i].Word)[k] != []rune(slice[j].Word)[k] {
			return []rune(slice[i].Word)[k] < []rune(slice[j].Word)[k]
		} else if k == len(slice[i].Word)-1 {
			return true
		} else if k == len(slice[j].Word)-1 {
			return false
		}
	}
	// needed, but should never get to this point
	return true
}

// WordCount : prints words and counts given a text file
func WordCount(filename string) KeyVals {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	defer file.Close()

	wordMap := make(map[string]int)
	// regex for puncuation
	reg := regexp.MustCompile("[^ '\\-a-zA-Z]+")

	scanner := bufio.NewScanner(file)
	var lineText string
	for scanner.Scan() {
		lineText = reg.ReplaceAllString(scanner.Text(), " ")
		for _, word := range strings.Split(lineText, " ") {
			if word != "" {
				wordMap[strings.ToLower(word)]++
			}
		}
	}

	var wordList KeyVals
	for k, v := range wordMap {
		wordList = append(wordList, KeyVal{Word: k, Count: v})
	}
	sort.Sort(wordList)
	return wordList
}
