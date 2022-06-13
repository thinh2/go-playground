package exchange

type edge struct {
	name string
	val  float32
}

type graph struct {
	adjList map[string][]edge
}

func (g *graph) addEdge(s, t string, val float32) {
	g.adjList[s] = append(g.adjList[s], edge{t, val})
}

func (g *graph) findBestValue(s, t string) float32 {
	visited := map[string]bool{}
	visited[s] = true
	return g.dfs(s, t, visited)
}

func (g *graph) dfs(s, t string, visited map[string]bool) float32 {
	if s == t {
		return float32(1.0)
	}
	maxRate := float32(-1.0)
	for _, adj := range g.adjList[s] {
		if visited[adj.name] {
			continue
		}
		visited[adj.name] = true
		tmp := g.dfs(adj.name, t, visited) * adj.val
		if tmp > maxRate {
			maxRate = tmp
		}
		visited[adj.name] = false
	}
	return maxRate
}
