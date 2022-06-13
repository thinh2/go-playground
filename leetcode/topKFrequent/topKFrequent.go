package topKFrequent

import (
	"container/heap"
	"fmt"
)

type Element struct {
	freq int
	val  int
}

type PriorityQueue []Element

func (pq *PriorityQueue) Len() int           { return len(*pq) }
func (pq *PriorityQueue) Swap(i, j int)      { (*pq)[i], (*pq)[j] = (*pq)[j], (*pq)[i] }
func (pq *PriorityQueue) Less(i, j int) bool { return (*pq)[i].freq < (*pq)[i].freq }
func (pq *PriorityQueue) Push(x any)         { *pq = append(*pq, x.(Element)) }
func (pq *PriorityQueue) Pop() any {
	n := pq.Len()
	item := (*pq)[n-1]
	*pq = (*pq)[:n-1]
	fmt.Println(*pq)
	return item
}

func topKFrequent(nums []int, k int) []int {
	pq := PriorityQueue{}
	freqCount := make(map[int]int)
	for _, num := range nums {
		freqCount[num]++
	}

	fmt.Println(freqCount)
	for num, freq := range freqCount {
		if pq.Len() < k {
			heap.Push(&pq, Element{freq: freq, val: num})
		} else if pq[0].freq < freq {
			heap.Pop(&pq)
			heap.Push(&pq, Element{freq: freq, val: num})
		}
	}

	topK := []int{}
	for {
		if pq.Len() == 0 {
			break
		}
		elem := heap.Pop(&pq).(Element)
		topK = append(topK, elem.val)
	}
	return topK
}
