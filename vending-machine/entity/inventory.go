package entity

import "sync"

type Inventory struct {
	mu       sync.RWMutex
	itemMap  map[string]*Item
	stockMap map[string]int
}

func NewInventory() *Inventory {
	return &Inventory{
		itemMap:  make(map[string]*Item),
		stockMap: make(map[string]int),
	}
}

func (i *Inventory) AddItem(code string, item *Item, quantity int) {
	i.mu.Lock()
	defer i.mu.Unlock()

	i.itemMap[code] = item
	i.stockMap[code] = quantity
}

func (i *Inventory) GetItem(code string) *Item {
	i.mu.RLock()
	defer i.mu.RUnlock()

	item, ok := i.itemMap[code]
	if !ok {
		return nil
	}

	return item
}

func (i *Inventory) IsAvailable(code string) bool {
	i.mu.RLock()
	defer i.mu.RUnlock()

	quantity, ok := i.stockMap[code]
	if !ok {
		return false
	}

	if quantity > 0 {
		return true
	}
	return false
}

func (i *Inventory) ReduceStock(code string) {
	i.mu.Lock()
	defer i.mu.Unlock()

	_, ok := i.stockMap[code]
	if !ok {
		return
	}

	if i.stockMap[code] > 0 {
		i.stockMap[code]--
	}
}
