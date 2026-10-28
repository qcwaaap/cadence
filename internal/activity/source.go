package activity

import "errors"

var ErrNotImplemented = errors.New("cadence source not implemented")

// CadenceSource это любой источник показаний каденса:
// Strava, BLE-датчик или заглушка для тестов.
type CadenceSource interface {
	Readings() (<-chan int, error)
}

// MockSource отдаёт заранее заданные значения.
type MockSource struct {
	Values []int
}

func (m MockSource) Readings() (<-chan int, error) {
	return SimulateCadenceReadings(m.Values), nil
}

// StravaSource берёт средний каденс тренировки из Strava.
type StravaSource struct {
	ActivityID int64
}

func (s StravaSource) Readings() (<-chan int, error) {
	act, err := FetchActivityFromStrava(s.ActivityID)
	if err != nil {
		return nil, err
	}
	return SimulateCadenceReadings([]int{act.AvgCadence}), nil
}

// BLESource пока заглушка: чтение с датчика появится на этапе с железом.
type BLESource struct{}

func (BLESource) Readings() (<-chan int, error) {
	return nil, ErrNotImplemented
}
