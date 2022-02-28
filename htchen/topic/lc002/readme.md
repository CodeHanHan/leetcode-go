# 2.两数相加

## 1. 题目描述

给你两个  **非空** 的链表，表示两个非负的整数。它们每位数字都是按照  **逆序**  的方式存储的，并且每个节点只能存储  **一位**  数字。

请你将两个数相加，并以相同形式返回一个表示和的链表。

你可以假设除了数字 0 之外，这两个数都不会以 0 开头。

 

 **示例 1：**   
<img alt="" src="https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2021/01/02/addtwonumber1.jpg" style="width: 483px; height: 342px;" />
```

输入：l1 = [2,4,3], l2 = [5,6,4]
输出：[7,0,8]
解释：342 + 465 = 807.

```
 **示例 2：** 

```

输入：l1 = [0], l2 = [0]
输出：[0]

```
 **示例 3：** 

```

输入：l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
输出：[8,9,9,9,0,0,0,1]

```
 

 **提示：** 
- 每个链表中的节点数在范围 `[1, 100]` 内
-  `0 <= Node.val <= 9` 
- 题目数据保证列表表示的数字不含前导零
 
**标签**
`递归` `链表` `数学` 


## 2. 解题
由于每位数字都是按照逆序的方式存储的，因此两个链表对应的位置可以直接相加，设相应的值为n1和n2，进位为carry，故结果链表对应的每个结点的值为(n1+n2)%10,进位为(n1+n2)/10。若最后还存在进位，需在链表尾部创建一个新的结点存储进位carry。