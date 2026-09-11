package main

import (
	"cadence/internal/activity"
	"fmt"
)

func main() {
	act, err := activity.FetchActivityFromStrava(123)
	if err != nil {
		wrapped := fmt.Errorf("failed to sync activity: %w", err)
		fmt.Println(wrapped)
	}
}
