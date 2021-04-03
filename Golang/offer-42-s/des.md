# 连续子数组的最大和

## 1. 题目描述
输入一个整型数组，数组中的一个或连续多个整数组成一个子数组。求所有子数组的和的最大值。

要求时间复杂度为O(n)。

## 2. 示例
```
输入: nums = [-2,1,-3,4,-1,2,1,-5,4]
输出: 6
解释: 连续子数组 [4,-1,2,1] 的和最大，为 6。
```

## 3. 解题
本题要求时间复杂度为O(n)，则很明显要求使用动态规划法

分析动态转移方程
设dp[i]为以下标为i的元素结尾的最大子数组和
则有
```
dp[i] = max(dp[i-1], 0) + arr[i]
```
为何是`max(dp[i-1], 0)`，因为若dp[i-1] > 0， 则对于dp[i]的值做出正贡献，即必有
`dp[i] < arr[i]`; 反之，若dp[i-1] < 0，则dp[i-1] 对于dp[i]的值做负贡献，即必有
`dp[i] > arr[i]`

> 一开始曾想
> ```
> dp[i] = max{ dp[i-1] + arr[i], arr[i], sum(arr[j:i])) }  (0<=j<i)
> ```  
> 很明显这种想法不对，因为dp[i-1]已经是arr[i]之前的最大子数组和，所以不存在其他子数组之和大于dp[i-1]
> 即`sum(arr[j:i]) = sum(arr[j:i-1]) + arr[i]`中的`sum(arr[j:i-1])`不可能大于dp[i-1]
> 因此，上述状态转移方程的第三项永远不可能是最大值。  
> 则正确的状态转移方程为：
> ```
> dp[i] = max{ dp[i-1] + arr[i], arr[i] } = max(dp[i-1], 0) + arr[i]
> ```