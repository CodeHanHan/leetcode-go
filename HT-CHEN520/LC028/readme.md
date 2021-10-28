# 实现strStr()

## 1.题目描述

实现 strStr() 函数。  

给你两个字符串 haystack 和 needle ，请你在 haystack 字符串中找出 needle 字符串出现的第一个位置（下标从 0 开始）。如果不存在，则返回  -1 。  
对于本题而言，当 needle 是空字符串时我们应当返回 0 。这与 C 语言的 strstr() 以及 Java 的 indexOf() 定义相符。

## 2. 示例
提示：
```
0 <= haystack.length, needle.length <= 5 * 104
haystack 和 needle 仅由小写英文字符组成
```

示例 1：
```
输入：haystack = "hello", needle = "ll"
输出：2
```
示例 2：
```
输入：haystack = "aaaaa", needle = "bba"
输出：-1
```
示例 3：
```
输入：haystack = "", needle = ""
输出：0
```

## 3. 解题
1.若字符串haystack长度小于needle长度，匹配失败，返回-1;    
2.当 needle是空字符串时应当返回 0 ;  
3.golang切片求解，从首位开始循环比较是否匹配，若匹配，返回首个匹配项下标，否则返回-1。  