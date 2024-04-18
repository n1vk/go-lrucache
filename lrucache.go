package lrucache

import "container/list"

type Node struct {
	Data    int
	Element *list.Element
}

type LRUCache struct {
	Cap    int
	Memory *list.List    // 只负责控制位置
	Hash   map[int]*Node // 负责存储k-v, k-元素索引
}

func New(capacity int) LRUCache {
	return LRUCache{
		Cap:    capacity,
		Memory: list.New(),
		Hash:   make(map[int]*Node),
	}
}

func (l *LRUCache) Put(key int, value int) {

	// 如果有，更新，并移动到最前
	if n, ok := l.Hash[key]; ok {
		n.Data = value
		l.Memory.MoveToFront(n.Element)
		return
	}

	// 如果满，删除最后一个
	if l.Memory.Len() == l.Cap {
		delete(l.Hash, l.Memory.Back().Value.(int))
		l.Memory.Remove(l.Memory.Back())
	}

	l.Hash[key] = &Node{
		Data:    value,
		Element: l.Memory.PushFront(key),
	}

}

func (l *LRUCache) Get(key int) int {

	if v, ok := l.Hash[key]; ok {
		l.Memory.MoveToFront(v.Element)
		return v.Data
	}

	return -1

}
