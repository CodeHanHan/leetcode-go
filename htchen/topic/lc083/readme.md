# 83.删除排序链表中的重复元素

## 1. 题目描述

给定一个已排序的链表的头<meta charset="UTF-8" /> `head` ， *删除所有重复的元素，使每个元素只出现一次* 。返回 *已排序的链表* 。

 

 **示例 1：** 
<img alt="" src="https://assets.leetcode.com/uploads/2021/01/04/list1.jpg" style="height: 160px; width: 200px;" />
```

输入：head = [1,1,2]
输出：[1,2]

```
 **示例 2：** 
<img alt="" src="https://assets.leetcode.com/uploads/2021/01/04/list2.jpg" style="height: 123px; width: 300px;" />
```

输入：head = [1,1,2,3,3]
输出：[1,2,3]

```
 

 **提示：** 
- 链表中节点数目在范围 `[0, 300]` 内
-  `-100 <= Node.val <= 100` 
- 题目数据保证链表已经按升序 **排列** 
 
**标签**
`链表` 


## 2. 解题
由于给定的链表是排好序的，我们只需要对链表进行一次遍历，就可以删除重复的元素。设哑结点dummy指向链表的头节点，设p指针指向head，随后开始对链表进行遍历。  
若p.Val与 p.next.Val对应的元素相同，那么我们就将p.next节点从链表中移除；否则说明链表中已经不存在其它与p指针对应的元素相同的节点。

