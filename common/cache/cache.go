package cache

// New returns a new concurrent-safe LRU cache with input capacity
func New(capacity uint64) *LRUCache {
	return &LRUCache{
		lru: NewLRUCache(capacity),
	}
}

// Add new entry to Cache return true if entry removed
func (l *LRUCache) Add(k, v interface{}) {
	l.m.Lock()
