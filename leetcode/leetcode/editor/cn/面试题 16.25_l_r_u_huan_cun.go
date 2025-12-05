package leetcode

//设计和构建一个“最近最少使用”缓存，该缓存会删除最近最少使用的项目。缓存应该从键映射到值(允许你插入和检索特定键对应的值)，并在初始化时指定最大容量。当缓存
//被填满时，它应该删除最近最少使用的项目。
//
// 它应该支持以下操作： 获取数据 get 和 写入数据 put 。
//
// 获取数据 get(key) - 如果密钥 (key) 存在于缓存中，则获取密钥的值（总是正数），否则返回 -1。 写入数据 put(key, value)
// - 如果密钥不存在，则写入其数据值。当缓存容量达到上限时，它应该在写入新数据之前删除最近最少使用的数据值，从而为新的数据值留出空间。
//
// 示例：
//
//
//LRUCache cache = new LRUCache( 2 /* 缓存容量 */ );
//
//cache.put(1, 1);
//cache.put(2, 2);
//cache.get(1);       // 返回  1
//cache.put(3, 3);    // 该操作会使得密钥 2 作废
//cache.get(2);       // 返回 -1 (未找到)
//cache.put(4, 4);    // 该操作会使得密钥 1 作废
//cache.get(1);       // 返回 -1 (未找到)
//cache.get(3);       // 返回  3
//cache.get(4);       // 返回  4
//
//
// Related Topics 设计 哈希表 链表 双向链表 👍 206 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
type Node struct {
	K    int
	V    int
	Prev *Node
	Next *Node
}
type LRUCache struct {
	Cache    map[int]*Node
	Head     *Node
	Tail     *Node
	Capacity int
	Size     int
}

func Constructor(capacity int) LRUCache {
	lru := LRUCache{
		Cache: make(map[int]*Node, capacity),
		Head: &Node{
			K: -1,
			V: -1,
		},
		Tail: &Node{
			K: -1,
			V: -1,
		},
		Capacity: capacity,
		Size:     0,
	}
	lru.Head.Next = lru.Tail
	lru.Tail.Prev = lru.Head
	return lru
}

func (t *LRUCache) removeNode(n *Node) {
	n.Prev.Next = n.Next
	n.Next.Prev = n.Prev
}

func (t *LRUCache) addToHead(n *Node) {
	n.Next = t.Head.Next
	n.Prev = t.Head
	t.Head.Next.Prev = n
	t.Head.Next = n
}

func (t *LRUCache) moveToHead(n *Node) {
	t.removeNode(n)
	t.addToHead(n)
}

func (t *LRUCache) removeTail() *Node {
	node := t.Tail.Prev
	t.removeNode(node)
	return node
}

func (t *LRUCache) Get(key int) int {
	item, ok := t.Cache[key]
	if !ok {
		return -1
	}
	t.moveToHead(item)
	return item.V
}

func (t *LRUCache) Put(key int, value int) {
	if item, ok := t.Cache[key]; ok {
		item.V = value
		t.moveToHead(item)
	} else {
		newNode := &Node{
			K: key,
			V: value,
		}
		t.Cache[key] = newNode
		t.addToHead(newNode)
		t.Size++

		if t.Size > t.Capacity {
			tail := t.removeTail()
			delete(t.Cache, tail.K)
			t.Size--
		}
	}
}

//	terst
/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
//leetcode submit region end(Prohibit modification and deletion)
