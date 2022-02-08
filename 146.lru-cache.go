/*
 * @lc app=leetcode id=146 lang=golang
 *
 * [146] LRU Cache
 */

// @lc code=start
type entry struct {
	key   int
	value int
}

type LRUCache struct {
	MaxEntries int
	ll         *list.List
	cache      map[int]*list.Element
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		MaxEntries: capacity,
		ll:         list.New(),
	}
}

func (this *LRUCache) Get(key int) int {
	if this.cache == nil {
		return -1
	}

	if ele, hit := this.cache[key]; hit {
		this.ll.MoveToFront(ele)
		return ele.Value.(*entry).value
	}

	return -1
}

func (this *LRUCache) Put(key int, value int) {

	if this.cache == nil {
		this.cache = make(map[int]*list.Element)
	}
	//使用したときにMoveFront
	v, ok := this.cache[key]
	if ok {
		this.ll.MoveToFront(v)
		v.Value.(*entry).value = value
		return
	}

	ele := this.ll.PushFront(&entry{key, value})
	this.cache[key] = ele

	if this.MaxEntries != 0 && this.ll.Len() > this.MaxEntries {
		this.RemoveOldest()
	}
}

func (this *LRUCache) RemoveOldest() {
	if this.cache == nil {
		return
	}

	ele := this.ll.Back()
	if ele != nil {
		this.removeElement(ele)
	}
}

func (this *LRUCache) removeElement(e *list.Element) {
	this.ll.Remove(e)
	entry := e.Value.(*entry)
	delete(this.cache, entry.key)

}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
// @lc code=end

