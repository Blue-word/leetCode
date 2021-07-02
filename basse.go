package main

import (
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
	//popStack := stack[len(stack)-1]
	stack = stack[:len(stack)-1]
	// 检查栈
	//len(stack) == 0

	/*
	 * 队列
	 */
	queue := make([]int, 10)
	// 入队
	queue = append(queue, 1)
	// 出队
	//popList := queue[0]
	queue = queue[1:]
	// 检查队列
	//len(queue) == 0

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
	//math.MaxInt32 // 1^31-1
	//math.MinInt32 // -1^31
	// int64 最大最小值（int默认是int64）
	//math.MaxInt64 // 1^31-1
	//math.MinInt64 // -1^31

	/*
	 * copy
	 */
	// 删除a[i]，可以用 copy 将i+1到末尾的值覆盖到i,然后末尾-1
	//s = []int{1, 2, 3}
	//copy(s[i:],s[i+1:])
	//s = s[:len(s)-1]
	//// make创建长度，则通过索引赋值
	//a:=make([]int,n)
	//a[n]=x
	//// make长度为0，则通过append()赋值
	//a:=make([]int,0)
	//a=append(a,x)

	/*
	 * 类型转换
	 */
	//s1 := "123456"
	//num1 := int(s[1]-'0')
	//str1 := string(s[0]) // "1"
	//b1 := byte(num1+'0') // '1'
	//
	//// 字符串转数字
	//num1,_=strconv.Atoi()
	//str1=strconv.Itoa()
}
