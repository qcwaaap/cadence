package matcher

import "sync"

// Result это ответ на одно показание каденса.
type Result struct {
    Cadence int
    BPM     int
}

// RecommendAll читает каденс из readings и считает BPM в фиксированном
// числе воркеров. Порядок результатов не гарантирован. Канал результатов
// закрывается, когда readings закрыт и все воркеры закончили.
func RecommendAll(estimator BPMEstimator, readings <-chan int, workers int) <-chan Result {
    if workers < 1 {
        workers = 1
    }
    results := make(chan Result)

    var wg sync.WaitGroup
    for i := 0; i < workers; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            for cadence := range readings {
                results <- Result{
                    Cadence: cadence,
                    BPM:     RecommendTrack(estimator, cadence),
                }
            }
        }()
    }

    // Закрывать results должен только тот, кто знает, что писать больше некому.
    go func() {
        wg.Wait()
        close(results)
    }()

    return results
}
