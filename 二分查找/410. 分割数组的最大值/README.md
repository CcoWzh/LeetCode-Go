#### [410. 分割数组的最大值](https://leetcode-cn.com/problems/split-array-largest-sum/)

给定一个非负整数数组和一个整数 *m*，你需要将这个数组分成 *m* 个非空的连续子数组。设计一个算法使得这 *m* 个子数组各自和的最大值最小。

**注意:**
数组长度 *n* 满足以下条件:

- 1 ≤ *n* ≤ 1000
- 1 ≤ *m* ≤ min(50, *n*)

**示例:**

```
输入:
nums = [7,2,5,10,8]
m = 2

输出:
18

解释:
一共有四种方法将nums分割为2个子数组。
其中最好的方式是将其分为[7,2,5] 和 [10,8]，
因为此时这两个子数组各自的和的最大值为18，在所有情况中最小。
```

---

第一次遇到这样的题目，即，分割数组，找最大值的最小。

动态规划和二分法都可以使用，但是二分法更加简单易懂。

[参考](https://leetcode-cn.com/problems/split-array-largest-sum/solution/er-fen-cha-zhao-by-coder233-2/)讲解得挺明白的：

```
解题思路：
由题意可知：子数组的最大值是有范围的，即在区间 [max(nums),sum(nums)]之中。
令 l=max(nums)，h=sum(nums)，mid=(l+h)/2，计算数组和最大值不大于mid对应的子数组个数 cnt(这个是关键！)
如果 cnt>m，说明划分的子数组多了，即我们找到的 mid 偏小，故 l=mid+1；
否则，说明划分的子数组少了，即 mid 偏大(或者正好就是目标值)，故 h=mid。
```

