package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

type Frequency struct {
	count int
	word  string
}

func Top10(s string) []string {
	if len(s) == 0 {
		return nil
	}

	tokenCount := make(map[string]int, len(s))
	for _, token := range strings.Fields(s) {
		tokenCount[token]++
	}

	frequencies := make([]Frequency, 0)
	for word, count := range tokenCount {
		frequencies = append(frequencies, Frequency{count, word})
	}

	sort.Slice(frequencies, func(i, j int) bool {
		return frequencies[i].count > frequencies[j].count
	})

	var result []string
	var currCount, nextCount int
	for n := 0; n < 10; {
		currCount = frequencies[n].count
		nextCount = frequencies[n+1].count

		if currCount > nextCount && n < 10 {
			result = append(result, frequencies[n].word)
			n++
			continue
		}

		var equalCounts []string
		prevCount := currCount
		for currCount == prevCount && n < 10 {
			equalCounts = append(equalCounts, frequencies[n].word)
			n++
			prevCount = currCount
			currCount = frequencies[n].count
		}
		sort.Strings(equalCounts)

		result = append(result, equalCounts...)
	}

	return result
}
