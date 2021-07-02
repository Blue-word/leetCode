package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
 * 排序算法
 */

/*
 * 快速排序
 * 给你一个整数数组 nums，请你将该数组升序排列
 */
func sortArray(nums []int) []int {
	if len(nums) == 0 {
		return nums
	}
	quickSort(nums, 0, len(nums)-1)
	return nums
}
func quickSort(nums []int, l, r int) {
	if l >= r {
		return
	}
	rand.Seed(time.Now().Unix())
	pivot := rand.Intn(r-l+1) + l
	nums[pivot], nums[r] = nums[r], nums[pivot]
	i := l - 1
	for j := l; j < r; j++ {
		if nums[j] <= nums[r] {
			i++
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	i = i + 1
	nums[i], nums[r] = nums[r], nums[i]
	quickSort(nums, l, i-1)
	quickSort(nums, i+1, r)
}

/*
 * 归并排序
 */
func MergeSort(nums []int) []int {
	res := mergeSort(nums)
	return res
}
func mergeSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	mid := len(nums) / 2
	left := mergeSort(nums[:mid])
	right := mergeSort(nums[mid:])
	return merge(left, right)
}
func merge(left, right []int) (result []int) {
	var l, r int
	for l < len(left) && r < len(right) {
		if right[r] < left[l] {
			result = append(result, right[r])
			r++
		} else {
			result = append(result, left[l])
			l++
		}
	}
	// 合并剩余
	result = append(result, left[l:]...)
	result = append(result, right[r:]...)
	return
}

/*
 * 冒泡排序
 */
func bubbleSort(nums []int) []int {
	length := len(nums)
	for i := 0; i < length; i++ {
		for j := 0; j < length-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
	return nums
}

/*
 * 选择排序
 */
func selectSort(nums []int) []int {
	length := len(nums)
	for i := 0; i < length-1; i++ {
		min := i
		for j := i + 1; j < length; j++ {
			if nums[j] < nums[min] {
				min = j
			}
		}
		nums[i], nums[min] = nums[min], nums[i]
	}
	return nums
}

/*
 * 插入排序
 */
func insertSort(nums []int) []int {
	for i := range nums {
		cur := nums[i]
		index := i - 1
		for index >= 0 && cur < nums[index] {
			nums[index+1] = nums[index]
			index--
		}
		nums[index+1] = cur
	}
	return nums
}

/*
 * 希尔排序
 */
func shellSort(nums []int) []int {
	for gap := len(nums) / 2; gap > 0; gap /= 2 {
		for i := gap; i < len(nums); i++ {
			temp := nums[i]
			j := i - gap
			for j >= 0 && temp < nums[j] {
				nums[j+gap] = nums[j]
				j -= gap
			}
			nums[j+gap] = temp
		}
	}
	return nums
}

func main() {
	//fmt.Println(sortArray([]int{2, 8, 4, 1, 3, 5, 6, 7}))
	fmt.Println(shellSort([]int{2, 8, 4, 1, 3, 5, 6, 7}))
}
