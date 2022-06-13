package topKFrequent

import (
	"sort"
	"testing"
)

func TestTopKFrequent(t *testing.T) {
	//[4,1,-1,2,-1,2,3], 2
	inputs := []int{1, 2, 1, 1, 3, 4, 4, 2, 2}
	topK := 2
	output := topKFrequent(inputs, topK)
	expected := []int{1, 2}
	sort.Ints(output)
	for i := range output {
		if output[i] != expected[i] {
			t.Errorf("failed to get topKFrequent, expected %v, received %v", expected, output)
		}
	}
}

func TestTopKFrequent2(t *testing.T) {
	inputs := []int{5, 2, 5, 3, 5, 3, 1, 1, 3}
	topK := 2
	output := topKFrequent(inputs, topK)
	expected := []int{3, 5}
	sort.Ints(output)

	//maybe test cache problem
	for i := range output {
		//fmt.Println(output[i], expected[i])
		if output[i] != expected[i] {
			t.Errorf("failed to get topKFrequent, expected %v, received %v", expected, output)
		}
	}
}
