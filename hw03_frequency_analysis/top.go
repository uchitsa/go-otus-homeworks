package hw03frequencyanalysis

import "strings"

func Top10(text string) []string {
	// Place your code here.
	statMap := make(map[string]int, 0)
	words := strings.Split(text, " ")

	for _, word := range words {
		w := strings.Trim(word, "?!,.-")
		statMap[strings.ToLower(w)]++
	}

	delete(statMap, "")

	topStats := make([]string, len(statMap))

	return topStats[:10]
}
