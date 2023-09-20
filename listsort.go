package piscine

type NodeI struct {
	Data int
	Next *NodeI
}

func ListSort(l *NodeI) *NodeI {
	if l == nil || l.Next == nil {
		return l
	}

	var prev *NodeI
	slow, fast := l, l

	for fast != nil && fast.Next != nil {
		prev = slow
		slow = slow.Next
		fast = fast.Next.Next
	}

	prev.Next = nil

	return merge(ListSort(l), ListSort(slow))
}

func merge(l1, l2 *NodeI) *NodeI {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	if l1.Data < l2.Data {
		l1.Next = merge(l1.Next, l2)
		return l1
	}

	l2.Next = merge(l1, l2.Next)
	return l2
}
