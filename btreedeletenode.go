package piscine

func BTreeDeleteNode(root, node *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if node == nil {
		return root
	}
	if node.Left == nil {
		root = BTreeTransplant(root, node, node.Right)
	}
	if node.Right == nil {
		root = BTreeTransplant(root, node, node.Left)
	}
	if node.Left != nil && node.Right != nil {
		rplc := BTreeMin(node.Right)
		if rplc.Parent != node {
			root = BTreeTransplant(root, rplc, rplc.Right)
			rplc.Right = node.Right
			rplc.Right.Parent = rplc
		}
		root = BTreeTransplant(root, node, rplc)
		rplc.Left = node.Left
		rplc.Left.Parent = rplc
	}
	return root
}
