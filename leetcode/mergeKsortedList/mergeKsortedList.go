package mergeksortedlist

import "container/heap"

type Node struct {
	val  int
	next *Node
}

type PriorityQueue []*Node

func (pq *PriorityQueue) Len() int           { return len(*pq) }
func (pq *PriorityQueue) Swap(i, j int)      { (*pq)[i], (*pq)[j] = (*pq)[j], (*pq)[i] }
func (pq *PriorityQueue) Less(i, j int) bool { return (*pq)[i].val < (*pq)[j].val }
func (pq *PriorityQueue) Push(x any)         { *pq = append(*pq, x.(*Node)) }
func (pq *PriorityQueue) Pop() any {
	n := len(*pq)
	item := (*pq)[n-1]
	*pq = (*pq)[:n-1]
	return item
}

func mergeKsortedList(kSortedList []*Node) *Node {
	pq := PriorityQueue{}
	var head *Node
	var curr *Node
	for _, elem := range kSortedList {
		if elem != nil {
			heap.Push(&pq, elem)
		}
	}

	for {
		if pq.Len() == 0 {
			break
		}
		elem := heap.Pop(&pq).(*Node)

		if head == nil {
			head = elem
			curr = head
		} else {
			curr.next = elem
			curr = curr.next
		}
		if elem.next != nil {
			heap.Push(&pq, elem.next)
		}
	}
	return head
}
