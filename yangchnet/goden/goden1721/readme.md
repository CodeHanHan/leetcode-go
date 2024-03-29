# 面试题 17.21.直方图的水量

## 1. 题目描述

给定一个直方图(也称柱状图)，假设有人从上面源源不断地倒水，最后直方图能存多少水量?直方图的宽度为 1。

<img src="https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2018/10/22/rainwatertrap.png" style="height: 161px; width: 412px;">

<small>上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的直方图，在这种情况下，可以接 6 个单位的水（蓝色部分表示水）。 **感谢 Marcos** 贡献此图。</small>

 **示例:**

```
输入: [0,1,0,2,1,0,1,3,2,1,2,1]
输出: 6
```

**标签**
`栈` `数组` `双指针` `动态规划` `单调栈`


## 2. 解题

对于下标 i，水能到达的最大高度等于下标 i 两边的最大高度的最小值，下标 i 处能接的水的量等于下标 i 处的水能到达的最大高度减去 height[i]。

可先倒序遍历，记录每个位置右边的最大值。再顺序遍历，记录每个位置左边的最大值。最后求出所盛水量。
