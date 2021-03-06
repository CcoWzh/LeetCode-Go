#### [53. 最大子序和](https://leetcode-cn.com/problems/maximum-subarray/)

难度简单2178收藏分享切换为英文关注反馈

给定一个整数数组 `nums` ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。

**示例:**

```
输入: [-2,1,-3,4,-1,2,1,-5,4],
输出: 6
解释: 连续子数组 [4,-1,2,1] 的和最大，为 6。
```

**进阶:**

如果你已经实现复杂度为 O(*n*) 的解法，尝试使用更为精妙的分治法求解。

#### 暴力解法

```go
func maxSubArray(nums []int) int {
    if len(nums) == 0 {
		return -2147483648
	}
	max, n := nums[0], len(nums)

	for i := 0; i < n; i++ {
		sum := 0
		for j := i; j < n; j++ {
			sum += nums[j]
			if max < sum {
				max = sum
			}
		}
	}
	return max
}
```

#### 动态规划

- dp[i] 表示以i结尾的，连续子数组的最大和
- dp[i] = `max(dp[i-1]+nums[i], nums[i])`

```go
func maxSubArray(nums []int) int {
	n := len(nums)
	if n == 0 {
		return -2147483648
	}
	dp := make([]int, n)
	//dp[i] 表示以i结尾的，连续子数组的最大和
	//dp[i] = max(dp[i-1]+nums[i], nums[i])
	dp[0] = nums[0]
	for i := 1; i < n; i++ {
		dp[i] = max(dp[i-1]+nums[i], nums[i])
	}
	// fmt.Println(dp)
	res := dp[0]
	for i := 0; i < n; i++ {
		if res < dp[i] {
			res = dp[i]
		}
	}
	return res
}

func max(a, b int) int {
	if a < b {
		a = b
	}
	return a
}
```

