# 面试题 17.13.恢复空格

## 1. 题目描述

哦，不！你不小心把一个长篇文章中的空格、标点都删掉了，并且大写也弄成了小写。像句子 `"I reset the computer. It still didn&rsquo;t boot!"` 已经变成了 `"iresetthecomputeritstilldidntboot"` 。在处理标点符号和大小写之前，你得先把它断成词语。当然了，你有一本厚厚的词典 `dictionary` ，不过，有些词没在词典里。假设文章用 `sentence` 表示，设计一个算法，把文章断开，要求未识别的字符最少，返回未识别的字符数。

**注意：** 本题相对原题稍作改动，只需返回未识别的字符数

**示例：**

```
输入：
dictionary = ["looked","just","like","her","brother"]
sentence = "jesslookedjustliketimherbrother"
输出： 7
解释： 断句后为"jess looked just like tim her brother"，共7个未识别字符。
```

**提示：**
- `0 <= len(sentence) <= 1000`
- `dictionary` 中总字符数不超过 150000。
- 你可以认为 `dictionary` 和 `sentence` 中只包含小写字母。

**标签**
`字典树` `数组` `哈希表` `字符串` `动态规划` `哈希函数` `滚动哈希`


## 2. 解题

Trie🌲+动态规划
dp[i]: 以i为结尾的字符处，最少未被确定的字符.

考虑转移方程，每次转移的时候我们考虑第 j(j≤i) 个到第 i 个字符组成的子串sentence[j:i] （注意字符串下标从 0 开始）是否能在词典中找到，如果能找到的话按照定义转移方程即为

dp[i]=min(dp[i],dp[j−1])

否则没有找到的话我们可以复用 dp[i−1] 的状态再加上当前未被识别的第 i 个字符，因此此时dp 值为

dp[i]=dp[i−1]+1

最后问题化简成了转移的时候如何快速判断当前子串是否存在于词典中
