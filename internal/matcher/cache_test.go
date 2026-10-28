package matcher

import (
	"sync"
	"testing"
)

func TestCacheGetSet(t *testing.T) {
	c := NewRecommendationCache(3)

	if _, ok := c.Get(80); ok {
		t.Error("пустой кэш не должен ничего возвращать")
	}
	c.Set(80, 160)
	if bpm, ok := c.Get(80); !ok || bpm != 160 {
		t.Errorf("ожидали 160, true; получили %d, %v", bpm, ok)
	}
	c.Set(80, 164)
	if bpm, _ := c.Get(80); bpm != 164 {
		t.Errorf("значение не обновилось: %d", bpm)
	}
	if c.Len() != 1 {
		t.Errorf("повторный Set не должен добавлять запись, Len=%d", c.Len())
	}
}

func TestCacheEvictsOldest(t *testing.T) {
	c := NewRecommendationCache(2)
	c.Set(70, 140)
	c.Set(80, 160)
	c.Set(90, 180) // вытесняет 70

	if _, ok := c.Get(70); ok {
		t.Error("70 должна быть вытеснена")
	}
	if _, ok := c.Get(80); !ok {
		t.Error("80 должна остаться")
	}
	if _, ok := c.Get(90); !ok {
		t.Error("90 должна остаться")
	}
	if c.Len() != 2 {
		t.Errorf("ожидали Len=2, получили %d", c.Len())
	}
}

func TestCacheConcurrent(t *testing.T) {
	c := NewRecommendationCache(10)
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			c.Set(n%50, n)
			c.Get(n % 50)
		}(i)
	}
	wg.Wait()

	if c.Len() > 10 {
		t.Errorf("кэш вырос выше лимита: %d", c.Len())
	}
}
