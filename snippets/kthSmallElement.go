package snippets

import (
	"golang.org/x/exp/constraints"
)

// 使用二分法从两个正序数组中找到最小的第k个数
// 基本思路是，每次比较两个正序数组中的第k/2大的数，截断其中较小中位数之前的所有数字，这些数字必不包含结果
// 直到某一个数组长度变为0，或k==1.

func kthSmallElement[T constraints.Ordered](nums1, nums2 []T, k int) T {
	var min1, min2 T     // 存储要比较的两个最小值
	var step1, step2 int // 存储要截断的长度

	for {
		// 三种结束情况
		if len(nums1) == 0 {
			return nums2[k-1]
		}
		if len(nums2) == 0 {
			return nums1[k-1]
		}
		if k == 1 {
			return MinN(nums1[0], nums2[0])
		}

		half := k / 2

		// 检查k/2是否越界，并做处理
		if len(nums1)-1 < half-1 {
			min1 = nums1[len(nums1)-1]
			step1 = len(nums1)
		} else {
			min1 = nums1[half-1]
			step1 = half
		}

		// 检查k/2是否越界，并做处理
		if len(nums2)-1 < half-1 {
			min2 = nums2[len(nums2)-1]
			step2 = len(nums2)
		} else {
			min2 = nums2[half-1]
			step2 = half
		}

		if min1 <= min2 {
			nums1 = nums1[step1:]
			k -= step1
		} else {
			nums2 = nums2[step2:]
			k -= step2
		}
	}
}
