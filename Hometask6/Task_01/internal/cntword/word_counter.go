package cntword

import (
	"bufio"
	"os"
	"slices"
)

type Word struct {
	Word  string
	Count int
}

func CountWords(filePath string) (res map[string]int, err error) {
	res = make(map[string]int)
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer func() {
		err = file.Close()
	}()
	s := bufio.NewScanner(file)
	s.Split(bufio.ScanWords)
	for s.Scan() {
		if s.Err() != nil {
			return nil, s.Err()
		}
		res[s.Text()]++
	}
	return res, err
}

func SortWords(words map[string]int) []Word {
	result := make([]Word, 0, len(words))
	for w, n := range words {
		result = append(result, Word{w, n})
	}
	slices.SortFunc(result, func(a, b Word) int {
		if a.Count > b.Count {
			return -1
		}
		if a.Count < b.Count {
			return 1
		}
		return 0
	})
	return result
}
