package cache

import "container/list"

// NewLRUCache returns a new non-concurrent-safe LRU cache with input capacity
func NewLRUCache(capacity uint64) *LRU {
	return &LRU{
		Cap:   capacity,
		l:     list.New(),
		items: make(map[interface{}]*list.Element),
	}
}

// Add adds a value to the cache
func (l *LRU) Add(key, value interface{}) {
	if f, o := l.items[key]; o {
		l.l.MoveToFront(f)
