package matcher

import (
	"errors"
	"testing"
)

func TestStyleTableEstimator(t *testing.T) {
	tests := []struct {
		name    string
		style   string
		cadence int
		want    int
	}{
		{"deep house, 62 об/мин", "deep-house", 62, 124},
		{"deep house, 64 об/мин", "deep-house", 64, 128},
		{"hip-hop, 90 об/мин", "hip-hop", 90, 90},
		{"нулевой каденс", "deep-house", 0, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e, err := NewStyleTableEstimator(tt.style)
			if err != nil {
				t.Fatalf("неожиданная ошибка: %v", err)
			}
			if got := RecommendTrack(e, tt.cadence); got != tt.want {
				t.Errorf("ожидали %d, получили %d", tt.want, got)
			}
		})
	}
}

func TestNewStyleTableEstimatorUnknown(t *testing.T) {
	_, err := NewStyleTableEstimator("polka")
	if !errors.Is(err, ErrUnknownStyle) {
		t.Errorf("ожидали ErrUnknownStyle, получили %v", err)
	}
}
