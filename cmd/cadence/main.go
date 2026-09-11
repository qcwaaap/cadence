package main

import (
	"cadence/internal/activity"
	"errors"
	"fmt"
)

func main() {
	act, err := activity.FetchActivityFromStrava(123)
	if err != nil {
		wrapped := fmt.Errorf("failed to sync activity: %w", err)

		var target *activity.StravaApiError
		if errors.As(wrapped, &target) {
			if target.StatusCode == 429 {
				fmt.Println("превышен лимит запросов, повторите позже")
			} else {
				fmt.Println("неизвестная ошибка API")
			}
		}
		return
	}

	fmt.Println("Activity loaded:", act)
}
