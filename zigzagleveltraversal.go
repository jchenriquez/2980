package zigzagleveltraversal

// TreeNode Simple tree node structure
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// LEFT or RIGHT represent the orientation that the traversal will have at any given time
const (
	LEFT = iota
	RIGHT
)

func traverse(root *TreeNode, queue []*TreeNode, orientation int, lists *[][]int) {

	nQueue := make([]*TreeNode, 0, len(queue)*2)
	nResult := make([]int, 0, len(queue))

	for len(queue) > 0 {
		node := queue[len(queue)-1]
		queue = queue[:len(queue)-1]
		nResult = append(nResult, node.Val)

		if orientation == LEFT {
			if node.Left != nil {
				nQueue = append(nQueue, node.Left)
			}

			if node.Right != nil {
				nQueue = append(nQueue, node.Right)
			}

		} else {
			if node.Right != nil {
				nQueue = append(nQueue, node.Right)
			}
			if node.Left != nil {
				nQueue = append(nQueue, node.Left)
			}
		}
	}

	*lists = append(*lists, nResult)

	if orientation == LEFT {
		traverse(root, nQueue, RIGHT, lists)
	} else {
		traverse(root, nQueue, LEFT, lists)
	}

}

// ZigzagLevelOrder will return an array of the tree in zigzag level order
func ZigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return make([][]int, 0)
	}
	result := make([][]int, 0, 10)
	result = append(result, []int{root.Val})
	resultPtr := &result

	traverse(root, 1, RIGHT, resultPtr)

	return *resultPtr
}
