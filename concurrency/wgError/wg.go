package wgerror

import (
	"sync"
)

func waitGroupExample(n int) []int {
	var wg sync.WaitGroup
	results := make([]int, n)
	for i := 0; i < len(results); i++ {
		go func(idx int) {
			wg.Add(1)
			results[idx] = idx
			wg.Done()
		}(i)
	}
	wg.Wait()
	return results
}
