package main

// 设计一个 LRU 缓存机制
// 双向链表 + 哈希表 实现 在双向链表的实现中，使用一个伪头部（dummy head）和伪尾部（dummy tail）标记界限，这样在添加节点和删除节点的时候就不需要检查相邻的节点是否存在。
// 靠近头部的键值对是最近使用的，而靠近尾部的键值对是最久未使用的。对于 Get 和 Put 操作的时间复杂度都是 O(1)，空间复杂度是 O(capacity)
type LRUCache struct {
	length, cap int
	hash        map[int]*DLinkedNode
	head, tail  *DLinkedNode
}

// 双向链表节点 包含指向前后节点的指针
type DLinkedNode struct {
	key, value int
	prev, next *DLinkedNode
}

// 生成链表节点
func initLinkedNode(key, val int) *DLinkedNode {
	return &DLinkedNode{key: key, value: val}
}

func Constructor(capacity int) LRUCache {
	cache := LRUCache{
		length: 0,
		cap:    capacity, // lru 缓存的容量
		hash:   map[int]*DLinkedNode{},
		head:   &DLinkedNode{},
		tail:   &DLinkedNode{},
	}
	cache.head.next = cache.tail
	cache.tail.prev = cache.head
	return cache
}

func (this *LRUCache) Get(key int) int {
	// key 不存在 => 返回 -1
	if _, ok := this.hash[key]; !ok {
		return -1
	}
	// move to head
	node := this.hash[key]
	this.moveToHead(node)
	return node.value
}

func (this *LRUCache) Put(key int, value int) {
	if _, ok := this.hash[key]; ok {
		// node 已存在的情况
		this.hash[key].value = value
		node := this.hash[key]
		this.moveToHead(node)
	} else {
		// node 不存在
		node := initLinkedNode(key, value)
		this.hash[key] = node
		this.addToHead(node)
		this.length++
		if this.length > this.cap {
			// 有淘汰策略
			this.removeTail()
		}
	}
}

func (this *LRUCache) addToHead(node *DLinkedNode) {
	node.prev = this.head
	node.next = this.head.next
	this.head.next.prev = node
	this.head.next = node
}

func (this *LRUCache) moveToHead(node *DLinkedNode) {
	this.removeNode(node)
	this.addToHead(node)
}

// 删除节点
func (this *LRUCache) removeNode(node *DLinkedNode) {
	next := node.next
	prev := node.prev
	prev.next = next
	next.prev = prev
}

// 缓存淘汰 淘汰尾部节点
func (this *LRUCache) removeTail() *DLinkedNode {
	last := this.tail.prev
	this.removeNode(last)
	delete(this.hash, last.key)
	this.length--
	return last
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
