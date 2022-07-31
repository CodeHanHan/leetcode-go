package queue

// 单调队列

type MonotonicQueue struct {
	queue []int
}

func NewMyQueue() *MonotonicQueue {
	return &MonotonicQueue{
		queue: make([]int, 0),
	}
}

func (m *MonotonicQueue) Front() int {
	return m.queue[0]
}

func (m *MonotonicQueue) Back() int {
	return m.queue[len(m.queue)-1]
}

func (m *MonotonicQueue) Empty() bool {
	return len(m.queue) == 0
}

func (m *MonotonicQueue) Push(val int) {
	// 每次入队时，将比自己小的值kill掉
	for !m.Empty() && val > m.Back() {
		m.queue = m.queue[:len(m.queue)-1]
	}
	m.queue = append(m.queue, val)
}

func (m *MonotonicQueue) Pop(val int) {
	// 每次弹出的时候，比较当前要弹出的数值是否等于队列出口元素的数值，如果相等则弹出。
	if !m.Empty() && val == m.Front() {
		m.queue = m.queue[1:]
	}
}
