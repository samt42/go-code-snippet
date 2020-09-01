package leetcode

func checkPossibility(nums []int) bool {
	numsSize := len(nums)
	// 数组个数小于2个 一定可以
	if numsSize < 3 {
		return true
	}

	// 思路如下：
	// 如果出现 a[i] > a[i+1]   改变一个数 就面临两种选择
	// 1. 把a[i]变大
	// 2. 把a[i+1] 变小
	// 这两种选择其实是有依据的 需要比较a[i-1] 与 a[i+1]的值
	// eg.  ... 1 4 3 ...   只能选择把4变小   ... 3 4 1 ... 只能选择把1变大
	// 改变完之后，记录改变次数，再检测是否升序
	// 如果次数大于1，至少改了两次 返回false

	// 先让前两个有序
	// 因为没有左边没有数 所以对于前两个数来说，最佳选择就是吧 a[0] 变小
	changeCount := 0
	if nums[0] > nums[1] {
		nums[0] = nums[1]
		changeCount++
	}

	for i := 1; i < numsSize - 1; i++ {
		right := nums[i+1]
		if nums[i] > right {
			changeCount++
			if changeCount > 1 {
				// 后面不用再看了
				return false
			}
			left := nums[i-1]
			if left > right {
				nums[i+1] = nums[i]
			} else {
				nums[i] = left
			}
		}
	}
	return true
}
