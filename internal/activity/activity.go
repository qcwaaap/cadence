package activity

import (
	"encoding/json"
	"fmt"
)

type StravaActivityJSON struct {
	ID         int64   `json:"id"`
	Name       string  `json:"name"`
	AvgCadence float64 `json:"average_cadence"`
}

type Activity struct {
	ID         int64
	AvgCadence int
	DistanceKm float64
}

type StravaApiError struct {
	StatusCode int
	EndPoint   string
}

func (e *StravaApiError) Error() string {
	return fmt.Sprintf("strava API error: GET %s returned %d", e.EndPoint, e.StatusCode)
}

func FetchActivityFromStrava(id int64) (*Activity, error) {
	return nil, &StravaApiError{
		StatusCode: 429,
		EndPoint:   "/activities",
	}
}

func (a *Activity) UpdateCadence(newCadence int) {
	a.AvgCadence = newCadence
}

func ParseActivity(data []byte) (*StravaActivityJSON, error) {
	var act StravaActivityJSON
	err := json.Unmarshal(data, &act)
	if err != nil {
		return nil, fmt.Errorf("failed to parse activity JSON: %w", err)
	}
	return &act, nil
}

func SimulateCadenceReadings(readings []int) <-chan int {
	ch := make(chan int)

	go func() {
		for _, r := range readings {
			ch <- r
		}
		close(ch)
	}()
	return ch
}
