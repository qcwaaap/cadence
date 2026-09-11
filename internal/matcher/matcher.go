package matcher

type BPMEstimator interface {
	EstimateTargetBPM(cadence int) int
}

type SimpleMultiplierEstimator struct {
	Multiplier float64
}

func (e SimpleMultiplierEstimator) EstimateTargetBPM(cadence int) int {
	return int(float64(cadence) * e.Multiplier)
}

func RecommendTrack(estimator BPMEstimator, cadence int) int {
	return estimator.EstimateTargetBPM(cadence)
}
