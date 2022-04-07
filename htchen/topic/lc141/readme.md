# 141.环形链表

## 1. 题目描述

给你一个链表的头节点 `head` ，判断链表中是否有环。

如果链表中有某个节点，可以通过连续跟踪 `next` 指针再次到达，则链表中存在环。 为了表示给定链表中的环，评测系统内部使用整数 `pos` 来表示链表尾连接到链表中的位置（索引从 0 开始）。 **注意： `pos` 不作为参数进行传递** 。仅仅是为了标识链表的实际情况。

 *如果链表中存在环* ，则返回 `true` 。 否则，返回 `false` 。

 

 **示例 1：** 

<img alt="" src="https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2018/12/07/circularlinkedlist.png" />

```

输入：head = [3,2,0,-4], pos = 1
输出：true
解释：链表中有一个环，其尾部连接到第二个节点。

```
 **示例 2：** 

<img alt="" src="https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2018/12/07/circularlinkedlist_test2.png" />

```

输入：head = [1,2], pos = 0
输出：true
解释：链表中有一个环，其尾部连接到第一个节点。

```
 **示例 3：** 

<img alt="" src="https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2018/12/07/circularlinkedlist_test3.png" />

```

输入：head = [1], pos = -1
输出：false
解释：链表中没有环。

```
 

 **提示：** 
- 链表中节点的数目范围是 `[0, 10^4]` 
-  `-10^5 <= Node.val <= 10^5` 
-  `pos` 为 `-1` 或者链表中的一个 **有效索引** 。
 

 **进阶：** 你能用 `O(1)` （即，常量）内存解决此问题吗？

 
**标签**
`哈希表` `链表` `双指针` 


## 2. 解题
采用快慢指针，快指针每次走两步，慢指针每次走一步，若快指针等于慢指针，则说明链表存在环，反之，不存在。
