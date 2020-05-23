package tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func KthSmallest(root *TreeNode, k int) int {
	return kthSmallest(root, k)
}

func kthSmallest(root *TreeNode, k int) int {
	_, _, result := find(root, k)
	return result
}

func find(node *TreeNode, k int) (found bool, nodeCount int, kth int) {
	leftNodeCount, rightNodeCount := 0, 0
	if node.Left != nil {
		var leftKth int
		found, leftNodeCount, leftKth = find(node.Left, k)
		if found {
			return found, 0, leftKth
		}
	}
	if leftNodeCount == k-1 {
		return true, 0, node.Val
	}
	if node.Right != nil {
		var rightKth int
		found, rightNodeCount, rightKth = find(node.Right, k-leftNodeCount-1)
		if found {
			return found, 0, rightKth
		}
	}
	return false, leftNodeCount + rightNodeCount + 1, 0
}
