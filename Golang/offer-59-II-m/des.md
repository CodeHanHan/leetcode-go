# 队列的最大值II

## 1. 题目描述
请定义一个队列并实现函数 `max_value` 得到队列里的最大值，要求函数`max_value`、`push_back` 和 `pop_front` 的均摊时间复杂度都是O(1)。

若队列为空，`pop_front` 和 `max_value`需要返回 -1

## 示例
```
输入: 
["MaxQueue","push_back","push_back","max_value","pop_front","max_value"]
[[],[1],[2],[],[],[]]
输出:[null,null,null,2,1,2]
```

```
输入: 
["MaxQueue","pop_front","max_value"]
[[],[],[]]
输出: [null,-1,-1]
```

**限制**
- 1 <= push_back,pop_front,max_value的总操作数 <= 10000
- 1 <= value <= 10^5

## 3. 解题
使用一个辅助队列，用于返回队列中的最大值。
使辅助队列保持（非）单调递减，当要返回最大值时，可直接返回第一个元素，当要添加元素时，删除队列中所有比要添加元素小的值
