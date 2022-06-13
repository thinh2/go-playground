package iter

import (
	"reflect"
	"testing"
)

func TestConver2d(t *testing.T) {
	input := [][]int{
		{1, 2, 3}, {4, 5}, {6}, {}, {7, 8, 9},
	}
	expected := []int{1, 4, 6, 7, 2, 5, 8, 3, 9}
	output := convert2d(input)
	if !reflect.DeepEqual(output, expected) {
		t.Errorf("failed to convert, expected %v, output %v", expected, output)
	}
}

func TestSlice2DIterator(t *testing.T) {
	input := [][][]int{
		{{1, 2, 3}, {4, 5}, {6}, {}, {7, 8, 9}},
		{{}, {1, 2, 3}, {4, 5}, {6}, {}, {7, 8, 9}},
	}
	expected := [][]int{
		{1, 4, 6, 7, 2, 5, 8, 3, 9},
		{1, 4, 6, 7, 2, 5, 8, 3, 9},
	}
	for id := range input {
		iterator := NewSlice2DIterator(input[id])
		iterator.Init()
		idx := 0
		for !iterator.HasNext() {
			val := iterator.Next()
			t.Log(val)
			if val != expected[id][idx] {
				t.Errorf("iterator failed, expected %v at index %v, output %v", expected[idx], idx, val)
			}
			idx++
		}
	}
}

func TestRangeIterator(t *testing.T) {
	rangeIterator := NewRangeIterator(2, 12, 2)
	expected := []int{2, 4, 6, 8, 10}
	idx := 0
	for !rangeIterator.HasNext() {
		val := rangeIterator.Next()
		if val != expected[idx] {
			t.Errorf("iterator failed, expected %v at index %v, output %v", expected[idx], idx, val)
		}
		idx++
	}
}

func TestRotationIterator(t *testing.T) {
	rangeIterator := NewRangeIterator(2, 12, 2)
	rotaionIterator := newRotationIterator(rangeIterator)
	rotaionIterator.Init()
	expected := []int{2, 4, 6, 8, 10}
	for i := 0; i < len(expected)*4; i++ {
		val := rotaionIterator.Next()
		if val != expected[i%len(expected)] {
			t.Errorf("iterator failed, expected %v at index %v, output %v", expected[i%len(expected)], i, val)
		}
	}
}

func TestRotationIterator2(t *testing.T) {
	input := [][]int{
		{1, 2, 3}, {4, 5}, {6}, {}, {7, 8, 9},
	}
	expected := []int{1, 4, 6, 7, 2, 5, 8, 3, 9}

	sliceIterator := NewSlice2DIterator(input)
	rotationIterator := newRotationIterator(sliceIterator)
	rotationIterator.Init()
	for i := 0; i < len(expected)*4; i++ {
		val := rotationIterator.Next()
		if val != expected[i%len(expected)] {
			t.Errorf("iterator failed, expected %v at index %v, output %v", expected[i%len(expected)], i, val)
		}
	}
}
