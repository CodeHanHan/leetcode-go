package heap

func buildMaxHeap(nums []int, n int) {
	for i := n / 2; i >= 1; i-- {
		maxHeapify(nums, n, i)
	}
}

// 对nums的前n个元素进行大顶堆化，开始位置为i
func maxHeapify(nums []int, n, i int) {
	for {
		maxPos := i
		if i*2 <= n && nums[maxPos] < nums[i*2] {
			maxPos = i * 2
		}
		if i*2+1 <= n && nums[maxPos] < nums[i*2+1] {
			maxPos = i*2 + 1
		}

		if maxPos == i {
			break
		}

		nums[i], nums[maxPos] = nums[maxPos], nums[i]
		i = maxPos
	}
}

func sortMax(nums []int, n int) {
	buildMaxHeap(nums, n)
	k := n
	for k > 1 {
		nums[k], nums[1] = nums[1], nums[k]
		k--
		maxHeapify(nums, k, 1)
	}
}
