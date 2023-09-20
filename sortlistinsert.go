package piscine

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
	for iterator.Next != nil {
		if data_ref < iterator.Next.Data {
			n.Next = iterator.Next
			iterator.Next = n
			return l
		}
		iterator = iterator.Next
	}
	iterator.Next = n
	return l
}
