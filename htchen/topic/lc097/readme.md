# 97.交错字符串

## 1. 题目描述

给定三个字符串 `s1` 、 `s2` 、 `s3` ，请你帮忙验证 `s3` 是否是由 `s1` 和 `s2`  **交错** 组成的。

两个字符串 `s` 和 `t` **交错** 的定义与过程如下，其中每个字符串都会被分割成若干 **非空** 子字符串：
-  `s = s<sub>1</sub> + s<sub>2</sub> + ... + s<sub>n</sub>` 
-  `t = t<sub>1</sub> + t<sub>2</sub> + ... + t<sub>m</sub>` 
-  `|n - m| <= 1` 
-  **交错** 是 `s<sub>1</sub> + t<sub>1</sub> + s<sub>2</sub> + t<sub>2</sub> + s<sub>3</sub> + t<sub>3</sub> + ...` 或者 `t<sub>1</sub> + s<sub>1</sub> + t<sub>2</sub> + s<sub>2</sub> + t<sub>3</sub> + s<sub>3</sub> + ...` 
 **注意：** `a + b` 意味着字符串 `a` 和 `b` 连接。

 

 **示例 1：** 
<img alt="" src="https://assets.leetcode.com/uploads/2020/09/02/interleave.jpg" />
```

输入：s1 = "aabcc", s2 = "dbbca", s3 = "aadbbcbcac"
输出：true

```
 **示例 2：** 

```

输入：s1 = "aabcc", s2 = "dbbca", s3 = "aadbbbaccc"
输出：false

```
 **示例 3：** 

```

输入：s1 = "", s2 = "", s3 = ""
输出：true

```
 

 **提示：** 
-  `0 <= s1.length, s2.length <= 100` 
-  `0 <= s3.length <= 200` 
-  `s1` 、 `s2` 、和 `s3` 都由小写英文字母组成
 

 **进阶：** 您能否仅使用 `O(s2.length)` 额外的内存空间来解决它?

 
**标签**
`字符串` `动态规划` 


## 2. 解题
动态规划
