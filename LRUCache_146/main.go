package main

type cacheItem struct {
	data  int
	key   int
	left  *cacheItem
	right *cacheItem
}

type LRUCache struct {
	items    map[int]*cacheItem
	head     *cacheItem
	tail     *cacheItem
	capacity uint32
}

func NewLRUCache(capacity uint32) LRUCache {
	cache := LRUCache{
		items:    make(map[int]*cacheItem),
		head:     &cacheItem{},
		tail:     &cacheItem{},
		capacity: capacity,
	}

	cache.tail.left = cache.head
	cache.head.right = cache.tail

	return cache
}

func (c *LRUCache) Put(key int, value int) {
	if item, ok := c.items[key]; ok {
		item.data = value
		c.Remove(item)
		c.Add(item)
	} else {
		item := &cacheItem{
			data: value,
			key:  key,
		}
		c.items[key] = item
		c.Add(item)
	}
}

func (c *LRUCache) Remove(item *cacheItem) {
	left := item.left
	right := item.right
	left.right = right
	right.left = left
	item.left = nil
	item.right = nil
}

func (c *LRUCache) Add(item *cacheItem) {
	item.right = c.head.right
	c.head.right.left = item
	c.head.right = item
	item.left = c.head

	if len(c.items) > int(c.capacity) {
		delete(c.items, c.tail.left.key)
		c.Remove(c.tail.left)
	}
}

func (c *LRUCache) Get(key int) int {
	item := c.items[key]
	if item == nil {
		return -1
	}

	c.Remove(item)
	c.Add(item)
	return item.data
}

func (c *LRUCache) moveToTop(item *cacheItem) {
	if c.head == nil {
		c.head = item
		c.tail = item
		return
	}

	if c.head == item {
		return
	}

	if item.left != nil && item.right != nil {
		item.left.right = item.right
		item.right.left = item.left
	} else if item.right == nil && item.left != nil {
		item.left.right = nil
		c.tail = item.left

		item.left = nil
	}

	c.head.left, item.right = item, c.head
	c.head = item

	if len(c.items) > int(c.capacity) {
		c.removeLast()
	}
}

func (c *LRUCache) removeLast() {
	if c.tail == nil {
		return
	}

	c.tail.left.right = nil
	delete(c.items, c.tail.key)
	c.tail = c.tail.left
}
