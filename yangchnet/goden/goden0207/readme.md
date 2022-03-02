# 面试题 02.07.链表相交

## 1. 题目描述

给你两个单链表的头节点  `headA` 和 `headB` ，请你找出并返回两个单链表相交的起始节点。如果两个链表没有交点，返回 `null` 。

图示两个链表在节点 `c1` 开始相交 **：** 

<a href="https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2018/12/14/160_statement.png" target="_blank"><img alt="" src="https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2018/12/14/160_statement.png" style="height: 130px; width: 400px;" /></a>

题目数据 **保证** 整个链式结构中不存在环。

 **注意** ，函数返回结果后，链表必须 **保持其原始结构** 。

 

 **示例 1：** 

<a href="https://assets.leetcode.com/uploads/2018/12/13/160_example_1.png" target="_blank"><img alt="" src="https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2018/12/14/160_example_1.png" style="height: 130px; width: 400px;" /></a>

```

输入：intersectVal = 8, listA = [4,1,8,4,5], listB = [5,0,1,8,4,5], skipA = 2, skipB = 3
输出：Intersected at '8'
解释：相交节点的值为 8 （注意，如果两个链表相交则不能为 0）。
从各自的表头开始算起，链表 A 为 [4,1,8,4,5]，链表 B 为 [5,0,1,8,4,5]。
在 A 中，相交节点前有 2 个节点；在 B 中，相交节点前有 3 个节点。

```
 **示例 2：** 

<a href="https://assets.leetcode.com/uploads/2018/12/13/160_example_2.png" target="_blank"><img alt="" src="https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2018/12/14/160_example_2.png" style="height: 136px; width: 350px;" /></a>

```

输入：intersectVal = 2, listA = [0,9,1,2,4], listB = [3,2,4], skipA = 3, skipB = 1
输出：Intersected at '2'
解释：相交节点的值为 2 （注意，如果两个链表相交则不能为 0）。
从各自的表头开始算起，链表 A 为 [0,9,1,2,4]，链表 B 为 [3,2,4]。
在 A 中，相交节点前有 3 个节点；在 B 中，相交节点前有 1 个节点。

```
 **示例 3：** 

<a href="https://assets.leetcode.com/uploads/2018/12/13/160_example_3.png" target="_blank"><img alt="" src="https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2018/12/14/160_example_3.png" style="height: 126px; width: 200px;" /></a>

```

输入：intersectVal = 0, listA = [2,6,4], listB = [1,5], skipA = 3, skipB = 2
输出：null
解释：从各自的表头开始算起，链表 A 为 [2,6,4]，链表 B 为 [1,5]。
由于这两个链表不相交，所以 intersectVal 必须为 0，而 skipA 和 skipB 可以是任意值。
这两个链表不相交，因此返回 null 。

```
 

 **提示：** 
-  `listA` 中节点数目为 `m` 
-  `listB` 中节点数目为 `n` 
-  `0 <= m, n <= 3 * 10^4` 
-  `1 <= Node.val <= 10^5` 
-  `0 <= skipA <= m` 
-  `0 <= skipB <= n` 
- 如果 `listA` 和 `listB` 没有交点， `intersectVal` 为 `0` 
- 如果 `listA` 和 `listB` 有交点， `intersectVal == listA[skipA + 1] == listB[skipB + 1]` 
 

 **进阶：** 你能否设计一个时间复杂度 `O(n)` 、仅用 `O(1)` 内存的解决方案？

 
**标签**
`哈希表` `链表` `双指针` 


## 2. 解题
由于两个链表的前面部分可能不一样长，因此我们不能一次性的找到那个交点。我们可以在A链表到尾部后再连到B链表上，B链表到尾部后连到A链表上，这样的话如果前面的部分不一样长，那么第二次就一定同时到那个交点上

两个链表相交的长度为c，A链表前面的部分长度为a，B链表前面的长度为b。那么第一次到相交的起始节点时，A的指针要走a个长度，B的指针要走b个长度。这样两个指针走的长度可能会不一致。
继续走到尾部，再互相连接后，A的指针走的距离是（a+c+b）,B的指针走的距离是（b+c+a）,这样两个指针就会同时到相交的节点了。

如果没有交点，那么两个就会同时到null，就会直接返回null


