package LC001

func twoSum(nums []int, target int) []int {
	for index, item := range nums {
		for i, a := range nums {
			if i != index && (item+a) == target {
				return []int{index, i}
			}
		}
	}
	return nil
}
