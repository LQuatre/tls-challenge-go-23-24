package piscine

func ListRemoveIf(l *List, data_ref interface{}) {
	if l.Head == nil {
		return
	}
	it := l.Head
	for it != nil {
		if it.Data == data_ref {
			l.Head = l.Head.Next
			it = l.Head
		} else if it.Next != nil && it.Next.Data == data_ref {
			it.Next = it.Next.Next
		} else {
			it = it.Next
		}
	}
}
