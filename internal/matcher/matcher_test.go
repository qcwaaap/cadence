package matcher

import "testing"

// fakeEstimator нужен, чтобы проверить RecommendTrack отдельно от
// настоящего расчёта: он просто возвращает заранее заданное число.
type fakeEstimator struct {
	result int
}

func (f fakeEstimator) EstimateTargetBPM(cadence int) int {
	return f.result
}

func TestRecommendTrack(t *testing.T) {
	tests := []struct {
		name      string
		estimator BPMEstimator
		cadence   int
		want      int
	}{
		{"множитель 2, каденс 80", SimpleMultiplierEstimator{Multiplier: 2.0}, 80, 160},
		{"множитель 2, каденс 90", SimpleMultiplierEstimator{Multiplier: 2.0}, 90, 180},
		{"множитель 1.5, каденс 80", SimpleMultiplierEstimator{Multiplier: 1.5}, 80, 120},
		{"дробный результат отбрасывается", SimpleMultiplierEstimator{Multiplier: 1.5}, 75, 112},
		{"нулевой каденс", SimpleMultiplierEstimator{Multiplier: 2.0}, 0, 0},
		{"нулевой множитель", SimpleMultiplierEstimator{Multiplier: 0}, 80, 0},
		{"подставной estimator", fakeEstimator{result: 128}, 62, 128},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := RecommendTrack(tt.estimator, tt.cadence)
			if got != tt.want {
				t.Errorf("RecommendTrack(cadence=%d) = %d, ожидали %d", tt.cadence, got, tt.want)
			}
		})
	}
}
