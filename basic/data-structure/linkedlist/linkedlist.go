package linkedlist

type Node struct {
	Data Object
	Next *Node
}

type LinkedList struct {
	Size int
	Head *Node
	Tail *Node
}

func (list *LinkedList) Append(node *Node) bool {
	if node == nil {
		return false
	}

	if list.Size == 0 {
		(*list).Head = node
	} else {
		(*(*list).Tail).Next = node
	}

	(*list).Tail = node
	(*list).Size ++

	return true
}

func (list *LinkedList) Insert(pos int, node *Node) bool {
	if node == nil || pos > (*list).Size {
		return false
	}

	// 如果是 0 则插入头部位置
	if pos == 0 {
		(*node).Next = (*list).Head
		(*list).Head = node
	} else {
		item := (*list).Head
		for i := 0; i < pos; i++ {
			item = (*item).Next
		}
		(*node).Next = (*item).Next
		(*item).Next = item
	}

	(*list).Size ++

	return true
}

func (list *LinkedList) Remove(pos int) bool {
	if pos >= (*list).Size {
		return false
	}

	// 删除头部元素
	if pos == 0 {
		(*list).Head = 	(*(*list).Head).Next

		// 只有一个元素
		if (*list).Size == 1 {
			(*list).Tail = nil
		}
	}

	item := (*list).Head
	for i := 0; i <= pos; i++ {
		item = (*item).Next
	}

	node := (*item).Next
	(*item).Next = (*node).Next

	// 如果是尾部则修改结尾指针
	if i == (*list).Size {
		(*list).Tail = node
	}

	return true
}

func (list *LinkedList) Get(pos int) *Node {
	if pos >= (*list).Size {
		return nil
	}

	item := (*list).Head

	for i = 0; i <= pos; i++ {
		item = (*item).Next
	}

	return item
}