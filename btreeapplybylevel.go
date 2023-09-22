package piscine

func BTreeApplyLevel(root *TreeNode, level int, f func(...interface{}) (int, error)) {
	if root == nil {
		return
	}
	if level == 1 {
		f(root.Data)
	} else if level > 1 {
		BTreeApplyLevel(root.Left, level-1, f)
		BTreeApplyLevel(root.Right, level-1, f)
	}
}

func BTreeApplyByLevel(root *TreeNode, f func(...interface{}) (int, error)) {
	if root == nil {
		return
	}
	level := BTreeLevelCount(root)
	for i := 1; i <= level; i++ {
		BTreeApplyLevel(root, i, f)
	}
}
