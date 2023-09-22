package psicne

func BTreeMax(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	for root.Right != nil {
		root = root.Right
	}
	return root
}
