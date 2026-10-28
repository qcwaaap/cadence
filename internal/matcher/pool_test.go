package matcher

import (
    "testing"
)

func feed(values []int) <-chan int {
    ch := make(chan int)
    go func() {
        for _, v := range values {
            ch <- v
        }
        close(ch)
    }()
    return ch
}

func TestRecommendAll(t *testing.T) {
    estimator := SimpleMultiplierEstimator{Multiplier: 2.0}

    tests := []struct {
        name     string
        readings []int
        workers  int
    }{
        {"один воркер", []int{75, 80, 85, 90}, 1},
        {"несколько воркеров", []int{75, 80, 85, 90}, 3},
        {"воркеров больше, чем данных", []int{80}, 8},
        {"ноль воркеров заменяется на один", []int{75, 80}, 0},
        {"пустой вход", nil, 4},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := make(map[int]int)
            for r := range RecommendAll(estimator, feed(tt.readings), tt.workers) {
                got[r.Cadence] = r.BPM
            }

            if len(got) != len(tt.readings) {
                t.Fatalf("ожидали %d результатов, получили %d", len(tt.readings), len(got))
            }
            for _, c := range tt.readings {
                if want := RecommendTrack(estimator, c); got[c] != want {
                    t.Errorf("каденс %d: ожидали %d, получили %d", c, want, got[c])
                }
            }
        })
    }
}

func TestRecommendAllManyReadings(t *testing.T) {
    estimator := SimpleMultiplierEstimator{Multiplier: 2.0}
    readings := make([]int, 1000)
    for i := range readings {
        readings[i] = i
    }

    count := 0
    for range RecommendAll(estimator, feed(readings), 4) {
        count++
    }
    if count != 1000 {
        t.Errorf("ожидали 1000 результатов, получили %d", count)
    }
}
