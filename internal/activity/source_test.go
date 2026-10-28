package activity

import (
	"errors"
	"testing"
)

func TestMockSource(t *testing.T) {
	var src CadenceSource = MockSource{Values: []int{75, 80, 85}}

	ch, err := src.Readings()
	if err != nil {
		t.Fatalf("неожиданная ошибка: %v", err)
	}
	var got []int
	for r := range ch {
		got = append(got, r)
	}
	if len(got) != 3 || got[0] != 75 || got[2] != 85 {
		t.Errorf("получили %v", got)
	}
}

func TestStravaSourceReturnsAPIError(t *testing.T) {
	var src CadenceSource = StravaSource{ActivityID: 123}

	_, err := src.Readings()

	var apiErr *StravaApiError
	if !errors.As(err, &apiErr) {
		t.Fatalf("ожидали *StravaApiError, получили %v", err)
	}
	if apiErr.StatusCode != 429 {
		t.Errorf("ожидали код 429, получили %d", apiErr.StatusCode)
	}
}

func TestBLESourceNotImplemented(t *testing.T) {
	var src CadenceSource = BLESource{}

	if _, err := src.Readings(); !errors.Is(err, ErrNotImplemented) {
		t.Errorf("ожидали ErrNotImplemented, получили %v", err)
	}
}
