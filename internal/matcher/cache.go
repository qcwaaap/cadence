package matcher

import "sync"

// RecommendationCache хранит уже посчитанные BPM по каденсу.
// Размер ограничен: когда место кончилось, вытесняется самая старая запись.
type RecommendationCache struct {
	mu    sync.RWMutex
	items map[int]int
	order []int // кольцо ключей в порядке добавления
	head  int   // позиция самого старого ключа, когда кольцо заполнено
	max   int
}

func NewRecommendationCache(max int) *RecommendationCache {
	if max < 1 {
		max = 1
	}
	return &RecommendationCache{
		items: make(map[int]int, max),
		order: make([]int, 0, max),
		max:   max,
	}
}

func (c *RecommendationCache) Get(cadence int) (int, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	bpm, ok := c.items[cadence]
	return bpm, ok
}

func (c *RecommendationCache) Set(cadence, bpm int) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if _, ok := c.items[cadence]; ok {
		c.items[cadence] = bpm
		return
	}
	if len(c.order) < c.max {
		c.order = append(c.order, cadence)
	} else {
		delete(c.items, c.order[c.head])
		c.order[c.head] = cadence
		c.head = (c.head + 1) % c.max
	}
	c.items[cadence] = bpm
}

func (c *RecommendationCache) Len() int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return len(c.items)
}
