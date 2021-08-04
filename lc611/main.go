package main

import (
	"fmt"
	"sort"
)

/*
思路与算法

我们可以对方法一进行优化。

我们将当 a = \textit{nums}[i], b = \textit{nums}[j]a=nums[i],b=nums[j] 时，最大的满足 \textit{nums}[k] < \textit{nums}[i] + \textit{nums}[j]nums[k]<nums[i]+nums[j] 的下标 kk 记为 k_{i, j}k
i,j
​
 。可以发现，如果我们固定 ii，那么随着 jj 的递增，不等式右侧 \textit{nums}[i] + \textit{nums}[j]nums[i]+nums[j] 也是递增的，因此 k_{i, j}k
i,j
​
  也是递增的。

这样一来，我们就可以将 jj 和 kk 看成两个同向（递增）移动的指针，将方法一进行如下的优化：

我们使用一重循环枚举 ii。当 ii 固定时，我们使用双指针同时维护 jj 和 kk，它们的初始值均为 ii；

我们每一次将 jj 向右移动一个位置，即 j \leftarrow j+1j←j+1，并尝试不断向右移动 kk，使得 kk 是最大的满足 \textit{nums}[k] < \textit{nums}[i] + \textit{nums}[j]nums[k]<nums[i]+nums[j] 的下标。我们将 \max(k - j, 0)max(k−j,0) 累加入答案。

当枚举完成后，我们返回累加的答案即可。

细节

与方法一中「二分查找的失败」类似，方法二的双指针中，也会出现不存在满足 \textit{nums}[k] < \textit{nums}[i] + \textit{nums}[j]nums[k]<nums[i]+nums[j] 的下标的情况。此时，指针 kk 不会出现在指针 jj 的右侧，即 k - j \leq 0k−j≤0，因此我们需要将 k - jk−j 与 00 中的较大值累加入答案，防止错误的负数出现。

*/

func triangleNumber(nums []int) int {

	if len(nums) < 3 {
		return 0
	}
	sort.Ints(nums)
	ans := 0

	for i := 0; i < len(nums)-1; i++ {
		k := i
		for j := i + 1; j < len(nums); j++ {
			for k+1 < len(nums) && nums[k+1] < nums[i]+nums[j] {
				k++
			}
			if k-j > 0 {
				ans += k - j
			}
		}
	}
	return ans
}

func main() {
	res := triangleNumber([]int{2, 2, 3, 4})
	fmt.Println("res", res)
}
