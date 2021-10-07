package solution0232

type MyQueue struct {
	stack    []int
	auxstack []int
}

func Constructor() MyQueue {
	return MyQueue{[]int{}, []int{}}
}

func (this *MyQueue) Push(x int) {
	// transfer everything to auxstack
	for i := len(this.stack) - 1; i >= 0; i-- {
		this.auxstack = append(this.auxstack, this.stack[i])
	}

	// add x to bottom of stack
	this.stack = []int{}
	this.stack = append(this.stack, x)

	// return everything from auxstack to stack
	for i := len(this.auxstack) - 1; i >= 0; i-- {
		this.stack = append(this.stack, this.auxstack[i])
	}
	this.auxstack = []int{}
}

func (this *MyQueue) Pop() int {
	if len(this.stack) == 0 {
		return 0
	}
	x := this.stack[len(this.stack)-1]
	this.stack = this.stack[:len(this.stack)-1]
	return x
}

func (this *MyQueue) Peek() int {
	if len(this.stack) == 0 {
		return 0
	}
	return this.stack[len(this.stack)-1]
}

func (this *MyQueue) Empty() bool {
	return len(this.stack) == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
