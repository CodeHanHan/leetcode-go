# 包含min函数的栈

## 1. 题目描述
定义栈的数据结构，请在该类型中实现一个能够得到栈的最小元素的 min 函数在该栈中，
调用 min、push 及 pop 的时间复杂度都是 O(1)。

## 2. 示例
示例:  
MinStack minStack = new MinStack();  
minStack.push(-2);  
minStack.push(0);  
minStack.push(-3);   
minStack.min();   --> 返回 -3.  
minStack.pop();  
minStack.top();      --> 返回 0.  
minStack.min();   --> 返回 -2.
提示：  
各函数的调用总次数不超过 20000 次

## 3. 解题
使用辅助栈  
使用一个数组data存储所有元素，其中所有元素先进后出，使用另一个辅助数组min存储最小值。  
没次入栈，先将元素存储至data中，然后比较入栈元素和min数组的顶部（当前最小值）的大小。  
若入栈元素比最小值小，则将入栈元素放入min数组顶部。

出栈时，首先从data数组顶部取出出栈值，若出栈值等于min数组顶部值，则将min数组最后一个元素删除。

取最小值时，直接返回min数组顶部值。