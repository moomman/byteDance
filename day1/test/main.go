package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	window := maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3)

	for i, _ := range window {
		fmt.Printf("%d ", window[i])
	}
}

func strToInt(str string) int {
	str = strings.Trim(str, " ")
	if len(str) == 0 {
		return 0
	}
	bustr := math.MaxInt32 / 10
	ind, sign := 1, 1
	//为了判断该数字是正数还是负数,结果用sign表示
	if str[0] == '-' {
		sign = -1
	} else if str[0] != '+' {
		ind = 0
	}
	stLen := len(str)
	sum := 0
	for i := ind; i < stLen; i++ {
		if str[i] >= '0' && str[i] <= '9' {

			if sum > bustr || (sum == bustr && str[i] > '7') {
				if sign == 1 {
					return math.MaxInt32
				} else {
					return math.MinInt32
				}
			}
			sum *= 10
			sum += int(str[i] - '0')
		} else {
			break
		}

	}

	return sign * sum
}

func maxSlidingWindow(nums []int, k int) []int {
	q := []int{}
	push := func(i int) {
		for len(q) > 0 && nums[i] >= nums[q[len(q)-1]] {
			q = q[:len(q)-1]
		}
		q = append(q, i)
	}

	for i := 0; i < k; i++ {
		push(i)
	}

	n := len(nums)
	ans := make([]int, 1, n-k+1)
	ans[0] = nums[q[0]]
	for i := k; i < n; i++ {
		push(i)
		for q[0] <= i-k {
			q = q[1:]
		}
		ans = append(ans, nums[q[0]])
	}
	return ans
}

//func maxSlidingWindow(nums []int, k int) []int {
//	ans := []int{}
//	que := []int{}
//	for i := range nums {
//		if len(que) > 0 && que[0] <= i - k {
//			que = que[1:] //元素过期了超出了窗口范围就踢出去
//		}
//		//新元素进来时，判断是不是可以更新单调队列
//
//		for len(que) > 0 && nums[que[len(que)-1]] <= nums[i] {
//			que = que[:len(que)-1]
//		}
//		que = append(que, i)
//		//超出队列容量时就可以更新答案了
//		if i >= k - 1{
//			ans = append(ans, nums[que[0]])
//		}
//	}
//	return ans
//}
