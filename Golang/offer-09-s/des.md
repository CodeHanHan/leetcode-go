# 用两个栈实现队列

## 1. 题目描述
用两个栈实现一个队列。队列的声明如下，请实现它的两个函数 appendTail 和 deleteHead ，
分别完成在队列尾部插入整数和在队列头部删除整数的功能。(若队列中没有元素，deleteHead操作返回 -1 )

## 2. 示例
```
输入：
["CQueue","appendTail","deleteHead","deleteHead"]
[[],[3],[],[]]
输出：[null,null,3,-1]
```

```
输入：
["CQueue","deleteHead","appendTail","appendTail","deleteHead","deleteHead"]
[[],[],[5],[2],[],[]]
输出：[null,-1,null,null,5,2]
```

**提示**
1 <= values <= 10000
最多会对 appendTail、deleteHead 进行 10000 次调用

## 3. 解题
两个栈分别用于尾部附加和头部删除， 设为tail和head   
对于appendTail,只需将元素进tail栈即可   
当要deleteHead时，若head不空，则直接对head出栈，若head空，则将tail中的元素依次转移到head中，再出栈，若tail也为空，则返回-1
