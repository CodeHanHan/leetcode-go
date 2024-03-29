# 接雨水

## 1. 题目描述
给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。

## 2. 示例
示例1
![1](https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2018/10/22/rainwatertrap.png)
```
输入：height = [0,1,0,2,1,0,1,3,2,1,2,1]
输出：6
解释：上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。
```

示例2
```
输入：height = [4,2,0,3,2,5]
输出：9
```

**提示**
- n == height.length
- $1 <= n <= 2 * 10^4$
- $0 <= height[i] <= 10^5$

## 3. 解题

1. 某一个位置上可盛放的雨水量由其左右两边的最大值中的最小值决定。例如某处位置高为1，其左侧最大值为3， 右侧最大值为2，则此处可盛放雨水量为min(3, 2)-1 = 1

则可先求出每个位置上对应的左最大值与右最大值，然后通过min(maxLeft[i], maxRight[i]) - height[i] 得出此处盛放雨水量。
