# 面试题 16.18.模式匹配

## 1. 题目描述

你有两个字符串，即 `pattern` 和 `value` 。 `pattern` 字符串由字母 `"a"` 和 `"b"` 组成，用于描述字符串中的模式。例如，字符串 `"catcatgocatgo"` 匹配模式 `"aabab"` （其中 `"cat"` 是 `"a"` ， `"go"` 是 `"b"` ），该字符串也匹配像 `"a"` 、 `"ab"` 和 `"b"` 这样的模式。但需注意 `"a"` 和 `"b"` 不能同时表示相同的字符串。编写一个方法判断 `value` 字符串是否匹配 `pattern` 字符串。

 **示例 1：**

```
输入： pattern = "abba", value = "dogcatcatdog"
输出： true

```
 **示例 2：**

```
输入： pattern = "abba", value = "dogcatcatfish"
输出： false

```
 **示例 3：**

```
输入： pattern = "aaaa", value = "dogcatcatdog"
输出： false

```
 **示例 4：**

```
输入： pattern = "abba", value = "dogdogdogdog"
输出： true
解释： "a"="dogdog",b=""，反之也符合规则

```
 **提示：**
-  `1 <= len(pattern) <= 1000`
-  `0 <= len(value) <= 1000`
- 你可以假设 `pattern` 只包含字母 `"a"` 和 `"b"` ， `value` 仅包含小写字母。

**标签**
`数学` `字符串` `回溯` `枚举`


## 2. 解题
由于a,b匹配的内容与长度都不明确，因此需要枚举出a，b的长度进行处理

1. 考虑pattern为空
2. pattern不为空，记录patter中a，b的数量
   1. value为"", 判断pattern是否为一种字符组成
   2. value不为""
      1. 处理patter只有一种字符串的情况
      2. 处理pattern中'a','b'可为""的情况
      3. 枚举'a', 'b'所代表的字符串长度

