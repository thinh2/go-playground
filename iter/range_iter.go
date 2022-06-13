package iter

import "container/list"

type Iterator interface {
	HasNext() bool
	Next() int
	Init()
}

type RangeIterator struct {
	start int
	end   int
	step  int
	curr  int
}

func NewRangeIterator(start, end, step int) *RangeIterator {
	return &RangeIterator{
		start: start,
		end:   end,
		step:  step,
		curr:  start,
	}
}

func (iter *RangeIterator) Init() {
	iter.curr = iter.start
}

func (iter *RangeIterator) HasNext() bool {
	return (iter.curr+iter.step <= iter.end)
}

func (iter *RangeIterator) Next() int {
	defer func() {
		iter.curr += iter.step
	}()
	return iter.curr
}

type element struct {
	row int
	col int
}

type Slice2DIterator struct {
	slice2d *[][]int
	queue   *list.List
}

func NewSlice2DIterator(slice2d [][]int) *Slice2DIterator {
	return &Slice2DIterator{
		slice2d: &slice2d,
	}
}

func (iter *Slice2DIterator) Init() {
	iter.queue = list.New()
	for i := 0; i < len(*iter.slice2d); i++ {
		if len((*iter.slice2d)[i]) > 0 {
			iter.queue.PushFront(element{i, 0})
		}
	}
}

func (iter *Slice2DIterator) HasNext() bool {
	return iter.queue.Len() != 0
}

func (iter *Slice2DIterator) Next() int {
	front := iter.queue.Back()
	ele := front.Value.(element)
	iter.queue.Remove(front)
	ret := (*iter.slice2d)[ele.row][ele.col]
	if ele.col+1 < len((*iter.slice2d)[ele.row]) {
		iter.queue.PushFront(element{ele.row, ele.col + 1})
	}
	return ret
}

type RotationIterator struct {
	baseIterator Iterator
}

func (iter *RotationIterator) Init() {
	iter.baseIterator.Init()
}

func (iter *RotationIterator) Next() int {
	if !iter.baseIterator.HasNext() {
		iter.baseIterator.Init()
	}
	ret := iter.baseIterator.Next()
	return ret
}

func (iter *RotationIterator) HasNext() bool {
	return true
}

func newRotationIterator(baseIterator Iterator) *RotationIterator {
	return &RotationIterator{
		baseIterator: baseIterator,
	}
}
