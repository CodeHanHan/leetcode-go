package bit

// Binary Indexed Tree 树状数组
type BinaryIndexedTree struct {
	array []int
}

// NewBinaryIndexedTree 构造一个新的bit
func NewBinaryIndexedTree(list []int) *BinaryIndexedTree {
	arr := make([]int, len(list)+1)
	for i := 0; i < len(list); i++ {
		arr[i+1] = list[i]
	}

	for i := 1; i < len(arr); i++ {
		j := i + (i & -i)
		if j < len(arr) {
			arr[j] += arr[i]
		}
	}

	return &BinaryIndexedTree{
		array: arr,
	}
}

// Update add delta to array[idx]
func (b *BinaryIndexedTree) Update(idx, delta int) {
	idx += 1
	for idx < len(b.array) {
		b.array[idx] += delta
		idx = idx + (idx & -idx)
	}
}

// Prefix return sum of [0:idx]
func (b *BinaryIndexedTree) PrefixSum(idx int) int {
	idx += 1
	res := 0
	for idx > 0 {
		res += b.array[idx]
		idx = idx - (idx & -idx)
	}

	return res
}

// RangeSum return sum of array[from:to]
func (b *BinaryIndexedTree) RangeSum(from, to int) int {
	return b.PrefixSum(to) - b.PrefixSum(from-1)
}
