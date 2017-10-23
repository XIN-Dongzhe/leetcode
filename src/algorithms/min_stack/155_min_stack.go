package min_stack

// https://leetcode.com/problems/min-stack/description/

type MinStack struct {
	data     []Item
	min_data int
}

type Item struct {
	data     int
	min_data int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		data: []Item{},
	}
}

func (this *MinStack) Push(x int) {
	if x < this.min_data || 0 == len(this.data) {
		this.min_data = x
	}

	this.data = append(this.data, Item{data: x, min_data: this.min_data})
}

func (this *MinStack) Pop() {
	if 0 == len(this.data) {
		return
	}

	this.data = append(this.data[:len(this.data)-1])

	if 0 == len(this.data) {
		return
	}

	this.min_data = this.data[len(this.data)-1].min_data
}

func (this *MinStack) Top() int {
	return this.data[len(this.data)-1].data
}

func (this *MinStack) GetMin() int {
	return this.min_data
}
