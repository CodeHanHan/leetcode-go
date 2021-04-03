# 问题描述
请实现一个函数，用来判断一棵二叉树是不是对称的。如果一棵二叉树和它的镜像一样，那么它是对称的。

例如，二叉树 [1,2,2,3,4,4,3] 是对称的。
```
    1
   / \
  2   2
 / \ / \
3  4 4  3
```

但是下面这个 [1,2,2,null,3,null,3] 则不是镜像对称的:
```
    1
   / \
  2   2
   \   \
   3    3
```

示例 1：   
输入：root = [1,2,2,3,4,4,3]  
输出：true  

示例 2：  
输入：root = [1,2,2,null,3,null,3]  
输出：false  


限制：  
0 <= 节点个数 <= 1000

# 解题思路
同样使用递归  
当两个节点对称时：
如：
```
    1
   / \
(lchild)2(r)2
 / \ / \
3  4 4  3
```
* lchild.val == r.val
* lchild.left.val == r.right.val
* lchild.right.val = r.left.val

递归结束条件：  
当以上任意一条不满足时，返回false  
当l==r==nil, 返回true

下一次递归：    
symmetry（lchild.left, r.right) && symmetry（lchild.right, r.left) 
 

