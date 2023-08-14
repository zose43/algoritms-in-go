package main

type Node struct {
	val        int
	prev, next *Node
}

func NewNode(val int) *Node {
	return &Node{next: nil, prev: nil, val: val}
}

type List struct {
	root *Node
	len  int
}

func (l *List) InsertBefore(node *Node, val int) {
	newNd := NewNode(val)
	if l.root == nil {
		l.PushToHead(newNd)
	} else {
		cur := l.root
		for cur.next != nil {
			if cur.val == node.val {
				if cur.prev == nil {
					l.root = newNd
				}
				newNd.prev = cur.prev
				cur.prev = newNd
				newNd.next = cur
				l.len++
				break
			}
			cur = cur.next
		}
	}
}

func (l *List) InsertAfter(node *Node, val int) {
	newNd := NewNode(val)
	if l.root == nil {
		l.PushToHead(newNd)
	} else {
		cur := l.root
		for cur.next != nil {
			if cur.val == node.val {
				newNd.next = cur.next
				cur.next = newNd
				newNd.prev = cur
				l.len++
				break
			}
			cur = cur.next
		}
	}
}

func (l *List) RemoveFromHead() {
	if l.root != nil {
		head := l.root.next
		head.prev = nil
		l.root = head
		l.len--
	}
}

func (l *List) RemoveFromBack() {
	if l.root != nil {
		cur := l.root
		for cur.next != nil {
			cur = cur.next
		}
		tail := cur.prev
		tail.next = nil
		l.len--
	}
}

func (l *List) PushToHead(node *Node) {
	if l.root == nil {
		l.root = node
	} else {
		head := l.root
		head.prev = node
		node.next = head
		l.root = node
	}
	l.len++
}

func (l *List) PushToBack(node *Node) {
	if l.root == nil {
		l.root = node
	} else {
		cur := l.root
		for cur.next != nil {
			cur = cur.next
		}
		cur.next = node
		node.prev = cur
	}
	l.len++
}

func (l *List) Length() int {
	return l.len
}
