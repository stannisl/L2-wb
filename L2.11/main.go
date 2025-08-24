package main

import (
	"context"
	"fmt"
	"sort"
	"strings"
	"sync"
	"time"
)

const (
	workerSize  = 3
	channelSize = 3
)

func main() {
	ctx := context.Background()
	input := []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик", "стол"}

	start := time.Now()
	result := AsyncParseAnnograms(Generator(ctx, input))
	fmt.Println("done after", time.Since(start))

	for k, v := range result {
		fmt.Printf("- \"%v\": %v\n", k, v)
	}
}

func Generator[T any](ctx context.Context, src []T) <-chan T {
	resCh := make(chan T, channelSize)

	go func() {
		defer close(resCh)

		for _, elem := range src {
			select {
			case <-ctx.Done():
				return
			case resCh <- elem:
			}
		}
	}()

	return resCh
}

func canonicalForm(word string) string {
	runeSeq := []rune(strings.ToLower(word))
	sort.Slice(runeSeq, func(i, j int) bool {
		return runeSeq[i] < runeSeq[j]
	})

	return string(runeSeq)
}

func AsyncParseAnnograms(words <-chan string) map[string][]string {
	parsedWords := make(map[string][]string, 8)
	mu := &sync.Mutex{}
	wg := new(sync.WaitGroup)

	for i := 0; i < workerSize; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			for word := range words {
				letters := canonicalForm(word)

				mu.Lock()
				parsedWords[letters] = append(parsedWords[letters], word)
				mu.Unlock()
			}
		}()
	}

	wg.Wait()

	result := make(map[string][]string, 8)

	for _, v := range parsedWords {
		if len(v) > 1 {
			for i := 0; i < len(v); i++ {
				result[v[0]] = append(result[v[0]], v[i])
			}
		}
	}

	return result
}
