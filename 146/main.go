package main

import "fmt"

type linkNode struct {
	key  int
	prev *linkNode
	next *linkNode
}
type cacheNode struct {
	val   int
	lnode *linkNode
}

type LRUCache struct {
	cache    map[int]cacheNode
	head     *linkNode
	tail     *linkNode
	Capacity int
	Used     int
}

func Constructor(capacity int) LRUCache {
	var lru LRUCache
	lru.Capacity = capacity
	lru.cache = make(map[int]cacheNode)
	return lru
}

func (this *LRUCache) promote(lnode *linkNode) {
	if lnode == this.head {
		return
	}
	if lnode.prev != nil {
		if this.tail == lnode {
			this.tail = lnode.prev
		}
		lnode.prev.next = lnode.next
	}
	if lnode.next != nil {
		lnode.next.prev = lnode.prev
	}
	lnode.next = this.head
	if this.head != nil {
		this.head.prev = lnode
	}
	lnode.prev = nil
	this.head = lnode
}

func (this *LRUCache) Get(key int) int {
	cnode, ok := this.cache[key]
	if !ok || cnode.val == -1 {
		return -1
	}
	//promote the linknode to head
	this.promote(cnode.lnode)
	return cnode.val
}

func (this *LRUCache) Put(key int, value int) {
	cnode, ok := this.cache[key]
	if !ok || cnode.val == -1 {
		// if exceed the capacity delete from tail
		if this.Used == this.Capacity {
			key1 := this.tail.key
			this.tail = this.tail.prev
			if this.tail != nil {
				this.tail.next = nil
			}
			//delete(this.cache, key)
			this.cache[key1] = cacheNode{-1, nil}
			//fmt.Println("delete", key1)
		} else {
			this.Used++
		}
		// create new link node
		lnode := new(linkNode)
		lnode.key = key
		// append the new node to head
		lnode.next = this.head
		if this.Used != 1 {
			this.head.prev = lnode
		} else {
			// for the first node, point tail to it
			this.tail = lnode
		}
		this.head = lnode

		// append {value, *node} to cache
		this.cache[key] = cacheNode{value, lnode}
	} else {
		// promote the linknode to head
		this.promote(cnode.lnode)
		// set value to cache
		this.cache[key] = cacheNode{value, cnode.lnode}
	}
}
func main() {
	fmt.Println("vim-go")
}
