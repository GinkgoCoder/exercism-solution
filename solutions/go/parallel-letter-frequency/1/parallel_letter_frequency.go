package letter

import "sync"

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(text string) FreqMap {
	frequencies := FreqMap{}
	for _, r := range text {
		frequencies[r]++
	}
	return frequencies
}

// ConcurrentFrequency counts the frequency of each rune in the given strings,
// by making use of concurrency.
func ConcurrentFrequency(texts []string) FreqMap {
	c := make(chan []any, 100)
	var wg sync.WaitGroup

	for _, text := range texts {
		wg.Add(1)
		go func(text string) {
			defer wg.Done()
			freq := map[rune]int{}
			for _, r := range text {
				freq[r]++
			}
			for r, count := range freq {
				c <- []any{r, count}
			}
		}(text)
	}
	go func() {
		wg.Wait()
		close(c)
	}()
	frequecies := FreqMap{}
	for freq := range c {
		frequecies[freq[0].(rune)] += freq[1].(int)
	}
	return frequecies
}
