package hw03frequencyanalysis

import (
	"regexp"
	"sort"
	"strings"
)

type Counter struct {
	word  string
	count int
}

var wordOnly = regexp.MustCompile(`\p{L}+.*\p{L}+|\p{L}+|^[\p{P}\p{S}]{2,}$`)

func substitute2(in string) string {
	out := wordOnly.FindString(in)
	out = strings.ToLower(out)
	return out
}

func Top10(text string) []string {
	words := strings.Fields(text)
	table := make(map[string]int, 10)
	for _, w := range words {
		key := substitute2(w)
		if key == "" {
			continue
		}
		table[key]++
	}
	sorted := make([]Counter, 0, len(table))
	for w, c := range table {
		sorted = append(sorted, Counter{word: w, count: c})
	}
	sort.Slice(sorted, func(i, j int) bool {
		if sorted[i].count == sorted[j].count { // сравнение строк при совпадении сётчиков
			return strings.Compare(sorted[i].word, sorted[j].word) < 0
		}
		return sorted[i].count > sorted[j].count // '>' большие в начале
	})
	result := []string{}
	for i := 0; i < 10 && i < len(sorted); i++ {
		result = append(result, sorted[i].word)
	}
	return result
}
