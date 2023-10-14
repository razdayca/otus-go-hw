package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

func Top10(text string) []string {
	if len(text) == 0 {
		return nil
	}
	counter := map[string]int{}
	splitText := strings.Fields(text)
	for _, val := range splitText {
		value, ok := counter[val]
		if ok {
			counter[val] = value + 1
		} else {
			counter[val] = 1
		}
	}
	intervalMap := map[int][]string{}
	for i := 0; i < 10; i++ {
		var largestValue int
		var largestkey string
		for key, value := range counter {
			if key != "" {
				if value > largestValue {
					largestValue = value
					largestkey = key
				}
			}
		}
		for key, val := range counter {
			if largestkey == key {
				intervalMap[val] = append(intervalMap[val], key)
				sort.Strings(intervalMap[val])
			}
		}
		delete(counter, largestkey)
	}
	values := make([]string, 0, 10)
	for len(intervalMap) > 0 {
		var largestkey int
		for key := range intervalMap {
			if key > largestkey {
				largestkey = key
			}
		}
		values = append(values, intervalMap[largestkey]...)
		delete(intervalMap, largestkey)
	}
	return values
}
