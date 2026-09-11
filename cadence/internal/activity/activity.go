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
func (a Activity) EstimateTargetBPM() int {
	return int(float64(a.AvgCadence) * 2.0)
} // пока так потом добавить таблицу интерфейс для расчета каденса

func (a *Activity) UpdateCadence(newCadence int) {
	a.AvgCadence = newCadence
}



