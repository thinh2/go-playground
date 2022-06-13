package mergeksortedlist

import "testing"

func TestMergeKsortedListTest(t *testing.T) {
	listValues := [][]int{
		{1, 4, 9}, {2, 3, 10}, {5, 7},
	}
	expected := []int{1, 2, 3, 4, 5, 7, 9, 10}
	kSortedList := buildKSortedList(listValues)
	output := mergeKsortedList(kSortedList)
	for _, expectedVal := range expected {
		if output == nil {
			t.Errorf("output length is not same with expected")
		}
		if output.val != expectedVal {
			t.Errorf("merge list error, expected %v, received %v", expectedVal, output.val)
		}
		output = output.next
	}
}

func buildKSortedList(valList [][]int) []*Node {
	kList := make([]*Node, 0, len(valList))
	for _, vals := range valList {
		kList = append(kList, buildList(vals))
	}
	return kList
}

func buildList(vals []int) *Node {
	var (
		head *Node
		curr *Node
	)

	for _, val := range vals {
		node := &Node{
			val:  val,
			next: nil,
		}
		if head == nil {
			head = node
			curr = node
		} else {
			curr.next = node
			curr = curr.next
		}
	}
	return head
}
