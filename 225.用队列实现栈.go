/*
 * @lc app=leetcode.cn id=225 lang=golang
 *
 * [225] 用队列实现栈
 */

// @lc code=start
type MyStack struct {
	Vals []int
}


/** Initialize your data structure here. */
func Constructor() MyStack {
	return MyStack{Vals : make([]int,0)}
}


/** Push element x onto stack. */
func (this *MyStack) Push(x int)  {
	this.Vals = append(this.Vals, x)
}


/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	length := len(this.Vals)
	if length > 0 {
		ret := this.Vals[length - 1]
		this.Vals = this.Vals[:length - 1]
		return ret
	}else{
		return 0
	}
}


/** Get the top element. */
func (this *MyStack) Top() int {
	length := len(this.Vals)
	if length > 0 {
		return this.Vals[length - 1]
	}else{
		return 0
	}
}


/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	return len(this.Vals) == 0
}


/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
// @lc code=end

