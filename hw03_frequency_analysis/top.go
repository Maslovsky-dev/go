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

	var pairs []kv
	for k, v := range freqMap {
		pairs = append(pairs, kv{k, v})
	}

	sort.Slice(pairs, func(i, j int) bool {
		if pairs[i].Value == pairs[j].Value {
			return pairs[i].Key < pairs[j].Key // по ключу asc
		}
		return pairs[i].Value > pairs[j].Value // по значению desc
	})

	var result []string

	for _, pair := range pairs {
		result = append(result, pair.Key)
	}

	if len(result) > 10 {
		return result[:10]
	}

	return result
}
