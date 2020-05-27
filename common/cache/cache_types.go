package cache

import (
	"container/list"
	"sync"
)

type LRUCache struct {
	lru *LRU
	m   sync.Mutex
}
type LRU struct {
	Cap   uint64
	l     *list.List
	items map[interface{}]*list.Element
}
type item struct {
	key   interface{}
	value interface{}
}
