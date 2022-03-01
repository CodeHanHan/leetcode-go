# 82.删除排序链表中的重复元素 II

## 1. 题目描述

给定一个已排序的链表的头 `head` ， *删除原始链表中所有重复数字的节点，只留下不同的数字* 。返回 *已排序的链表* 。

 

 **示例 1：** 
<img alt="" src="https://assets.leetcode.com/uploads/2021/01/04/linkedlist1.jpg" style="height: 142px; width: 500px;" />
```

输入：head = [1,2,3,3,4,4,5]
输出：[1,2,5]

```
 **示例 2：** 
<img alt="" src="https://assets.leetcode.com/uploads/2021/01/04/linkedlist2.jpg" style="height: 164px; width: 400px;" />
```

输入：head = [1,1,1,2,3]
输出：[2,3]

```
 

 **提示：** 
- 链表中节点数目在范围 `[0, 300]` 内
-  `-100 <= Node.val <= 100` 
- 题目数据保证链表已经按升序 **排列** 
 
**标签**
`链表` `双指针` 


## 2. 解题
由于给定的链表是排好序的，所以重复的元素在链表中出现的位置是连续的。为了便于删除头结点，先创建一个哑结点dummy指向头结点，设指针p初始指向头结点，若p.Val=p.Next.Val，则存在相同结点，需循环删除；若p.Val!=p.Next.Val，则说明该结点只有一个，跳过继续遍历。
