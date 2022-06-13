package wgerror

import (
	"testing"
)

func TestConcurrency(t *testing.T) {
	results := waitGroupExample(10000)
	for i, val := range results {
		if val != i {
			t.Error("concurrency error, failed to update at index ", i)
		}
	}

}
