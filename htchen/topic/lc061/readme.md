# 61.旋转链表

## 1. 题目描述

给你一个链表的头节点 `head` ，旋转链表，将链表每个节点向右移动  `k` * * 个位置。

 

 **示例 1：** 
<img alt="" src="https://assets.leetcode.com/uploads/2020/11/13/rotate1.jpg" style="width: 600px; height: 254px;" />
```

输入：head = [1,2,3,4,5], k = 2
输出：[4,5,1,2,3]

```
 **示例 2：** 
<img alt="" src="https://assets.leetcode.com/uploads/2020/11/13/roate2.jpg" style="width: 472px; height: 542px;" />
```

输入：head = [0,1,2], k = 4
输出：[2,0,1]

```
 

 **提示：** 
- 链表中节点的数目在范围 `[0, 500]` 内
-  `-100 <= Node.val <= 100` 
-  `0 <= k <= 2 * 10^9` 
 
**标签**
`链表` `双指针` 


## 2. 解题
若链表的长度为len，则链表移动的次数为(k%len),要达到反转的效果，可以先把原链表的尾部和头部相连，再从新链表的头结点前断开即可。
