#### [560. 和为K的子数组](https://leetcode-cn.com/problems/subarray-sum-equals-k/)

给定一个整数数组和一个整数 **k，**你需要找到该数组中和为 **k** 的连续的子数组的个数。

**示例 1 :**

```
输入:nums = [1,1,1], k = 2
输出: 2 , [1,1] 与 [1,1] 为两种不同的情况。
```

**说明 :**

1. 数组的长度为 [1, 20,000]。
2. 数组中元素的范围是 [-1000, 1000] ，且整数 **k** 的范围是 [-1e7, 1e7]。

----

前缀和：

使用前缀和可以很方便的计算`[i:j]`之间的和，不需要重复计算

```go
func subarraySum(nums []int, k int) int {
	//前缀和
	n := len(nums)
	preSum := make([]int, n+1)
	for i := 1; i <= n; i++ {
		preSum[i] = preSum[i-1] + nums[i-1]
	}
	//前缀求和，preSum[j]-preSum[i] == k
	count := 0
	for i := 0; i < n; i++ {
		for j := i + 1; j <= n; j++ {
			//fmt.Println(nums[i:j], preSum[j], preSum[i])
			if preSum[j]-preSum[i] == k {
				count++
			}
		}
	}
	return count
}
```

但是，前缀和可以优化，加一个哈希表即可：

```go
public class Solution {
    public int subarraySum(int[] nums, int k) {
        int count = 0, pre = 0;
        HashMap < Integer, Integer > mp = new HashMap < > ();
        mp.put(0, 1);
        for (int i = 0; i < nums.length; i++) {
            pre += nums[i];
            if (mp.containsKey(pre - k))
                count += mp.get(pre - k);
            mp.put(pre, mp.getOrDefault(pre, 0) + 1);
        }
        return count;
    }
}
```

