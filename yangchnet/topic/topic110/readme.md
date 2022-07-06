# 110.平衡二叉树

## 1. 题目描述

给定一个二叉树，判断它是否是高度平衡的二叉树。

本题中，一棵高度平衡二叉树定义为：

<blockquote>
一个二叉树 *每个节点 * 的左右两个子树的高度差的绝对值不超过 1 。
</blockquote>

 

 **示例 1：** 
<img alt="" src="https://assets.leetcode.com/uploads/2020/10/06/balance_1.jpg" style="width: 342px; height: 221px;" />
```

输入：root = [3,9,20,null,null,15,7]
输出：true

```
 **示例 2：** 
<img alt="" src="https://assets.leetcode.com/uploads/2020/10/06/balance_2.jpg" style="width: 452px; height: 301px;" />
```

输入：root = [1,2,2,3,3,null,null,4,4]
输出：false

```
 **示例 3：** 

```

输入：root = []
输出：true

```
 

 **提示：** 
- 树中的节点数在范围 `[0, 5000]` 内
-  `-10^4 <= Node.val <= 10^4` 
 
**标签**
`树` `深度优先搜索` `二叉树` 


## 2. 解题

