# 59.螺旋矩阵 II

## 1. 题目描述

给你一个正整数  `n` ，生成一个包含 `1` 到  `n^2`  所有元素，且元素按顺时针顺序螺旋排列的  `n x n` 正方形矩阵 `matrix` 。

 

 **示例 1：** 
<img alt="" src="https://assets.leetcode.com/uploads/2020/11/13/spiraln.jpg" style="width: 242px; height: 242px;" />
```

输入：n = 3
输出：[[1,2,3],[8,9,4],[7,6,5]]

```
 **示例 2：** 

```

输入：n = 1
输出：[[1]]

```
 

 **提示：** 
-  `1 <= n <= 20` 
 
**标签**
`数组` `矩阵` `模拟` 


## 2. 解题
初始位置设为矩阵的左上角，初始方向设为向右。若下一步的位置超出矩阵边界，或者是之前访问过的位置，则顺时针旋转，进入下一个方向。如此反复直至填入n*n个元素。


