package piscine

type NodeI struct {
	Data int
	Next *NodeI
}

func SortListInsert(l *NodeI, data_ref int) *NodeI {
	n := &NodeI{Data: data_ref}
	if l == nil {
		return n
	}
	if data_ref < l.Data {
		n.Next = l
		return n
	}
	iterator := l
	for iterator.Next != nil && iterator.Next.Data < data_ref {
		iterator = iterator.Next
	}
	n.Next = iterator.Next
	iterator.Next = n
	return l
}

func ListSort(l *NodeI) *NodeI {
	if l == nil {
		return nil
	}
	return SortListInsert(l, l.Data)
}
