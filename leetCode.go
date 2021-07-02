package main

func main() {
	subsets([]int{1, 2, 3})
}

/*
 * 给你两个字符串 haystack 和 needle ，请你在 haystack 字符串中找出 needle 字符串出现的第一个位置（下标从 0 开始）。
 * 如果不存在，则返回  -1
 */
func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	var i, j int
	// i不需要到len-1
	for i = 0; i < len(haystack)-len(needle)+1; i++ {
		for j = 0; j < len(needle); j++ {
			if haystack[i+j] != needle[j] {
				break
			}
		}
		// 判断字符串长度是否相等
		if len(needle) == j {
			return i
		}
	}
	return -1
}

/*
 * 给你一个整数数组 nums ，数组中的元素 互不相同 。返回该数组所有可能的子集（幂集）。
 * 解集 不能 包含重复的子集。你可以按 任意顺序 返回解集
 * 输入：nums = [1,2,3]
 * 输出：[[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]
 */
func subsets(nums []int) [][]int {
	var i, j int
	for i = 0; i < len(nums); i++ {

	}
}
