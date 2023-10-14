package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

func Top10(text string) []string { //nolint:gocognit
	if len(text) > 0 {
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
		test := map[int][]string{}
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
					test[val] = append(test[val], key)
					sort.Strings(test[val])
				}
			}
			delete(counter, largestkey)
		}
		values := make([]string, 0, 10)
		for len(test) > 0 {
			var largestkey int
			for key := range test {
				if key > largestkey {
					largestkey = key
				}
			}
			values = append(values, test[largestkey]...)
			delete(test, largestkey)
		}
		return values
	}
	return nil
}
