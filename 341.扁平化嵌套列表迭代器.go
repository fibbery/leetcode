/*
 * @lc app=leetcode.cn id=341 lang=golang
 *
 * [341] 扁平化嵌套列表迭代器
 */
package main

type NestedInteger struct {
	Val  int
	Data []*NestedInteger
}

func (this NestedInteger) IsInteger() bool {
	return len(this.Data) == 0 && this.Val > 0
}

func (this NestedInteger) GetInteger() int {
	return this.Val
}

func (this *NestedInteger) SetInteger(value int) {
	this.Val = value
}

func (this *NestedInteger) Add(elem NestedInteger) {
	this.Data = append(this.Data, &elem)
}

func (this *NestedInteger) GetList() []*NestedInteger {
	return this.Data
}

// @lc code=start
/**
 * // This is the interface that allows for creating nested lists.
 * // You should not implement it, or speculate about its implementation
 * type NestedInteger struct {
 * }
 *
 * // Return true if this NestedInteger holds a single integer, rather than a nested list.
 * func (this NestedInteger) IsInteger() bool {}
 *
 * // Return the single integer that this NestedInteger holds, if it holds a single integer
 * // The result is undefined if this NestedInteger holds a nested list
 * // So before calling this method, you should have a check
 * func (this NestedInteger) GetInteger() int {}
 *
 * // Set this NestedInteger to hold a single integer.
 * func (n *NestedInteger) SetInteger(value int) {}
 *
 * // Set this NestedInteger to hold a nested list and adds a nested integer to it.
 * func (this *NestedInteger) Add(elem NestedInteger) {}
 *
 * // Return the nested list that this NestedInteger holds, if it holds a nested list
 * // The list length is zero if this NestedInteger holds a single integer
 * // You can access NestedInteger's List element directly if you want to modify it
 * func (this NestedInteger) GetList() []*NestedInteger {}
 */

type NestedIterator struct {
	data    []int
	current int
}

func Constructor(nestedList []*NestedInteger) *NestedIterator {
	data := make([]int, 0)
	for _, value := range nestedList {
		data = append(data, traverse(value)...)
	}
	return &NestedIterator{data: data, current: -1}
}

func traverse(value *NestedInteger) []int {
	data := make([]int, 0)
	if value.IsInteger() {
		return append(data, value.GetInteger())
	} else {
		for _, item := range value.GetList() {
			data = append(data, traverse(item)...)
		}
	}
	return data
}

func (this *NestedIterator) Next() int {
	this.current++
	return this.data[this.current]
}

func (this *NestedIterator) HasNext() bool {
	return this.current+1 < len(this.data)
}

// @lc code=end
