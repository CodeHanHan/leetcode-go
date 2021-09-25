# 寻找两个正序数组的中位数

## 1. 题目描述
给定两个大小分别为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。请你找出并返回这两个正序数组的中位数 。

## 2. 示例

示例 1：
```
输入：nums1 = [1,3], nums2 = [2]
输出：2.00000
解释：合并数组 = [1,2,3] ，中位数 2
```

示例 2：
```
输入：nums1 = [1,2], nums2 = [3,4]
输出：2.50000
解释：合并数组 = [1,2,3,4] ，中位数 (2 + 3) / 2 = 2.5
```

示例 3：
```
输入：nums1 = [0,0], nums2 = [0,0]
输出：0.00000
```

示例 4：
```
输入：nums1 = [], nums2 = [1]
输出：1.00000
```

示例 5：
```
输入：nums1 = [2], nums2 = []
输出：2.00000
```

提示：
- nums1.length == m
- nums2.length == n
- 0 <= m <= 1000
- 0 <= n <= 1000
- 1 <= m + n <= 2000
- -106 <= nums1[i], nums2[i] <= 106

## 3. 解题方法
最容易想到的方法是先合并两个数组，然后从中取中位数，时间复杂度为O(max(m, n))，空间复杂度为O(m+n)

更好的方法是双指针法，对于中位数，我们完全可以有数组长度确定它的位置，维护两个指针，分别指向两个数组的头部，每次将
指向数值较小的指针向后移动一位，如果某一个数组到头，则只需移动另一个数组上的指针，直到到达中位数的位置。