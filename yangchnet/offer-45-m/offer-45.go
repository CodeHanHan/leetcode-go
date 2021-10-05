package offer_45_m_

import (
	"bytes"
	"sort"
	"strconv"
)

func minNumber(nums []int) string {
	var arr Arr = nums
	sort.Sort(arr)
	var buffer bytes.Buffer
	for _, v := range arr {
		buffer.WriteString(strconv.Itoa(v))
	}
	return buffer.String()
}

type Arr []int

func (a Arr) Len() int {
	return len(a)
}

func (a Arr) Less(i, j int) bool {
	strI := strconv.Itoa(a[i])
	strJ := strconv.Itoa(a[j])
	x := strI + strJ
	y := strJ + strI
	numX, _ := strconv.Atoi(x)
	numY, _ := strconv.Atoi(y)
	return numX < numY
}

func (a Arr) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
