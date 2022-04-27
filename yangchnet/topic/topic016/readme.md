# 最接近的三数之和

## 1. 题目描述
给定一个包括 n 个整数的数组 nums 和 一个目标值 target。找出 nums 中的三个整数，使得它们的和与 target 最接近。返回这三个数的和。假定每组输入只存在唯一答案。

## 2. 示例
```
输入：nums = [-1,2,1,-4], target = 1
输出：2
解释：与 target 最接近的和是 2 (-1 + 2 + 1 = 2) 。
```

**提示**
- $3 <= nums.length <= 10^3$
- $-10^3 <= nums[i] <= 10^3$
- $-10^4 <= target <= 10^4$

## 3. 解题
使用类似topic-15的方法，排序加双指针

- 特判：如果数组长度为3，直接返回数组全部元素和
- 对数组进行排序
- 设最接近的和closest初始为最大的32位正整数
- 遍历排序后数组（i=[0,len(nums)-2]）：
  - 令l = i+1, r = len(nums)-1, 设sum = nums[i] + nums[l] + nums[r], 当l < r时：
    - 若closest > abs(sum - target), 则closest = sum
    - 若sum-target等于0， 那么直接返回sum
    - 若sum-target小于0，则l++
    - 若sum-target大于0，则r--
- 返回closest
