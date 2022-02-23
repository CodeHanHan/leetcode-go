# 螺旋矩阵

## 1. 题目描述
给你一个 m 行 n 列的矩阵 matrix ，请按照 顺时针螺旋顺序 ，返回矩阵中的所有元素。

## 2. 示例
示例1
![1](https://assets.leetcode.com/uploads/2020/11/13/spiral1.jpg)
```
输入：matrix = [[1,2,3],[4,5,6],[7,8,9]]
输出：[1,2,3,6,9,8,7,4,5]
```

示例2
![2](https://assets.leetcode.com/uploads/2020/11/13/spiral.jpg)
```
输入：matrix = [[1,2,3,4],[5,6,7,8],[9,10,11,12]]
输出：[1,2,3,4,8,12,11,10,9,5,6,7]
```

提示：
- m == matrix.length
- n == matrix[i].length
- 1 <= m, n <= 10
- -100 <= matrix[i][j] <= 100

## 3. 解题
螺旋矩阵每次碰到边界或访问过的元素时，方向的更改具有一定规律，可利用这个规律使其在碰到边界或访问过元素时自动改变方向，如果改变方向后仍然碰到访问过的元素，则说明遍历完毕。
