package main

/*
 * 链表
 * null/nil 异常处理
 * dummy node 哑巴节点
 * 快慢指针
 * 插入一个节点到排序链表
 * 从一个链表中移除一个节点
 * 翻转链表
 * 合并两个链表
 * 找到链表的中间节点
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortArrToListNode(arr []int) *ListNode {
	head := &ListNode{Val: arr[0]}
	other := head
	for i := 1; i < len(arr); i++ {
		temp := &ListNode{Val: arr[i]}
		other.Next = temp
		other = other.Next
	}
	return head
}

/*
 * 删除排序链表中的重复元素
 * 存在一个按升序排列的链表，给你这个链表的头节点 head ，请你删除所有重复的元素，使每个元素只出现一次
 * 返回同样按升序排列的结果链表
 */
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	cur := head
	for cur.Next != nil {
		if cur.Val == cur.Next.Val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return head
}

/*
 * 删除排序链表中的重复元素 II
 * 存在一个按升序排列的链表，给你这个链表的头节点 head ，请你删除链表中所有存在数字重复情况的节点，只保留原始链表中 没有重复出现 的数字。
 * 返回同样按升序排列的结果链表
 */
func deleteDuplicates2(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	dummy := &ListNode{
		Val:  0,
		Next: head,
	}
	head = dummy
	var del int
	for head.Next != nil && head.Next.Next != nil {
		if head.Next.Val == head.Next.Next.Val {
			del = head.Next.Val
			for head.Next != nil && head.Next.Val == del {
				head.Next = head.Next.Next
			}
		} else {
			head = head.Next
		}
	}
	return dummy.Next
}

/*
 * 反转链表
 * 给你单链表的头节点 head ，请你反转链表，并返回反转后的链表
 * 双指针迭代
 */
func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	var pre *ListNode
	cur := head
	for cur != nil {
		temp := cur.Next
		cur.Next = pre
		pre = cur
		cur = temp
	}
	return pre
}

/*
 * 反转链表
 * 给你单链表的头节点 head ，请你反转链表，并返回反转后的链表
 * 递归解法
 */
func reverseList1(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	cur := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return cur
}

/*
 * 反转链表 II
 * 给你单链表的头指针 head 和两个整数 left 和 right ，其中 left <= right
 * 请你反转从位置 left 到位置 right 的链表节点，返回 反转后的链表
 * 迭代-头插法
 */
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	dummy := &ListNode{Val: -1, Next: head}
	pre := dummy
	for i := 1; i < left; i++ {
		pre = pre.Next
	}
	cur := pre.Next
	for i := 0; i < right-left; i++ {
		cur.Next, cur.Next.Next, pre.Next = cur.Next.Next, pre.Next, cur.Next
	}
	return dummy.Next
}

/*
 * 合并两个有序链表
 * 将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的
 * 迭代解法
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{Val: -1}
	temp := dummy
	for l1 != nil && l2 != nil {
		if l1.Val > l2.Val {
			temp.Next = l2
			l2 = l2.Next
		} else {
			temp.Next = l1
			l1 = l1.Next
		}
		temp = temp.Next
	}
	if l1 == nil {
		temp.Next = l2
	} else {
		temp.Next = l1
	}
	return dummy.Next
}

/*
 * 合并两个有序链表
 * 将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的
 * 递归解法
 */
func mergeTwoLists1(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val > l2.Val {
		l2.Next = mergeTwoLists(l1, l2.Next)
		return l2
	} else {
		l1.Next = mergeTwoLists(l1.Next, l2)
		return l1
	}
}

/*
 * 分隔链表
 * 给你一个链表的头节点 head 和一个特定值 x ，请你对链表进行分隔，使得所有 小于 x 的节点都出现在 大于或等于 x 的节点之前。
 * 你应当 保留 两个分区中每个节点的初始相对位置。
 */
func partition(head *ListNode, x int) *ListNode {
	left := new(ListNode)
	right := new(ListNode)
	tempLeft := left
	tempRight := right
	for head != nil {
		if head.Val < x {
			tempLeft.Next = head
			tempLeft = tempLeft.Next
		} else {
			tempRight.Next = head
			tempRight = tempRight.Next
		}
		head = head.Next
	}
	tempRight.Next = nil
	tempLeft.Next = right.Next
	return left.Next
}

/*
 * 两两交换链表中的节点
 * 给定一个链表，两两交换其中相邻的节点，并返回交换后的链表。
 * 你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换
 */
func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{Val: -1, Next: head}
	pre := dummy
	cur := dummy.Next
	for cur != nil && cur.Next != nil {
		cur.Next, cur.Next.Next, pre.Next = cur.Next.Next, pre.Next, cur.Next
		pre = pre.Next.Next
		cur = cur.Next
	}
	return dummy.Next
}

func main() {
	//deleteDuplicates(sortArrToListNode([]int{1, 1, 2, 3, 3}))
	//reverseList(sortArrToListNode([]int{1, 2, 3, 4, 5}))
	//reverseBetween(sortArrToListNode([]int{9, 7, 2, 5, 4, 3, 6}), 2, 4)
	partition(sortArrToListNode([]int{1, 4, 3, 2, 5, 2}), 3)
}
