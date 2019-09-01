package main

import "container/list"

/*
 * @lc app=leetcode id=146 lang=golang
 *
 * [146] LRU Cache
 */

// LRUCache lru
type LRUCache struct {
	l        *list.List
	values   map[int]*list.Element
	keys     map[*list.Element]int
	capacity int
	size     int
}

// Constructor constructor
func Constructor(capacity int) LRUCache {
	l := list.New()
	return LRUCache{
		l:        l,
		keys:     make(map[*list.Element]int, capacity),
		values:   make(map[int]*list.Element, capacity),
		capacity: capacity,
	}
}

// Get get
func (lru *LRUCache) Get(key int) int {
	if v, ok := lru.values[key]; ok {
		lru.l.MoveToFront(v)
		return v.Value.(int)
	}
	return -1
}

// Put put
func (lru *LRUCache) Put(key int, value int) {
	if v, ok := lru.values[key]; ok {
		lru.l.MoveToFront(v)
		v.Value = value
		return
	}
	if lru.size < lru.capacity {
		e := lru.l.PushFront(value)
		lru.keys[e] = key
		lru.values[key] = e
		lru.size++
		return
	}
	back := lru.l.Back()
	removedKey := lru.keys[back]
	delete(lru.keys, back)
	delete(lru.values, removedKey)
	lru.l.MoveToFront(back)
	lru.l.Remove(back)
	e := lru.l.PushFront(value)
	lru.keys[e] = key
	lru.values[key] = e
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
