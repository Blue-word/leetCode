package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	//res := postorderTraversal(sortedArrayToBST([]int{-10, -3, 0, 5, 9}))
	//res := MergeSort([]int{8, 4, 5, 7, 1, 3, 6, 2})
	res := minDepth(sortedArrayToBST([]int{2, 0, 3, 0, 4, 0, 5, 0, 6}))
	fmt.Println(res)
}

func sortedArrayToBST(nums []int) *TreeNode {
	return buildBinaryTree(nums, 0, len(nums)-1)
}

// 将有序数组转换为二叉搜索树
func buildBinaryTree(nums []int, left, right int) *TreeNode {
	if left > right {
		return nil
	}
	mid := (left + right) / 2
	root := &TreeNode{Val: nums[mid]}
	root.Left = buildBinaryTree(nums, left, mid-1)
	root.Right = buildBinaryTree(nums, mid+1, right)
	return root
}

/*
 * 二叉树
 * 前序遍历：先访问根节点，再前序遍历左子树，再前序遍历右子树
 * 中序遍历：先中序遍历左子树，再访问根节点，再中序遍历右子树
 * 后序遍历：先后序遍历左子树，再后序遍历右子树，再访问根节点
 */
// 前序递归
func preorderTraversalRecursion(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Println(root.Val)                  // 根节点
	preorderTraversalRecursion(root.Left)  // 左子树
	preorderTraversalRecursion(root.Right) // 右子树
}

// 前序非递归
func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var res []int
	stack := make([]*TreeNode, 0)
	for root != nil {
		for root != nil {
			res = append(res, root.Val)
			stack = append(stack, root.Left)
			root = root.Left
		}
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		root = node.Right
	}
	return res
}

// 中序非递归
func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var res []int
	stack := make([]*TreeNode, 0)
	for root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, node.Val)
		root = node.Right
	}
	return res
}

// 后续非递归
func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var res []int
	stack := make([]*TreeNode, 0)
	var pre *TreeNode
	for root != nil || len(stack) != 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		node := stack[len(stack)-1]
		if node.Right == nil || node.Right == pre {
			res = append(res, node.Val)
			stack = stack[:len(stack)-1]
			pre = node
		} else {
			root = node.Right
		}
	}
	return res
}

// DFS深度优先搜索-从上到下
func dfsTraversal(root *TreeNode) []int {
	result := make([]int, 0)
	dfs(root, &result)
	return result
}

// V1：深度遍历，结果指针作为参数传入到函数内部
func dfs(root *TreeNode, result *[]int) {
	if root == nil {
		return
	}
	*result = append(*result, root.Val)
	dfs(root.Left, result)
	dfs(root.Right, result)
}

// DFS 深度搜索-从下向上（分治法）
func divideTraversal(root *TreeNode) []int {
	result := divideAndConquer(root)
	return result
}

// V2：通过分治法遍历
func divideAndConquer(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var res []int
	divideAndConquer(root.Left)
	divideAndConquer(root.Right)
	res = append(res, root.Val)
	return res
}

// BFS 层序遍历-自顶向下
func levelOrder(root *TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}
	var queue []*TreeNode
	queue = append(queue, root)
	for len(queue) > 0 {
		var temp []int
		length := len(queue)
		for i := 0; i < length; i++ {
			node := queue[0]
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			temp = append(temp, node.Val)
		}
		res = append(res, temp)
	}
	return res
}

// BFS 层序遍历-自底向上
func levelOrderBottom(root *TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}
	var queue []*TreeNode
	queue = append(queue, root)
	for len(queue) > 0 {
		var temp []int
		length := len(queue)
		for i := 0; i < length; i++ {
			node := queue[0]
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			temp = append(temp, node.Val)
		}
		res = append(res, temp)
	}
	//var newRes [][]int
	//for i := len(res) - 1; i >= 0; i-- {
	//	newRes = append(newRes, res[i])
	//}
	// 反转
	for i := 0; i < len(res)/2; i++ {
		res[i], res[len(res)-i-1] = res[len(res)-i-1], res[i]
	}
	return res
}

// 二叉树的最小深度
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var count []int
	var queue []*TreeNode
	queue = append(queue, root)
	count = append(count, 1)
	for i := 0; i < len(queue); i++ {
		node := queue[i]
		depth := count[i]
		if node.Left == nil && node.Right == nil {
			return depth
		}
		if node.Left != nil {
			queue = append(queue, root.Left)
			count = append(count, depth+1)
		}
		if node.Right != nil {
			queue = append(queue, root.Right)
			count = append(count, depth+1)
		}
	}
	return 0
}
