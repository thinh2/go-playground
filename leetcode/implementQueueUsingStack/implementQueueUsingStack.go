package implementqueueusingstack

import (
	"container/list"
)

type MyQueue struct {
	pushStack *list.List
	popStack  *list.List
}

func Constructor() MyQueue {
	return MyQueue{
		pushStack: list.New(),
		popStack:  list.New(),
	}
}

func (this *MyQueue) Push(x int) {
	this.pushStack.PushBack(x)
	/*for curr := this.pushStack.Front(); curr != this.pushStack.Back(); curr = curr.Next() {
		fmt.Printf("%v -> ", curr.Value)
	}
	fmt.Println()*/
}

func (this *MyQueue) Pop() int {
	if this.popStack.Len() == 0 {
		this.migratePushToPop()
	}
	val := this.popStack.Back()
	ret := val.Value.(int)
	this.popStack.Remove(val)
	return ret
}

func (this *MyQueue) migratePushToPop() {
	for this.pushStack.Len() != 0 {
		val := this.pushStack.Back()
		this.popStack.PushBack(val.Value.(int))
		this.pushStack.Remove(val)
	}
}

func (this *MyQueue) Peek() int {
	if this.popStack.Len() == 0 {
		this.migratePushToPop()
	}
	return this.popStack.Back().Value.(int)
}

func (this *MyQueue) Empty() bool {
	return this.popStack.Len() == 0 && this.pushStack.Len() == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
