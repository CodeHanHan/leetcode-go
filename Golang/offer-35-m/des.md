# 复杂链表的复制

## 1. 问题描述
请实现 copyRandomList 函数，复制一个复杂链表。在复杂链表中，每个节点除了有一个 next 指针指向下一个节点，还有一个 random 指针指向链表中的任意节点或者 null。

**示例1**
![复杂链表](https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2020/01/09/e1.png)
```
输入：head = [[7,null],[13,0],[11,4],[10,2],[1,0]]
输出：[[7,null],[13,0],[11,4],[10,2],[1,0]]
```
**示例2**
![复杂链表2](https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2020/01/09/e2.png)
```
输入：head = [[1,1],[2,1]]
输出：[[1,1],[2,1]]
```

## 2. 提示
提示：

-10000 <= Node.val <= 10000
Node.random 为空（null）或指向链表中的节点。
节点数目不超过 1000 。

## 3. 解题思路
本题难点在于random指针的复制，要想复制完整个链表，就需要复制全部的random指针
但有些random指针在复制完整个链表前不能复制。

使用切片存储所有新节点，使用字典存储节点对应的顺序
通过对字典中值的查取重新链接Random