# 54.螺旋矩阵

## 1. 题目描述

给你一个 `m` 行 `n` 列的矩阵  `matrix` ，请按照 **顺时针螺旋顺序** ，返回矩阵中的所有元素。

 

 **示例 1：** 
<img alt="" src="https://assets.leetcode.com/uploads/2020/11/13/spiral1.jpg" style="width: 242px; height: 242px;" />
```

输入：matrix = [[1,2,3],[4,5,6],[7,8,9]]
输出：[1,2,3,6,9,8,7,4,5]

```
 **示例 2：** 
<img alt="" src="https://assets.leetcode.com/uploads/2020/11/13/spiral.jpg" style="width: 322px; height: 242px;" />
```

输入：matrix = [[1,2,3,4],[5,6,7,8],[9,10,11,12]]
输出：[1,2,3,4,8,12,11,10,9,5,6,7]

```
 

 **提示：** 
-  `m == matrix.length` 
-  `n == matrix[i].length` 
-  `1 <= m, n <= 10` 
-  `-100 <= matrix[i][j] <= 100` 
 
**标签**
`数组` `矩阵` `模拟` 


## 2. 解题
模拟螺旋矩阵的路径。初始位置是矩阵的左上角，初始方向是向右，当路径超出界限或者进入之前访问过的位置时，顺时针旋转，进入下一个方向。直至遍历完。
