package piscine

func ListReverse(l *List) {
	if l.Head == nil {
		return
	}
	list := &List{}
	it := l.Head
	for it != nil {
		ListPushFront(list, it.Data)
		it = it.Next
	}
	l.Head = list.Head
}
