package offer_41_d_

/* ============== 这种写法超出时间限制==============
type MedianFinder struct {
	Data []int
	Len int
}


 // initialize your data structure here.
func Constructor() MedianFinder {
	return MedianFinder{
		Data:   make([]int, 0),
		Len:    0,
	}
}

func (this *MedianFinder) AddNum(num int)  {
	this.Data = append(this.Data, num)
	sort.Ints(this.Data)    // 应该是这里超了，每次都要排序，可以优化为找出中位数前排序
	this.Len ++
}

func (this *MedianFinder) FindMedian() float64 {
	if this.Len == 1 {
		return float64(this.Data[0])
	}
	var odd bool = (this.Len % 2) == 0
	if odd {
		mid := this.Len / 2
		return float64(this.Data[mid-1] + this.Data[mid]) / 2
	}
	mid := int(this.Len / 2)
	return float64(this.Data[mid])
}

*/

/*
大根堆：存放数据流中较小的一半元素。
小根堆：存放数据流中较大的一半元素。

让 大根堆的长度 = 小根堆的长度 或者 大根的堆长度 + 1 = 小根堆的长度

*/

// 下面的代码有个bug，没找出来。。

type MedianFinder struct {
	minArr []int
	maxArr []int
}

func Constructor() MedianFinder {
	heap := MedianFinder{
		minArr: make([]int, 0),
		maxArr: make([]int, 0),
	}

	return heap
}

func swap(heapArr []int, i, j int) {
	tmp := heapArr[i]
	heapArr[i] = heapArr[j]
	heapArr[j] = tmp
}

// 小根堆添加元素后更新，元素上浮
func (this *MedianFinder) minHeapUp(val int) {
	this.minArr = append(this.minArr, val)
	i := len(this.minArr) - 1
	for i > 0 && this.minArr[i] < this.minArr[(i-1)/2] {
		swap(this.minArr, i, (i-1)/2)
		i = (i - 1) / 2
	}
}

// 大根堆添加元素后更新，元素上浮
func (this *MedianFinder) maxHeapUp(val int) {
	this.maxArr = append(this.maxArr, val)
	i := len(this.maxArr) - 1
	for i > 0 && this.maxArr[i] > this.maxArr[(i-1)/2] {
		swap(this.maxArr, i, (i-1)/2)
		i = (i - 1) / 2
	}
}

// 小根堆更新堆顶元素后的操作
func (this *MedianFinder) minHeapDown(val int) {
	index := 0
	lchild := index*2 + 1 // left child

	minChild := 0
	this.minArr[0] = val

	for lchild < len(this.minArr) {
		// 找出左右孩子中较小的
		if lchild+1 < len(this.minArr) && this.minArr[lchild] > this.minArr[lchild+1] {
			minChild = lchild + 1 // right child
		} else {
			minChild = lchild
		}

		if this.minArr[index] < this.minArr[minChild] { // 当前值小于其孩子中的较小值
			minChild = index

		}

		if index == minChild {
			break
		}

		swap(this.minArr, index, minChild) // 将当前值与其孩子中较小值互换

		// 更新指针
		index := minChild
		lchild = index*2 + 1
	}
}

// 大根堆更新堆顶元素后的操作
func (this *MedianFinder) maxHeapDown(val int) {
	index := 0
	lchild := index*2 + 1

	maxChild := 0
	this.maxArr[0] = val

	for lchild < len(this.maxArr) {
		if lchild+1 < len(this.maxArr) && this.maxArr[lchild] < this.maxArr[lchild+1] {
			maxChild = lchild + 1
		} else {
			maxChild = lchild
		}

		if this.maxArr[index] > this.maxArr[maxChild] {
			maxChild = index

		}

		if index == maxChild {
			break
		}

		swap(this.maxArr, index, maxChild)

		index = maxChild
		lchild = index*2 + 1
	}
}

func (this *MedianFinder) AddNum(num int) {
	if len(this.minArr) == 0 {
		this.minHeapUp(num)
		return
	}

	//元素比小根堆的堆顶元素大
	if num > this.minArr[0] {
		// 判断如果将该元素扔到小根堆时，如果此时小根堆长度比大根堆大1
		if len(this.minArr)+1-len(this.maxArr) > 1 {
			this.maxHeapUp(this.minArr[0]) // 将小根堆的堆顶元素扔到大根堆
			this.minHeapDown(num)          // 该元素放在小根堆的堆顶
		} else {
			this.minHeapUp(num) // 直接扔进小根堆
		}
	} else { // 元素比小根堆的堆顶元素小，需要扔到大根堆
		if len(this.maxArr)+1 > len(this.minArr) { // 如果将该元素扔到大根堆后，长度比小根堆要大，则不能直接扔进大根堆
			if num < this.maxArr[0] { // 该元素比大根堆的堆顶元素小
				this.minHeapUp(this.maxArr[0]) // 将大根堆的堆顶元素扔进小根堆
				this.maxHeapDown(num)          // 将该元素放在大根堆堆顶
			} else {
				this.minHeapUp(num) // 该元素比大根堆的堆顶元素大，直接扔到小根堆
			}
		} else {
			this.maxHeapUp(num)
		}
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if len(this.minArr) == len(this.maxArr) {
		return (float64(this.minArr[0]) + float64(this.maxArr[0])) / 2
	} else {
		return float64(this.minArr[0])
	}
}
