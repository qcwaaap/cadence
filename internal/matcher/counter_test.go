package matcher

import (
    "sync"
    "testing"
)

func TestCounterRace(t *testing.T) {
    var c Counter
    var wg sync.WaitGroup

    for i := 0; i < 1000; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            c.Increment()
        }()
    }
    wg.Wait()

    if c.Value() != 1000 {
        t.Errorf("ожидали 1000, получили %d", c.Value())
    }
}
