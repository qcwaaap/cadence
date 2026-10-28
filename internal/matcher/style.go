package matcher

import (
	"errors"
	"fmt"
)

var ErrUnknownStyle = errors.New("unknown music style")

// styleMultipliers: сколько ударов в минуту (BPM) приходится на один оборот
// педалей в минуту. Deep house ~60-64 об/мин → 120-128 BPM, это ×2.
// Hip-hop идёт в темпе ~85-95 BPM, это почти ×1 к привычному каденсу.
var styleMultipliers = map[string]float64{
	"deep-house": 2.0,
	"hip-hop":    1.0,
}

// StyleTableEstimator считает BPM по таблице стилей.
type StyleTableEstimator struct {
	style      string
	multiplier float64
}

func NewStyleTableEstimator(style string) (StyleTableEstimator, error) {
	m, ok := styleMultipliers[style]
	if !ok {
		return StyleTableEstimator{}, fmt.Errorf("%w: %q", ErrUnknownStyle, style)
	}
	return StyleTableEstimator{style: style, multiplier: m}, nil
}

func (e StyleTableEstimator) EstimateTargetBPM(cadence int) int {
	return int(float64(cadence) * e.multiplier)
}
