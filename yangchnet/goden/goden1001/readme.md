# 面试题 10.01.合并排序的数组

## 1. 题目描述

给定两个排序后的数组 A 和 B，其中 A 的末端有足够的缓冲空间容纳 B。 编写一个方法，将 B 合并入 A 并排序。

初始化 A 和 B 的元素数量分别为 *m* 和 *n* 。

**示例:**

```
输入:
A = [1,2,3,0,0,0], m = 3
B = [2,5,6],       n = 3

输出: [1,2,2,3,5,6]
```

**说明:**

- `A.length == n + m`

**标签**
`数组` `双指针` `排序`

## 2. 解题

双指针法

初始i，j分别指向a，b的最后一个元素, cur指向待插入位置，初始为a[-1]：

- 如果a[i] > b[j]，a[cur] = a[i], i--, cur--
- 如果a[i] < b[j], a[cur] = b[j], j--, cur--
- 重复上述两步，直到i<0 || j<0