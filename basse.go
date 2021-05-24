package leetCode

import (
	"math"
	"sort"
)

func baseStruct() {
	/*
	 * 栈
	 */
	stack := make([]int, 10)
	// 压栈
	stack = append(stack, 1)
	// 出栈
	popStack := stack[len(stack)-1]
	stack = stack[:len(stack)-1]
	// 检查栈
	len(stack) == 0

	/*
	 * 队列
	 */
	queue := make([]int, 10)
	// 入队
	queue = append(queue, 1)
	// 出队
	popList := queue[0]
	queue = queue[1:]
	// 检查队列
	len(queue) == 0

	/*
	 * 字典
	 */
	dict := make(map[string]int)
	// 设置kv
	dict["a"] = 1
	// 删除kv
	delete(dict, "a")
	// 遍历
	for k, v := range dict {
		println(k, v)
	}

	/*
	 * sort
	 */
	s := []int{1, 2, 3}
	sort.Ints([]int{})                  // int排序
	sort.Strings([]string{})            // 字符串排序
	sort.Slice(s, func(i, j int) bool { // 自定义排序
		return s[i] < s[j] // 降序
	})

	/*
	 * math
	 */
	// int32 最大最小值
	math.MaxInt32 // 1^31-1
	math.MinInt32 // -1^31
	// int64 最大最小值（int默认是int64）
	math.MaxInt64 // 1^31-1
	math.MinInt64 // -1^31

	/*
	 * copy
	 */

}