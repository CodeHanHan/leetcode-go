package goden1003

// 暴力搜索
func search(arr []int, target int) int {
	for i := 0; i < len(arr); i++ {
		if arr[i] == target {
			return i
		}
	}

	return -1
}

// 二分查找
func search2(arr []int, target int) int {
	l, r := 0, len(arr)-1

	var mid int
	for l <= r {
		mid = l + (r-l)/2
		if arr[l] == target {
			return l
		} else if arr[l] == arr[mid] {
			l = l + 1
		} else if arr[l] < arr[mid] {
			if target < arr[l] || arr[mid] < target {
				l = mid
			} else {
				l++
				r = mid
			}
		} else if arr[l] > arr[mid] {
			if target < arr[l] && arr[mid] < target {
				l = mid
			} else {
				l++
				r = mid
			}
		}
	}

	return -1
}
