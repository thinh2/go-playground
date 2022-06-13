package heap

import "testing"

func TestTopK(t *testing.T) {
	input := []Item{
		{name: "a", priority: 10}, {name: "b", priority: 23}, {name: "c", priority: 9},
		{name: "d", priority: 48}, {name: "e", priority: 7},
	}
	k := 3
	expected := []Item{
		{name: "e", priority: 7}, {name: "c", priority: 9}, {name: "a", priority: 10},
	}
	output := topK(input, k)
	for i := 0; i < len(output); i++ {
		if expected[i] != output[i] {
			t.Errorf("failed to get from priority, expected %v, return %v", expected, output)
			return
		}
	}
}
