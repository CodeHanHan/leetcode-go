# 题目描述
在一个 n * m 的二维数组中，每一行都按照从左到右递增的顺序排序，每一列都按照从上到下递增的顺序排序。请完成一个函数，输入这样的一个二维数组和一个整数，判断数组中是否含有该整数。

 

示例:

现有矩阵 matrix 如下：
```
[
  [1,   4,  7, 11, 15],  
  [2,   5,  8, 12, 19],  
  [3,   6,  9, 16, 22],  
  [10, 13, 14, 17, 24],  
  [18, 21, 23, 26, 30]  
] 
```
 
给定 target = 5，返回 true。

给定 target = 20，返回 false。

 

限制：

0 <= n <= 1000

0 <= m <= 1000

# 解题思路
1. 暴力法
2. 对角线扫描
    首先扫描对角线，如果存在，则返回
    如果不存在，则定位到大于给定数的元素，再从此元素，往上和往左扫描。
    > 这种方法在矩阵为n*n时可以，但当矩阵为n*m时，则需要进行额外的判断，因此不太好。
3. 标志法  
    从左下角开始寻找，如果当前值大于目标值，则向上移动，如果当前值小于目标值，则向右移动。
4. 二叉排序树
    从矩阵的右上角看，可看做是一棵二叉排序树
    
    
# bug
在写代码的时候，定义了一个
```java
    int m = matrix[0].length;
```
整个代码如下：
```java
public boolean findNumberIn2DArray(int[][] matrix, int target) {
        int m = matrix[0].length;
        int i = matrix.length-1, j = 0;  //定位到左下角, i是行，j是列
        while(i >= 0 && j < matrix[0].length){
            if(matrix[i][j] > target){  //大于目标值，往上
                i--;
            }else if(matrix[i][j] < target){  //小于目标值，往右
                j++;
            }else{
                return true;
            }
        }
        return false;
    }
```

但是在提交的时候出现了错误，Line 5: java.lang.ArrayIndexOutOfBoundsException: Index 0 out of bounds for length 0
原因在于，当输入为空数组的时候，索引0是越界的

改成如下代码：
```java
public boolean findNumberIn2DArray(int[][] matrix, int target) {
        int i = matrix.length-1, j = 0;  //定位到左下角, i是行，j是列
        while(i >= 0 && j < matrix[0].length){
            if(matrix[i][j] > target){  //大于目标值，往上
                i--;
            }else if(matrix[i][j] < target){  //小于目标值，往右
                j++;
            }else{
                return true;
            }
        }
        return false;
    }
```
问题来了，这里也有访问matrix[0]，为什么不会出错呢
原因在于，在while的判断结构里，首先判断i>=0,这时i=-1,就会判断为false，表达式就被短路了，所以后面的一句实际上没有被执行。因此不会产生错误。
