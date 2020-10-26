package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	set := Constructor()
	set.Insert(0)
	set.Insert(1)
	set.Remove(0)
	set.Insert(2)
	set.Remove(1)
	r := set.GetRandom()
	fmt.Println(r)
}

/*
 * @lc app=leetcode.cn id=380 lang=golang
 *
 * [380] 常数时间插入、删除和获取随机元素
 */

// @lc code=start
type RandomizedSet struct {
	values      []int
	value2Index map[int]int
}

/** Initialize your data structure here. */
func Constructor() RandomizedSet {
	return RandomizedSet{
		value2Index: make(map[int]int),
	}
}

/** Inserts a value to the set. Returns true if the set did not already contain the specified element. */
func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.value2Index[val]; !ok {
		this.values = append(this.values, val)
		this.value2Index[val] = len(this.values) - 1
		return true
	}
	return false
}

/** Removes a value from the set. Returns true if the set contained the specified element. */
func (this *RandomizedSet) Remove(val int) bool {
	if idx, ok := this.value2Index[val]; ok {
		//swap with last element
		this.swap(idx, len(this.values)-1)
		this.values = this.values[:len(this.values)-1]
		delete(this.value2Index, val)
		return true
	}
	return false
}

/** Get a random element from the set. */
func (this *RandomizedSet) GetRandom() int {
	rand.Seed(time.Now().UnixNano())
	idx := rand.Intn(len(this.values))
	return this.values[idx]
}

func (this *RandomizedSet) swap(a, b int) {
	this.values[a], this.values[b] = this.values[b], this.values[a]
	this.value2Index[this.values[a]] = a
	this.value2Index[this.values[b]] = b
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
// @lc code=end
