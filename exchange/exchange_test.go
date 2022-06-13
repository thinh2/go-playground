package exchange

import "testing"

func TestExchange(t *testing.T) {
	edgeList := map[string][]edge{
		"usd": {{"vnd", 10.1}, {"sgd", 1.3}, {"cny", 2.1}},
		"vnd": {{"sgd", 2}, {"usd", 3}},
	}
	graph := &graph{
		adjList: edgeList,
	}
	val := graph.findBestValue("usd", "sgd")
	if val < float32(20) {
		t.Errorf("compute exchange failed, output %v", val)
	}
}
