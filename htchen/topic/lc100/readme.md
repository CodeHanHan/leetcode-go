# 相同的树


## 1. 题目描述
给你两棵二叉树的根节点 p 和 q ，编写一个函数来检验这两棵树是否相同。  
如果两个树在结构上相同，并且节点具有相同的值，则认为它们是相同的。

## 2. 示例

提示：
```
两棵树上的节点数目都在范围 [0, 100] 内
-104 <= Node.val <= 104
```

示例 1：  
![1](https://raw.githubusercontent.com/HT-CHEN520/gopl/main/lc100-1.jpg)
```
输入：p = [1,2,3], q = [1,2,3]
输出：true
```
示例 2：  
![2](https://raw.githubusercontent.com/HT-CHEN520/image-lc/main/lc100-2.jpg)
```
输入：p = [1,2], q = [1,null,2]
输出：false
```

示例 3：  
![3](https://raw.githubusercontent.com/HT-CHEN520/image-lc/main/lc100-3.jpg)
```
输入：p = [1,2,1], q = [1,1,2]
输出：false
```

## 3. 解题
1.若p或q其中一个为空结点，那么两颗树不相同，返回false。  
2.递归遍历两个数的左右结点和结点的值是否相同，全部相同则两棵树相同，返回true；反之，返回false。