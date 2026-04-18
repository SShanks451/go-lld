package main

import (
	"container/list"
	"fmt"
)

type entry struct {
	key  int
	val  int
	freq int
}

type LFUCache struct {
	capacity   int
	size       int
	minFreq    int
	keyToNode  map[int]*list.Element
	freqToList map[int]*list.List
}

func NewLFUCache(capacity int) *LFUCache {
	return &LFUCache{
		capacity:   capacity,
		keyToNode:  make(map[int]*list.Element),
		freqToList: make(map[int]*list.List),
	}
}

func (l *LFUCache) Get(key int) int {
	if l.capacity == 0 {
		return -1
	}

	node, ok := l.keyToNode[key]
	if ok {
		l.touch(node)
		return node.Value.(*entry).val
	}
	return -1
}

func (l *LFUCache) Put(key, val int) {
	if l.capacity == 0 {
		return
	}

	if node, ok := l.keyToNode[key]; ok {
		e := node.Value.(*entry)
		e.val = val
		l.touch(node)
		return
	}

	if l.size == l.capacity {
		l.evictLFU()
	}

	e := &entry{
		key:  key,
		val:  val,
		freq: 1,
	}
	if _, ok := l.freqToList[1]; !ok {
		l.freqToList[1] = list.New()
	}

	node := l.freqToList[1].PushFront(e)
	l.keyToNode[key] = node
	l.minFreq = 1
	l.size++
}

func (l *LFUCache) touch(node *list.Element) {
	e := node.Value.(*entry)
	oldFreq := e.freq
	l.freqToList[oldFreq].Remove(node)

	if l.freqToList[oldFreq].Len() == 0 {
		delete(l.freqToList, oldFreq)
		if l.minFreq == oldFreq {
			l.minFreq++
		}
	}

	e.freq++
	if _, ok := l.freqToList[e.freq]; !ok {
		l.freqToList[e.freq] = list.New()
	}

	l.keyToNode[e.key] = l.freqToList[e.freq].PushFront(e)
}

func (l *LFUCache) evictLFU() {
	lst := l.freqToList[l.minFreq]
	victim := lst.Back()
	if victim == nil {
		return
	}
	lst.Remove(victim)

	if lst.Len() == 0 {
		delete(l.freqToList, l.minFreq)
	}

	delete(l.keyToNode, victim.Value.(*entry).key)
	l.size--
}

func main() {
	fmt.Println("========================================")
	fmt.Println("   LFU Cache — Demo Operations")
	fmt.Println("========================================")

	// ---- 1. Basic Insert & Get ----
	fmt.Println("\n--- 1. Basic Insert & Get (capacity=3) ---")
	cache := NewLFUCache(3)
	cache.Put(1, 10)
	cache.Put(2, 20)
	cache.Put(3, 30)

	fmt.Printf("  Get(1) = %d\n", cache.Get(1)) // 10, freq(1) bumps to 2
	fmt.Printf("  Get(2) = %d\n", cache.Get(2)) // 20, freq(2) bumps to 2
	fmt.Printf("  Get(9) = %d\n", cache.Get(9)) // -1, miss
	fmt.Println("  After Gets:")

	// ---- 2. Eviction: least frequent key removed ----
	fmt.Println("\n--- 2. Eviction (key 3 has freq=1, others have freq=2) ---")
	cache.Put(4, 40) // should evict key=3 (lowest freq=1)

	fmt.Printf("  Get(3) = %d  (should be -1, evicted)\n", cache.Get(3))
	fmt.Printf("  Get(4) = %d\n", cache.Get(4))

	// ---- 3. LRU tiebreaker among same frequency ----
	fmt.Println("\n--- 3. LRU Tiebreaker (capacity=2) ---")
	cache2 := NewLFUCache(2)
	cache2.Put(10, 100)
	cache2.Put(20, 200)
	// Both have freq=1. key=10 was inserted first → it's the LRU.
	fmt.Println("  Before eviction:")
	cache2.Put(30, 300) // evicts key=10 (LRU among freq=1)
	fmt.Println("  After inserting key=30:")
	fmt.Printf("  Get(10) = %d  (should be -1)\n", cache2.Get(10))
	fmt.Printf("  Get(20) = %d\n", cache2.Get(20))

	// ---- 4. Update existing key ----
	fmt.Println("\n--- 4. Update Existing Key ---")
	cache3 := NewLFUCache(2)
	cache3.Put(1, 100)
	cache3.Put(2, 200)
	fmt.Println("  Before update:")

	cache3.Put(1, 999) // update key=1's value, freq bumps to 2
	fmt.Println("  After Put(1, 999):")
	fmt.Printf("  Get(1) = %d  (should be 999)\n", cache3.Get(1))

	// Now insert key=3 → should evict key=2 (freq=1), not key=1 (freq=3)
	cache3.Put(3, 300)
	fmt.Println("  After inserting key=3:")

	// ---- 5. Repeated access builds frequency ----
	fmt.Println("\n--- 5. Frequency Building (capacity=3) ---")
	cache4 := NewLFUCache(3)
	cache4.Put(1, 10)
	cache4.Put(2, 20)
	cache4.Put(3, 30)

	// Hammer key=1 five times
	for i := 0; i < 5; i++ {
		cache4.Get(1)
	}
	// Access key=2 twice
	cache4.Get(2)
	cache4.Get(2)
	fmt.Println("  After hammering key=1 (freq=6) and key=2 (freq=3):")

	// Insert key=4 → evicts key=3 (freq=1)
	cache4.Put(4, 40)
	fmt.Println("  After inserting key=4:")

	// Insert key=5 → evicts key=4 (freq=1, the newest low-freq entry)
	cache4.Put(5, 50)
	fmt.Println("  After inserting key=5:")

	// ---- 6. Zero capacity edge case ----
	fmt.Println("\n--- 6. Zero Capacity Edge Case ---")
	cache5 := NewLFUCache(0)
	cache5.Put(1, 10)
	fmt.Printf("  Get(1) = %d  (should be -1, cache can't hold anything)\n", cache5.Get(1))

	// ---- 7. Single capacity ----
	fmt.Println("\n--- 7. Single Capacity ---")
	cache6 := NewLFUCache(1)
	cache6.Put(1, 10)
	fmt.Printf("  Get(1) = %d\n", cache6.Get(1))
	cache6.Put(2, 20) // evicts key=1
	fmt.Printf("  Get(1) = %d  (should be -1)\n", cache6.Get(1))
	fmt.Printf("  Get(2) = %d\n", cache6.Get(2))

	fmt.Println("\n========================================")
	fmt.Println("   All operations completed!")
	fmt.Println("========================================")
}
