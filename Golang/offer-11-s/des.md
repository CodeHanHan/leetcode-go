# 旋转数组的最小数字
## 1. 题目描述
把一个数组最开始的若干个元素搬到数组的末尾，我们称之为数组的旋转。输入一个递增排序的数组的一个旋转，输出旋转数组的最小元素。例如，数组 [3,4,5,1,2] 为 [1,2,3,4,5] 的一个旋转，该数组的最小值为1。

## 2. 示例
```
输入：[3,4,5,1,2]
输出：1
```

```
输入：[2,2,2,0,1]
输出：0
```

## 3. 解题
最容易想到的方法是从头开始遍历，找到最小值即可，时间复杂度为O(n)  
除此之外，可利用其有序特性，使用二分查找法，
当某一个元素的前一个值比它大，则最小值即为当前元素