package tencent_ai_agent

import (
	"math"
	"sync"
)

type IData struct {
	Val int
	A   int        // 访问计数器
	mu  sync.Mutex // 每个数据项的独立锁
}

type Lru struct {
	Data map[string]*IData // 改用普通map+读写锁
	Size int               // 最大容量
	mu   sync.RWMutex      // 全局读写锁
}

func NewLru(size int) *Lru {
	return &Lru{
		Data: make(map[string]*IData),
		Size: size,
	}
}

func (l *Lru) Get(key string) int {
	l.mu.RLock()
	item, exists := l.Data[key]
	l.mu.RUnlock()

	if !exists {
		return -1
	}

	item.mu.Lock()
	defer item.mu.Unlock()
	item.A++
	return item.Val
}

func (l *Lru) Put(key string, value int) {
	l.mu.Lock()
	defer l.mu.Unlock()

	// 更新现有条目
	if item, exists := l.Data[key]; exists {
		item.mu.Lock()
		defer item.mu.Unlock()
		item.A += 2
		item.Val = value
		return
	}

	// 添加新条目
	newItem := &IData{
		Val: value,
		A:   2,
	}
	l.Data[key] = newItem

	// 执行LRU淘汰策略
	if l.Size > 0 && len(l.Data) > l.Size {
		l.evict()
	}
}

func (l *Lru) evict() {
	var oldestKey string
	minA := math.MaxInt64

	// 遍历查找最小A值的条目
	for k, v := range l.Data {
		v.mu.Lock()
		if a := v.A; a < minA {
			minA = a
			oldestKey = k
		}
		v.mu.Unlock()
	}

	// 删除最旧条目
	if oldestKey != "" {
		delete(l.Data, oldestKey)
	}
}
