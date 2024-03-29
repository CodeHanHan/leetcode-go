# 面试题 17.08.马戏团人塔

## 1. 题目描述

有个马戏团正在设计叠罗汉的表演节目，一个人要站在另一人的肩膀上。出于实际和美观的考虑，在上面的人要比下面的人矮一点且轻一点。已知马戏团每个人的身高和体重，请编写代码计算叠罗汉最多能叠几个人。

**示例：**

```
输入：height = [65,70,56,75,60,68] weight = [100,150,90,190,95,110]
输出：6
解释：从上往下数，叠罗汉最多能叠 6 层：(56,90), (60,95), (65,100), (68,110), (70,150), (75,190)
```
 **提示：**
-  `height.length == weight.length <= 10000`

**标签**
`数组` `二分查找` `动态规划` `排序`


## 2. 解题

最长递增子序列问题

先按身高排序，然后根据体重求最长上升子序列。

求最长上升子序列：
1. 动态规划法
dp[i]: 以i为结尾的最长上升子序列
dp[i] = max{dp[w1], dp[w2], ...} + 1, 其中wx的体重小于wi的体重

2. 二分法
见[topic300](../../topic/topic300/readme.md)
