package leetcode

import "container/list"

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
type LRUCache struct {
	cache    map[int]*list.Element
	list     *list.List
	capacity int
}

type Item struct {
	Key   int
	Value int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{list: list.New(), capacity: capacity, cache: make(map[int]*list.Element)}
}

func (t *LRUCache) Get(key int) int {
	data, ok := t.cache[key]
	if ok {
		t.list.MoveToFront(data)
		return data.Value.(*Item).Value
	}
	return -1
}

func (t *LRUCache) getItem(key int) *list.Element {
	data, ok := t.cache[key]
	if ok {
		return data
	}
	return nil
}

func (t *LRUCache) Put(key int, value int) {
	item := t.getItem(key)
	if item != nil {
		item.Value.(*Item).Value = value
		t.list.MoveToFront(item)
		return
	}

	if t.list.Len() >= t.capacity {
		// 删除最少使用的
		tail := t.list.Back()
		if tail != nil {
			delete(t.cache, tail.Value.(*Item).Key)
			t.list.Remove(tail)
		}
	}
	t.cache[key] = t.list.PushFront(&Item{key, value})
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
//leetcode submit region end(Prohibit modification and deletion)
