package hw03frequencyanalysis

import (
	"math"
	"regexp"
	"sort"
	"strings"
)

var regexpSplit = regexp.MustCompile(`\s-\s|[!"#$%&'()*+,\\./:;<=>?@[\]^_\x60{|}~\s]+`)

func Top10(text string) []string {
	// Place your code here.
	statMap := make(map[string]int, 0)

	for _, w := range regexpSplit.Split(text, -1) {
		statMap[strings.ToLower(w)]++
	}

	topStats := make([]string, len(statMap))
	for s := range statMap {
		topStats = append(topStats, s)
	}

	sort.Slice(topStats, func(i, j int) bool {
		if statMap[topStats[i]] == statMap[topStats[j]] {
			return topStats[i] < topStats[j]
		}
		return statMap[topStats[i]] > statMap[topStats[j]]
	})

	limit := int(math.Min(float64(len(topStats)), 10))

	return topStats[:limit]
}
