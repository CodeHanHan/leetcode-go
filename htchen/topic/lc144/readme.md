# 144.二叉树的前序遍历

## 1. 题目描述

给你二叉树的根节点 `root` ，返回它节点值的  **前序** * * 遍历。

 

 **示例 1：** 
<img alt="" src="https://assets.leetcode.com/uploads/2020/09/15/inorder_1.jpg" style="width: 202px; height: 324px;" />
```

输入：root = [1,null,2,3]
输出：[1,2,3]

```
 **示例 2：** 

```

输入：root = []
输出：[]

```
 **示例 3：** 

```

输入：root = [1]
输出：[1]

```
 **示例 4：** 
<img alt="" src="https://assets.leetcode.com/uploads/2020/09/15/inorder_5.jpg" style="width: 202px; height: 202px;" />
```

输入：root = [1,2]
输出：[1,2]

```
 **示例 5：** 
<img alt="" src="https://assets.leetcode.com/uploads/2020/09/15/inorder_4.jpg" style="width: 202px; height: 202px;" />
```

输入：root = [1,null,2]
输出：[1,2]

```
 

 **提示：** 
- 树中节点数目在范围 `[0, 100]` 内
-  `-100 <= Node.val <= 100` 
 

 **进阶：** 递归算法很简单，你可以通过迭代算法完成吗？

 
**标签**
`栈` `树` `深度优先搜索` `二叉树` 


## 2. 解题
1.递归  
2.迭代
