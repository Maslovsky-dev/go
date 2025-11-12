package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

func Top10(s string) []string {
	freqMap := map[string]int{}

	for _, word := range strings.Fields(s) {
		freqMap[word]++
	}

	type kv struct {
		Key   string
		Value int
	}

	// Создаем слайл для пар kv
	pairs := make([]kv, 0, len(freqMap))
	for k, v := range freqMap {
		pairs = append(pairs, kv{k, v})
	}

	sort.Slice(pairs, func(i, j int) bool {
		if pairs[i].Value == pairs[j].Value {
			return pairs[i].Key < pairs[j].Key // по ключу asc
		}
		return pairs[i].Value > pairs[j].Value // по значению desc
	})

	// Определяем размер итогового слайса
	n := len(pairs)
	if n > 10 {
		n = 10
	}

	result := make([]string, 0, n)
	for i := 0; i < n; i++ {
		result = append(result, pairs[i].Key)
	}

	return result
}
