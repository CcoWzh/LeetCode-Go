package main

import "fmt"

type papapa []int

func main() {
	var head papapa
	head = append(head, 0)

	for i := 0; i < 100; i++ {
		head.insert(i)
	}
	for i := 0; i < 100; i++ {
		fmt.Println(head.deleteMax())
		fmt.Println(head)
	}
}

//insert 插入元素
func (this *papapa) insert(key int) {
	*this = append(*this, key)
	this.swim(len(*this) - 1)
}

func (this *papapa) deleteMax() int {
	(*this)[1], (*this)[len(*this)-1] = (*this)[len(*this)-1], (*this)[1]
	res := (*this)[len(*this)-1]
	(*this) = (*this)[:len(*this)-1]
	this.sink(1)
	return res
}

//swim 上浮
func (this *papapa) swim(index int) {
	if index/2 == 0 {
		return
	}
	if (*this)[index] > (*this)[index/2] {
		(*this)[index], (*this)[index/2] = (*this)[index/2], (*this)[index]
		this.swim(index / 2)
	}
	return
}

//sink 下沉
func (this *papapa) sink(index int) {
	maxIndex := len(*this) - 1
	if maxIndex >= 2*index+1 {
		//找出子节点最大的
		max := 2 * index
		if (*this)[2*index+1] > (*this)[2*index] {
			max = 2*index + 1
		}
		if (*this)[index] < (*this)[max] {
			(*this)[index], (*this)[max] = (*this)[max], (*this)[index]
			this.sink(max)
		}
	} else if maxIndex >= 2*index {
		if (*this)[index] < (*this)[2*index] {
			(*this)[index], (*this)[2*index] = (*this)[2*index], (*this)[index]
			this.sink(2 * index)
		}
	}
	return
}
