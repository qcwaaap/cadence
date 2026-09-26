package main

import (
	"cadence/internal/activity"
	"cadence/internal/matcher"
	"fmt"
)

func main() {
	// act, err := activity.FetchActivityFromStrava(123)
	// if err != nil {
	// 	wrapped := fmt.Errorf("failed to sync activity: %w", err)
	// 	var target *activity.StravaApiError
	// 	if errors.As(wrapped, &target) {
	// 		if target.StatusCode == 429 {
	// 			fmt.Println("превышен лимит запросов, повторите позже")
	// 		} else {
	// 			fmt.Println("неизвестная ошибка API")
	// 		}
	// 	}
	// 	return
	// }
	// fmt.Println("Activity loaded:", act)

	readings := []int{75, 80, 85, 90}
	estimator := matcher.SimpleMultiplierEstimator{Multiplier: 2.0}
	ch := activity.SimulateCadenceReadings(readings)
	for r := range ch {
		bpm := matcher.RecommendTrack(estimator, r)
		fmt.Println("Каденс:", r, "→ Рекомендуемый BPM:", bpm)
	}

}
