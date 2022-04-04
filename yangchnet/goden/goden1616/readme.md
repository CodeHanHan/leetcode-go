# 面试题 16.16.部分排序

## 1. 题目描述

给定一个整数数组，编写一个函数，找出索引 `m` 和 `n` ，只要将索引区间 `[m,n]` 的元素排好序，整个数组就是有序的。注意： `n-m` 尽量最小，也就是说，找出符合条件的最短序列。函数返回值为 `[m,n]` ，若不存在这样的 `m` 和 `n` （例如整个数组是有序的），请返回 `[-1,-1]` 。

**示例：**
```
输入： [1,2,4,7,10,11,7,12,6,7,16,18,19]
输出： [3,9]
```

**提示：**
-  `0 <= len(array) <= 1000000`

**标签**
`栈` `贪心` `数组` `双指针` `排序` `单调栈`

## 2. 解题

所求区间[l, r]

对于l，其右侧必存在小于array[l]的值，且array[l]为'小于右侧最小值的'最大值

对于r，其左侧必存在大于array[r]的值，且array[r]为'大于左侧最大值的'最小值