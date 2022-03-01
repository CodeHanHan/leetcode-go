# 面试题 01.08.零矩阵

## 1. 题目描述

编写一种算法，若M × N矩阵中某个元素为0，则将其所在的行与列清零。

 

 **示例 1：** 

```
输入：
[
  [1,1,1],
  [1,0,1],
  [1,1,1]
]
输出：
[
  [1,0,1],
  [0,0,0],
  [1,0,1]
]

```
 **示例 2：** 

```
输入：
[
  [0,1,2,0],
  [3,4,5,2],
  [1,3,1,5]
]
输出：
[
  [0,0,0,0],
  [0,4,5,0],
  [0,3,1,0]
]

```
 
**标签**
`数组` `哈希表` `矩阵` 


## 2. 解题
1. 先记录为0的位置
2. 将记录位置的行列设为0
