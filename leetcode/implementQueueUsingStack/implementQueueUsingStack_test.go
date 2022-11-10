package implementqueueusingstack

import (
	"fmt"
	"testing"
)

func TestQueue(t *testing.T) {
	queue := Constructor()
	queue.Push(1)
	queue.Push(2)
	fmt.Println("pop:", queue.Pop())
	queue.Push(3)
	fmt.Println("pop:", queue.Pop())
}
