package heap

import (
	"container/heap"
	"fmt"
)

type Item struct {
	priority int
	name     string
}

type ItemList []Item

func (l ItemList) Len() int           { return len(l) }
func (l ItemList) Less(i, j int) bool { return l[i].priority < l[j].priority }
func (l *ItemList) Swap(i, j int)     { (*l)[i], (*l)[j] = (*l)[j], (*l)[i] }
func (l *ItemList) Push(item any)     { *l = append((*l), item.(Item)) }
func (l *ItemList) Pop() any {
	fmt.Printf("%v", l)
	item := (*l)[len(*l)-1]
	*l = (*l)[:len(*l)-1]
	return item
}

func topK(items []Item, k int) []Item {
	itemList := ItemList(items)
	heap.Init(&itemList)
	// why do we need to take pointer ?
	// https://github.com/golang/go/wiki/MethodSets#variables
	// what is the different between value receiver and pointer receiver for interface ?
	topItems := make([]Item, 0, k)
	for i := 0; i < k; i++ {
		item := heap.Pop(&itemList).(Item)
		topItems = append(topItems, item)
	}
	return topItems
}
